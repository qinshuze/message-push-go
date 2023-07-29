package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ArrayToString[T any](v []T) string {
	var str = strings.Builder{}
	for i, t := range v {
		if i != 0 {
			str.WriteString(",")
		}

		if any(t) == nil {
			str.WriteString("nil")
			continue
		}

		switch v1 := any(t).(type) {
		case string:
			str.WriteString(v1)
			continue
		case int:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int8:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int16:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int32:
			str.WriteString(strconv.FormatInt(int64(v1), 10))
			continue
		case int64:
			str.WriteString(strconv.FormatInt(v1, 10))
			continue
		case byte:
			str.WriteString(string(v1))
			continue
		case float32:
			str.WriteString(strconv.FormatFloat(float64(v1), 'f', -1, 64))
			continue
		case float64:
			str.WriteString(strconv.FormatFloat(v1, 'f', -1, 64))
			continue
		case bool:
			str.WriteString(strconv.FormatBool(v1))
			continue
		}
	}

	return str.String()
}

func T1() {
	startTime := time.Now().UnixNano()

	//sampleRegexp := regexp.MustCompile(`[\[\]]`)
	//for i := 0; i < 1000000; i++ {
	//	sampleRegexp.ReplaceAllString("[1, 2, 3]", "")
	//}
	//
	//for i := 0; i < 1000000; i++ {
	//	strings.ReplaceAll("[1, 2, 3]", "[", "")
	//}

	//for i := 0; i < 1000000; i++ {
	//	strings.Trim(strings.Trim("[1, 2, 3]", "["), "]")
	//}

	//fmt.Println(strings.Trim(strings.Trim("[1, 2, 3]", "["), "]"))
	//fmt.Println(fmt.Sprintf("%v", []int{1,2,3}))
	//for i := 0; i < 1000000; i++ {
	//	ArrayToString([]int{1,2,3})
	//	strings.Trim(strings.Trim("[1, 2, 3]", "["), "]")
	//}

	//fmt.Println(ArrayToString([]any{1,"2",'3',4.001,false,[]string{"a","b","c"}}))

	//fmt.Println(str.String())
	//fmt.Println(ArrayToString([]int{1, 2, 3}))

	for i := 0; i < 1000000; i++ {
		var str = strings.Builder{}
		for i, i3 := range []int{1, 2, 3} {
			if i != 0 {
				str.WriteString(",")
			}
			str.WriteString(strconv.FormatInt(int64(i3), 10))
		}
		str.String()
	}

	fmt.Println((time.Now().UnixNano() - startTime) / 1000000)
}

type DataChannel struct {
	data   []byte
	status string
}
type Reader struct {
	state       string
	req         *http.Request
	ioReader    *bufio.Reader
	dataChannel chan DataChannel
}

func (r *Reader) Pause() {
	r.state = "pause"
}
func (r *Reader) Start() {
	r.state = "reader"
}

func NewReader(r *http.Request) *Reader {
	reader := &Reader{state: "pause", dataChannel: make(chan DataChannel)}
	reader.req = r
	reader.ioReader = bufio.NewReader(r.Body)
	return reader
}

var m = make(map[string]*Reader)

func E1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "Close")

	req := m["10"]
	req.Start()
	for {
		select {
		case dataChannel := <-req.dataChannel:
			switch dataChannel.status {
			case "read":
				_, err := w.Write(dataChannel.data)
				if err != nil {
					fmt.Println("写入错误：", err)
					return
				}
				break
			case "end":
				_, err := w.Write(dataChannel.data)
				if err != nil {
					fmt.Println("写入错误：", err)
					return
				}
				return
			case "error":
				w.WriteHeader(502)
			case "outTime":
				w.WriteHeader(504)
			}
		case <-time.After(60 * time.Second):
			w.WriteHeader(408)
			return
		}
	}
}

