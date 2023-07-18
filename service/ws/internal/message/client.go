package message

import (
	"ccps.com/service/ws/internal/svc"
	"context"
	"encoding/base64"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/rest/httpc"
	"k8s.io/utils/strings/slices"
	"net/http"
	"strings"
	"sync"
)

import (
	"encoding/json"
)

// 初始化websocket连接选项
var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义客户端连接状态常量
const (
	ReadyStateNot        int = 0
	ReadyStateOpen       int = 1
	ReadyStateClose      int = 2
	ReadyStateDisconnect int = 3
)

// Client 客户端连接对象结构体
type Client struct {
	id         string
	name       string
	roomIds    []string
	tags       []string
	conn       *websocket.Conn
	isLive     bool
	realm      string
	svcCtx     *svc.ServiceContext
	extendData map[string]string
	readyState int
}

// openHandler 客户端连接成功处理方法
func (c *Client) openHandler(ctx *svc.ServiceContext) {
	Append(c)
	c.readyState = ReadyStateOpen
	color.Green("客户端[%s::%s]连接成功，当前客户端连接数：%d", c.id, c.name, Count())
}

// closeHandler 客户端连接关闭处理方法
// 当客户端连接关闭时触发该方法
func (c *Client) closeHandler(ctx *svc.ServiceContext) {
	c.conn.SetCloseHandler(func(code int, text string) error {
		// 离开房间
		for _, s := range c.RoomIds() {
			c.LeaveRoom(s)
		}

		// 将其从连接池中移除
		Remove(c)

		c.readyState = ReadyStateClose

		// 如果是非活动连接，则将其状态更改为断开
		if !c.isLive {
			c.readyState = ReadyStateDisconnect
		}

		// 将其设为非活动连接
		c.isLive = false
		color.Red("客户端[%s::%s]连接关闭：%d %s，当前客户端连接数：%d", c.id, c.name, code, text, Count())
		return nil
	})
}

// pongHandler 心跳包处理方法
// 当接收到新的ping心跳检测包后触发该方法
func (c *Client) pongHandler(ctx *svc.ServiceContext) {
	c.conn.SetPongHandler(func(appData string) error {
		// 收到ping心跳检测包后将当前客户端连接设置为活动连接
		c.isLive = true
		return nil
	})
}

// messageHandler 客户端消息处理方法
// 循环等待处理消息
func (c *Client) messageHandler(ctx *svc.ServiceContext) {
	for {
		messageType, p, err := c.conn.ReadMessage()
		if err != nil {
			color.Red("消息接收失败：%s", err.Error())
			_ = c.Close(websocket.CloseGoingAway, err.Error())
			return
		}

		if messageType != websocket.TextMessage {
			color.Red("无效消息类型：%s", messageType)
			_ = c.Close(websocket.CloseUnsupportedData, "CLOSE_UNSUPPORTED")
			return
		}

		message, err := NewReceiveMessageByJsonStr(string(p))
		if err != nil {
			color.Red("消息解析错误: %s - %s", err.Error(), string(p))
			_ = c.Close(websocket.CloseUnsupportedData, "CLOSE_UNSUPPORTED")
			return
		}

		c.messageRelay(message)
	}
}

// messageRelay 转发消息
func (c *Client) messageRelay(message *ReceiveMessage) {
	color.Blue("接收到消息：%v", message)
	sendMsg := SendMessage{Type: Relay, Content: message.Content(), Sender: c.name}
	clientColl := Clients().GetByRealm(c.realm).Filter(c)

	// 如果房间号存在，则发送给指定房间号下的所有客户端
	for _, id := range message.roomIds {
		clientColl.GetByRoomId(id)
	}

	// 如果接收人存在，则发送给到指定的接收人客户端
	for _, s := range message.names {
		clientColl.GetByName(s)
	}

	// 如果标签存在，则发送给到拥有指定标签的客户端
	for _, s := range message.tags {
		clientColl.GetByTag(s)
	}

	clientColl.Send(sendMsg)
}

