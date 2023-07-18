package message

import (
	"ccps.com/service/ws/ws"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"k8s.io/utils/strings/slices"
	"sync"
	"time"
)

var clients []*Client

type ClientCollect struct {
	clients []*Client
}

var operateLock = sync.Mutex{}

func init() {
	startHeartbeat()
}

// startHeartbeat 开启客户端心跳检测
func startHeartbeat() {
	time.AfterFunc(10*time.Second, func() {
		for _, client := range clients {
			// 如果客户端非活动连接，则将其关闭
			if !client.isLive {
				err := client.Close(websocket.CloseAbnormalClosure, "CLOSE_ABNORMAL")
				if err != nil {
					color.Red("关闭客户端连接失败：%s", err.Error())
				}
			} else {
				err := client.Ping()
				if err != nil {
					color.Red("发送心跳异常：%s", err.Error())
				}
			}
		}
		startHeartbeat()
	})
}

// NewClientCollect 新建客户端集合
func NewClientCollect(clients []*Client) *ClientCollect {
	return &ClientCollect{
		clients: clients,
	}
}

// Clients 获取客户端集合
func Clients() *ClientCollect {
	return NewClientCollect(clients)
}

// Append 将加客户端添加到连接池
func Append(c *Client) {
	operateLock.Lock()
	defer operateLock.Unlock()

	clients = append(clients, c)
}

// Remove 将客户端从连接池中移除
func Remove(c *Client) {
	operateLock.Lock()
	defer operateLock.Unlock()

	var _tmp []*Client
	for _, c2 := range clients {
		if c2 != c {
			_tmp = append(_tmp, c2)
		}
	}

	clients = _tmp
}

// GetByName 根据名称获取客户端连接
func (co *ClientCollect) GetByName(name string) *ClientCollect {
	var _tmp []*Client
	for _, client := range co.clients {
		if client.name == name {
			_tmp = append(_tmp, client)
		}
	}

	co.clients = _tmp
	return co
}

// GetByTag 根据标签获取客户端连接
func (co *ClientCollect) GetByTag(tag string) *ClientCollect {
	var _tmp []*Client
	for _, client := range co.clients {
		if slices.Index(client.tags, tag) != -1 {
			_tmp = append(_tmp, client)
		}
	}

	co.clients = _tmp
	return co
}

// GetByRoomId 根据房间号获取客户端连接
func (co *ClientCollect) GetByRoomId(roomId string) *ClientCollect {
	var _tmp []*Client
	for _, client := range co.clients {
		if slices.Index(client.roomIds, roomId) != -1 {
			_tmp = append(_tmp, client)
		}
	}

	co.clients = _tmp
	return co
}

// GetByRealm 根据域获取客户端连接
func (co *ClientCollect) GetByRealm(realm string) *ClientCollect {
	var _tmp []*Client
	for _, client := range co.clients {
		if client.realm == realm {
			_tmp = append(_tmp, client)
		}
	}

	co.clients = _tmp
	return co
}

// Send 发送消息给到集合里所有的客户端
func (co *ClientCollect) Send(msg SendMessage) {
	go func() {
		for _, client := range co.clients {
			err := client.Send(msg)
			if err != nil {
				color.Red("消息发送失败：%s", err.Error())
			}
		}
	}()
}

// Filter 过滤客户端
func (co *ClientCollect) Filter(c ...*Client) *ClientCollect {
	var _tmp []*Client
	for _, client := range co.clients {
		var skip = false
		for _, c2 := range c {
			if client == c2 {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		_tmp = append(_tmp, client)
	}

	co.clients = _tmp
	return co
}

// Slice 将集合转为切片
func (co *ClientCollect) Slice() []*Client {
	return co.clients
}

// Count 获取客户端连接数量
func (co *ClientCollect) Count() int {
	return len(co.clients)
}

// GetBySelectOption 根据查询选项获取客户端连接
func (co *ClientCollect) GetBySelectOption(options *ws.SelectOptions) *ClientCollect {
	for _, realm := range options.Realms {
		co.GetByRealm(realm)
	}

	for _, id := range options.RoomIds {
		co.GetByRoomId(id)
	}

	for _, tag := range options.Tags {
		co.GetByTag(tag)
	}

	for _, name := range options.Names {
		co.GetByName(name)
	}

	return co
}

// Count 获取客户端连接数量
func Count() int {
	return len(clients)
}