func T2(w http.ResponseWriter, r *http.Request) {
	//uid := strconv.FormatInt(int64(time.Now().Second()), 10)
	uid := "10"
	reader := NewReader(r)
	outTime := time.Now().UnixMilli() + 60000
	readeEndHandler := func(status string) {
		fmt.Println(status)
		delete(m, "10")
	}

	fmt.Println("uuid: ", uid)
	m[uid] = reader
	for {
		if time.Now().UnixMilli() > outTime {
			readeEndHandler("超时")
			reader.dataChannel <- DataChannel{data: []byte{}, status: "outTime"}
			break
		}

		if reader.state == "pause" {
			continue
		}

		bytes, err := reader.ioReader.ReadBytes('\n')
		if err == io.EOF {
			reader.dataChannel <- DataChannel{data: bytes, status: "end"}
			readeEndHandler("读取结束")
			break
		}

		if err != nil {
			reader.dataChannel <- DataChannel{data: bytes, status: "error"}
			readeEndHandler("失败：" + err.Error())
			break
		}

		if len(bytes) > 0 {
			outTime = time.Now().UnixMilli() + 60000
			reader.dataChannel <- DataChannel{data: bytes, status: "read"}
			fmt.Println(string(bytes))
		}
	}
}

func T3(ch chan<- bool) {
	ch <- true
}

type QRCode struct {
	Id         string `json:"id"`
	ClientIp   string `json:"client_ip"`
	ClientPort string `json:"client_port"`
	DeviceId   string `json:"device_id"`
	DeviceIp   string `json:"device_ip"`
	DevicePort string `json:"device_port"`
	Status     string `json:"status"`
	Expires    int64  `json:"_"`
}

var qrMap = map[string]*QRCode{}

func Q1(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var expires, _ = strconv.ParseInt(r.URL.Query().Get("expires"), 10, 64)
	var result = qrMap[id]

	if result == nil {
		result = &QRCode{Id: id, Status: "not", Expires: expires}
		qrMap[id] = result
		log.Println(result)
	} else if result.Status == "complete" {
		delete(qrMap, id)
	}

	w.Header().Set("Content-Type", "application/json")
	marshal, _ := json.Marshal(map[string]string{
		"id":        result.Id,
		"status":    result.Status,
		"device_id": result.DeviceId,
	})
	_, _ = w.Write(marshal)
}

func Q3(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var deviceId = r.URL.Query().Get("device_id")
	var result = map[string]any{"msg": "ok", "code": 200}

	if qrMap[id] != nil {
		if (time.Now().UnixMilli() / 1000) > qrMap[id].Expires {
			result["msg"] = "二维码过期"
			result["code"] = 423
			delete(qrMap, id)
		} else if (time.Now().UnixMilli() / 1000) < qrMap[id].Expires {
			qrMap[id].Status = "complete"
		} else if deviceId != qrMap[id].DeviceIp {
			result["msg"] = "非法操作"
			result["code"] = 400
		} else {
			result["msg"] = "非法操作"
			result["code"] = 400
		}
	} else {
		result["msg"] = "无效id"
		result["code"] = 422
	}

	w.Header().Set("Content-Type", "application/json")
	marshal, _ := json.Marshal(result)
	_, _ = w.Write(marshal)
}

func Q2(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var deviceId = r.URL.Query().Get("device_id")
	var result = map[string]any{"msg": "ok", "code": 200}
	if qrMap[id] == nil {
		result["msg"] = "无效id"
		result["code"] = 422
	} else if (time.Now().UnixMilli() / 1000) < qrMap[id].Expires {
		log.Println(time.Now().UnixMilli()/1000, qrMap[id].Expires)
		qrMap[id].Status = "used"
		qrMap[id].DeviceId = deviceId
	} else {
		result["msg"] = "二维码过期"
		result["code"] = 423
		delete(qrMap, id)
	}

	w.Header().Set("Content-Type", "application/json")
	marshal, _ := json.Marshal(result)
	_, _ = w.Write(marshal)
}

// 自定义中间件函数来拦截HTTP请求
func interceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 执行任何预处理逻辑
		fmt.Println("拦截请求:", r.URL.Path)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		// 调用下一个处理器
		next.ServeHTTP(w, r)
	})
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/qrcode", Q1)
	server.HandleFunc("/qrcode/step1", Q2)
	server.HandleFunc("/qrcode/step2", Q3)

	handler := interceptor(server)
	log.Fatalln(http.ListenAndServe("0.0.0.0:8004", handler))
}