// NewClient 新建客户端连接
func NewClient(ctx *svc.ServiceContext, r *http.Request, w http.ResponseWriter) *Client {
	// 解析请求参数
	queryParams := r.URL.Query()
	extendData := map[string]string{}
	paramsExtendData := queryParams.Get("extend_data")
	if paramsExtendData != "" {
		var decodeOk = true
		decodeString, err := base64.StdEncoding.DecodeString(paramsExtendData)
		if err != nil {
			decodeOk = false
		}

		err = json.Unmarshal(decodeString, &extendData)
		if err != nil {
			decodeOk = false
		}

		if !decodeOk {
			color.Red("参数解析失败：%s - %s", err.Error(), paramsExtendData)
			w.WriteHeader(422)
			_, _ = w.Write([]byte("param `extend_data` analysis error, invalid data format."))
			return nil
		}
	}

	// 升级客户端请求为websocket请求
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		color.Red("客户端请求提升失败：%s", err.Error())
		w.WriteHeader(400)
		return nil
	}

	// 客户端认证
	if !clientAuth(ctx, r, w) {
		err := conn.CloseHandler()(3001, "client auth fail")
		if err != nil {
			color.Red("CloseHandler：%s", err.Error())
			return nil
		}
		return nil
	}

	// 创建客户端连接对象
	client := &Client{
		id:         uuid.New().String(),
		name:       queryParams.Get("name"),
		tags:       strings.Split(queryParams.Get("tags"), ","),
		realm:      queryParams.Get("realm"),
		extendData: extendData,
		isLive:     true,
		conn:       conn,
		readyState: ReadyStateNot,
	}

	// 为客户端连接添加事件侦听处理
	client.openHandler(ctx)
	client.closeHandler(ctx)
	go client.pongHandler(ctx)
	go client.messageHandler(ctx)

	// 将客户端连接加入到指定房间
	roomIds := strings.Split(queryParams.Get("room_ids"), ",")
	for _, id := range roomIds {
		if id != "" {
			client.JoinRoom(id)
		}
	}

	return client
}

// clientAuth 客户端连接认证
// 如果在配置中填写了clientAuth项，则视为开启客户端认证
// 客户端认证开启后，当有新的客户端连接加入，则会先将客户端请求转发给到clientAuth选项中配置的地址
func clientAuth(ctx *svc.ServiceContext, r *http.Request, w http.ResponseWriter) bool {

	// 检查当前请求域名是否在白名单内，如果在白名单内则，跳过认证步骤
	if slices.Index(ctx.Config.WebSocket.WhiteList, strings.Split(r.Host, ":")[0]) != -1 {
		return true
	}

	// 如果客户端认证选项不为空，则对当前客户端请求进行认证
	if ctx.Config.WebSocket.ClientAuth != "" {
		// 将当前客户端请求转发到认证服务器
		var u = ctx.Config.WebSocket.ClientAuth + "?" + r.URL.Query().Encode()
		res, err := httpc.Do(context.Background(), http.MethodGet, u, nil)
		if err != nil {
			color.Red("请求客户端认证服务器失败：%s", err.Error())
			//w.Header().Add("Connection", "close")
			//w.WriteHeader(503)
			return false
		}

		if res.StatusCode != 200 {
			color.Red("客户端认证失败：%s - %v", res.Status, u)
			//w.Header().Add("Connection", "close")
			//w.WriteHeader(401)
			return false
		}
	}
	return true
}

// =============Getter Setter Start================

func (c *Client) Id() string {
	return c.id
}

func (c *Client) IsLive() bool {
	return c.isLive
}

func (c *Client) Ping() error {
	sendLock.Lock()
	defer sendLock.Unlock()

	c.isLive = false
	return c.conn.WriteMessage(websocket.PingMessage, []byte{})
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) RoomIds() []string {
	return c.roomIds
}

func (c *Client) Tags() []string {
	return c.tags
}

func (c *Client) Realm() string {
	return c.realm
}

func (c *Client) ExtendData() map[string]string {
	return c.extendData
}

func (c *Client) ReadyState() int {
	return c.readyState
}

// =============Getter Setter End================

// CheckState 检查客户端连接状态
func (c *Client) CheckState(state int) bool {
	return c.readyState == state
}

var sendLock = sync.RWMutex{}

// Send 发送消息到当前客户端
func (c *Client) Send(message SendMessage) error {
	sendLock.Lock()
	defer sendLock.Unlock()

	marshal, _ := json.Marshal(message)
	return c.conn.WriteMessage(websocket.TextMessage, marshal)
}

// Close 关闭当前客户端连接
func (c *Client) Close(code int, text string) error {
	if c.readyState == ReadyStateClose || c.readyState == ReadyStateDisconnect {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return err
	}

	return c.conn.CloseHandler()(code, text)
}

// JoinRoom 将当前客户端连接加入到指定房间
func (c *Client) JoinRoom(roomId string) {
	c.roomIds = append(c.roomIds, roomId)
	var roomIdMap = map[string]int{}
	for i, id := range c.roomIds {
		roomIdMap[id] = i
	}
	var newRoomIds []string
	for s := range roomIdMap {
		newRoomIds = append(newRoomIds, s)
	}
	c.roomIds = newRoomIds

	// 发送消息给到房间内所有人，告知有新的客户端连接加入
	content, err := json.Marshal(map[string]any{
		"id":          c.id,
		"name":        c.name,
		"tags":        c.tags,
		"extend_data": c.extendData,
	})
	if err != nil {
		color.Red("消息内容解析错误：%s", err.Error())
		return
	}
	sendMsg := SendMessage{Type: Join, Content: string(content), Sender: c.name}
	Clients().GetByRealm(c.realm).GetByRoomId(roomId).Filter(c).Send(sendMsg)
	color.Green("客户端[%s::%s]加入房间[%s]，当前房间人数 %d",
		c.id,
		c.name,
		roomId,
		Clients().GetByRealm(c.realm).GetByRoomId(roomId).Count(),
	)
}

// LeaveRoom 将当前客户端连接从指定房间中移除
func (c *Client) LeaveRoom(roomId string) {
	var newRoomIds []string
	for _, id := range c.roomIds {
		if id != roomId {
			newRoomIds = append(newRoomIds, id)
		}
	}

	c.roomIds = newRoomIds

	// 发送消息给到房间内所有人，告知有客户端连接离开
	content, err := json.Marshal(map[string]any{
		"id":          c.id,
		"name":        c.name,
		"tags":        c.tags,
		"extend_data": c.extendData,
	})
	if err != nil {
		color.Red("消息内容解析错误：%s", err.Error())
		return
	}

	sendMsg := SendMessage{Type: Leave, Content: string(content), Sender: c.name}
	Clients().GetByRealm(c.realm).GetByRoomId(roomId).Filter(c).Send(sendMsg)
	color.Green("客户端[%s::%s]离开房间[%s]，当前房间人数 %d",
		c.id,
		c.name,
		roomId,
		Clients().GetByRealm(c.realm).GetByRoomId(roomId).Count(),
	)
}

// AddTag 为当前客户端连接添加标签
func (c *Client) AddTag(tag string) {
	c.tags = append(c.tags, tag)

	var tagMap = map[string]int{}
	for i, id := range c.tags {
		tagMap[id] = i
	}

	var newTags []string
	for s := range tagMap {
		newTags = append(newTags, s)
	}

	c.tags = newTags
}

// RemoveTag 从当前客户端连接中移除指定标签
func (c *Client) RemoveTag(tag string) {
	var newTags []string
	for _, id := range c.tags {
		if id != tag {
			newTags = append(newTags, id)
		}
	}

	c.tags = newTags
}
