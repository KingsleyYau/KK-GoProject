package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"pushserver"
	"strconv"
//	"syscall"
	"time"
)

var clientMap 	*pushserver.ClientMap
var appMap 		*pushserver.AppMap

func sendRespone(bf *bufio.Writer, buf []byte) error {
	repmsg := "HTTP/1.1 200 OK\r\n"
	repmsg += "Connection: keep-alive\r\n"
	repmsg += "Content-Length: " + strconv.Itoa(len(buf)) + "\r\n"
	repmsg += "Content-type: text/html; charset=utf8\r\n"
	repmsg += "\r\n"
		
	repbuf := append([]byte(repmsg), buf...)	
			
	_, err := bf.Write([]byte(repbuf))
	if err != nil {
		return err
	}
	err = bf.Flush()
	if err != nil {
		return err
	}
	return nil
}
func readRequest(conn net.Conn, b *bufio.Reader, timeout time.Duration) (req *http.Request, err error) {
	log.Println("########## " + conn.RemoteAddr().String() + " 等待下次请求 #########")
	if timeout > 0 {
		// 读取客户端请求, 阻塞, timeout秒超时, 防止请求挑战值后占用连接
		conn.SetReadDeadline(time.Now().Add(timeout))
	} else {
		conn.SetReadDeadline(time.Time{})
	}

  	if req, err = http.ReadRequest(b); err != nil {
//  		if err == io.EOF {
//  			log.Println(conn.RemoteAddr().String(), "读取请求失败, 原因:", err.Error())
//  	  	} else {
//  	  		if terr, ok := err.(net.Error); ok && terr.Timeout() {
//  	  			log.Println(conn.RemoteAddr().String(), "读取请求失败, 超时:", terr.Error())
//  	  		}
//  	  	}
		// 连接已经断开
		return req, err
	}
	log.Println("########## " + conn.RemoteAddr().String() + " 收到请求 #########")
	req.RemoteAddr = conn.RemoteAddr().String()
	log.Println(conn.RemoteAddr().String(), "Http请求 Header:", req.Proto, req.Header)
	return req, nil
}

/*
 * 处理客户端和服务端推送协议协程
 * 协程由http包里面开启, 这里只做逻辑处理
 */
func clientHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("来自:", r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	log.Println("Connection:", r.Header.Get("Connection"))
	w.Header().Set("Connection", "keep-alive")
    
    // 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var req *http.Request
	
	req = r
	
	// 逻辑
	const (
		GetChallenge = iota
		Register
		GetPushMessage
		RegisterServer
		PushMessage
	)
	const (
		ClientRequest = iota
		ServerRequest
	)
	setpmsg := []string{
			"请求Challenge",
			"注册客户端Token",
			"请求接受推送",
			"注册服务端",
			"发送推送消息",
			}
	var step int
	var ctype int
	var err error
    var msg string
    var token string
    var repmsg string
    var challenge string // 挑战值
    var checkcode string // 检验码
    var timeout time.Duration
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println("械劫连接", r.RemoteAddr, "失败, 原因:", err.Error())
			return
		}
		// 
		defer log.Println("关闭连接:", r.RemoteAddr)
		defer conn.Close()
		log.Println("械劫连接", r.RemoteAddr, "成功")
	}
		
	step = GetChallenge
	ctype = ClientRequest
	
	
	for {	
		// 默认阻塞
		timeout = 0
		
		// 第一次连接不用读取http header, 所以先读取 http body
		buf, err := ioutil.ReadAll(req.Body)
	    if err != nil {
	    	log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    	return
	    }	
	    req.Body.Close()
		log.Println(conn.RemoteAddr().String(), "Http请求 Body:", string(buf))
		
		// 解析协议
		var jsonMap map[string]interface{}
		if err := json.Unmarshal(buf, &jsonMap); err != nil {
			log.Println(conn.RemoteAddr().String(), "解析json失败:", err.Error())
			sendRespone(bf.Writer, []byte("解析json失败, 非法请求"))
			return
		}
		
		// 处理请求
		if cmd, ok := jsonMap["cmd"]; ok {
			switch cmd {
				case "getserverchanllenge":{
					// 服务端第一步, GetChallenge
					msg = "服务端请求Challenge"
					log.Println(conn.RemoteAddr().String(), msg)
					
					// http body
					jsonData := make(map[string]interface{})
					jsonData["ret"] = "2001"
						
					if step != GetChallenge {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因: 非法请求:[", setpmsg[0], "]", "当前步骤应该为:[", setpmsg[step] ,"]")
					} else {
						jsonData["ret"] = "0"
						un := time.Now().Unix()
						challenge = fmt.Sprintf("%x%x", un, un << 2)
						jsonData["chanllenge"] = challenge
					}

					jsonbuf, _ := json.Marshal(jsonData)
					
					err = sendRespone(bf.Writer, []byte(jsonbuf))
					if err != nil {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
						return
					}
					
					if jsonData["ret"] != "0" {
						log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", jsonData["ret"])
						return
					}
					
					// 请求Challenge成功
					log.Println(conn.RemoteAddr().String(), msg, "成功")
					step = RegisterServer
					ctype = ServerRequest
					timeout = 15 * time.Second
				}
				case "getclientchanllenge": {
					// 客户端第一步, GetChallenge
					msg = "客户端请求Challenge"
					log.Println(conn.RemoteAddr().String(), msg)
					
					jsonData := make(map[string]interface{})
					jsonData["ret"] = "2001"
					
					if step != GetChallenge {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因: 非法请求:[", setpmsg[0], "]", "当前步骤应该为:[", setpmsg[step] ,"]")
					} else {
						jsonData["ret"] = "0"
						un := time.Now().Unix()
						challenge = fmt.Sprintf("%x%x", un, un << 2)
						jsonData["chanllenge"] = challenge
					}
					
					jsonbuf, _ := json.Marshal(jsonData)
					err = sendRespone(bf.Writer, []byte(jsonbuf))
					if err != nil {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
						return
					}

					if jsonData["ret"] != "0" {
						log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", jsonData["ret"])
						return
					}
					
					// 请求Challenge成功
					log.Println(conn.RemoteAddr().String(), msg, "成功")
					step = Register
					ctype = ClientRequest
					timeout = 15 * time.Second
				}
				case "register": {
					// 第二步, Register
					msg = "注册"
					log.Println(conn.RemoteAddr().String(), msg)
					
					// http body
					jsonData := make(map[string]interface{})
					jsonData["ret"] = "2001"
					
					if ctype == ClientRequest {
						// 客户端
						if step != Register {
							log.Println(conn.RemoteAddr().String(), msg, "失败, 原因: 非法请求:[", setpmsg[1], "]", "当前步骤应该为:[", setpmsg[step] ,"]")
						} else {
							type body struct {
								AppId		string		`json:"appid"`
								CheckCode	string		`json:"checkcode"`
								AppVer		string		`json:"appver"`
								Model		string		`json:"model"`
								System		string		`json:"system"`
							}
							type jsonStruct struct {	
								TokenId	string	`json:"tokenid"`
								Body	body	`json:"body"`
							}
						
							var js jsonStruct
							if err := json.Unmarshal(buf, &js); err != nil {
								log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
							}

							if js.TokenId != "" && js.Body.CheckCode != "" && js.Body.AppId != "" {
								// 生成校验码
								token = js.TokenId
								appKey, _ := appMap.Get(js.Body.AppId)
							
//								appKey = md5tohex(appKey)
								checkcode = js.TokenId + challenge + appKey;
								checkcode = pushserver.Md5tohex(checkcode)
								
								if checkcode == js.Body.CheckCode {
									// 注册成功, 加入在线列表, 创建消息队列
									c := pushserver.NewClient(token)
									c.Address = conn.RemoteAddr().String()
									c.AppVer = js.Body.AppVer
									c.Model = js.Body.Model
									c.System = js.Body.System
									
									bFlag := clientMap.Set(token, c)
									if bFlag {	
										logString := fmt.Sprintf("%s, %s, %s, %s, %s, 客户端上线", 
																conn.RemoteAddr().String(),
																token, 
																js.Body.AppVer, 
																js.Body.Model, 
																js.Body.System)
										pushserver.LogFile(logString)
										jsonData["ret"] = "0"
										// 断线销毁消息队列, 移除在线列表
										defer func() {
											log.Println(conn.RemoteAddr().String(), "断线移除在线列表")
											logString := fmt.Sprintf("%s, %s, %s, %s, %s, 客户端下线", 
																conn.RemoteAddr(),
																token, 
																js.Body.AppVer, 
																js.Body.Model, 
																js.Body.System)
											pushserver.LogFile(logString)
											clientMap.Delete(token)
										}()
		
										step = GetPushMessage
									}
								}
							}
						}
					} else {
						// 服务端
						if step != RegisterServer {
							log.Println(conn.RemoteAddr().String(), msg, "失败, 原因: 非法请求:[", setpmsg[3], "]", "当前步骤应该为:[", setpmsg[step] ,"]")
						} else {
							type body struct {
								AppId		string		`json:"appid"`
								CheckCode	string		`json:"checkcode"`
							}
							type jsonStruct struct {	
								Body	body	`json:"body"`
							}
							var js jsonStruct
							if err := json.Unmarshal(buf, &js); err != nil {
								log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
							}
							if js.Body.CheckCode != "" && js.Body.AppId != "" {
								// 生成校验码
								appKey, _ := appMap.Get(js.Body.AppId)
								
//								appKey = md5tohex(appKey)
								checkcode = challenge + appKey
								checkcode = pushserver.Md5tohex(checkcode)
								
								if true/*checkcode == js.Body.CheckCode*/ {
									jsonData["ret"] = "0"
									// 注册成功
									step = PushMessage
								}
							}
						}
					}

					// 	返回数据
					jsonbuf, _ := json.Marshal(jsonData)
					
					err = sendRespone(bf.Writer, []byte(jsonbuf))
					if err != nil {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
						return
					}
					
					if jsonData["ret"] != "0" {
						log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", jsonData["ret"])
						return
					}
					log.Println(conn.RemoteAddr().String(), msg, "成功")
				}
				case "pushmsg": {
					// 服务端第三步, 推送消息
					msg = "推送消息"
					log.Println(conn.RemoteAddr().String(), msg)
					
					// http body
					jsonData := make(map[string]interface{})
					jsonData["ret"] = "2001"
					
					if ctype != ServerRequest || step !=  PushMessage {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因: 非法请求:[", setpmsg[4], "]", "当前步骤应该为:[", setpmsg[step] ,"]")
					} else {
						type body struct {
							TokenIdList	[]string	`json:"tokenidlist"`
							Body		string		`json:"body"`
						}
						type jsonStruct struct {	
							Body	body	`json:"body"`
						}
						
						var js jsonStruct
						if err := json.Unmarshal(buf, &js); err != nil {
							log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
						} 
						
						if len(js.Body.TokenIdList) > 0 {
							// 发送成功
							jsonData["ret"] = "0"
							msg := pushserver.Message {
									Content:js.Body.Body,
									RecvTime:time.Now().Format("2006-01-02 15:04:05"),
							}
							clientMap.SendMessage(js.Body.TokenIdList, msg)
							
						}
					}
						
					// 	返回数据
					jsonbuf, _ := json.Marshal(jsonData)
					
					err = sendRespone(bf.Writer, []byte(jsonbuf))
					if err != nil {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因:", err.Error())
						return
					}
					
					if jsonData["ret"] != "0" {
						log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", jsonData["ret"])
						return
					}

				}
				case "getpushmessage": {
					// 第三步, StartPush
					msg = "请求接受推送"
					log.Println(conn.RemoteAddr().String(), msg)
					
					if step != GetPushMessage {
						log.Println(conn.RemoteAddr().String(), msg, "失败, 原因: 非法请求:[", setpmsg[2], "]", "当前步骤应该为:[", setpmsg[step] ,"]")
						return
					}
					
					var ml *pushserver.MessageList
					if v, ok := clientMap.Get(token); ok {
						log.Println(conn.RemoteAddr().String(), "成功获取消息队列")
						ml = v.GetMessageList()
					} else {
						log.Println(conn.RemoteAddr().String(), "获取消息队列失败")
						return
					}
										
					log.Println(conn.RemoteAddr().String(), "开始推送消息")
					
					// http header
					var repheader = "HTTP/1.1 200 OK\r\n"
   					repheader += "Connection: keep-alive\r\n"
					repmsg = repheader
					repmsg += "\r\n"
					
					_, err = bf.Write([]byte(repmsg))
					if err != nil {
						log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
						return
					}
					err = bf.Flush()
					if err != nil {
						log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
						return			
					}
									
					var push_msg pushserver.Message	
					for i := 0;; i++ {
						// 检查消息队列
						count := ml.Count()
						if count > 0 {
							log.Println(conn.RemoteAddr().String(), "消息队列还有:", count, "条消息")
							if pop_msg, ok := ml.Pop(); ok {
								push_msg = pop_msg
								log.Println(conn.RemoteAddr().String(), "消息:", push_msg)
							} else {
								log.Println(conn.RemoteAddr().String(), "没有需要推送的消息")
								time.Sleep(100 * time.Millisecond)
								continue
							}
						} else {
							// 设置读取超时, 检查客户端在线
							conn.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
							_, err := bf.ReadByte()
							if err != nil {
								if err == io.EOF {
									log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
									return
								}
							}
							time.Sleep(100 * time.Millisecond)
							continue
						}
				
						// 组合推送信息
						const length = "Context-Length:"
						
						jsonAps := make(map[string]interface{})
						jsonAps["id"] = "1"
						jsonAps["alert"] = push_msg.Content
						jsonAps["badge"] = "100"
						jsonAps["type"] = "event"
						
						jsonBody := make(map[string]interface{})
						jsonBody["aps"] = jsonAps
					
						jsonData := make(map[string]interface{})
						jsonData["cmd"] = "pushmsg"
						jsonData["tokenid"] = token
						jsonData["body"] = jsonBody
						
						jsonBuf, _ := json.Marshal(jsonData)

						repmsg = fmt.Sprintf("Context-Length:%v\r\n\r\n%s", len(jsonBuf), jsonBuf)
						log.Println(conn.RemoteAddr().String(), "推送消息:", repmsg)
		
						// 发送推送信息
						_, err = bf.Write([]byte(repmsg))
						if err != nil {
							log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
							return
						}
						err = bf.Flush()
						if err != nil {
							log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
							log.Println(conn.RemoteAddr().String(), "缓冲区剩余数据[", bf.Writer.Buffered(), "]")
							return			
						}
		
						// 删除数据库记录
					}
				}
				default: {
					log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", "没有找到请求的命令")
					return
				}
			}
		}
	
    	// 开始下一次请求读取, 读取客户端请求 Header
  	  	if req, err = readRequest(conn, bf.Reader, timeout); err != nil {
  	  		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
  	  		return
		}
	}
}

/*
 *	获取在线客户端列表
 */
func showclientsHandleFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("来自:", r.RemoteAddr, r.Proto, "的请求显示在线客户端")
	
	var html string
	var clients string
	
	items := clientMap.Items()
	for key, _ := range items {
		clients += key
		clients += "<br>"
	}
	
	html = "<html><title>在线客户端列表</title><body>" + clients + "</body></html>"
	
	w.Write([]byte(html))
}
func InitSocket() {
	// 解除进程打开socket数量限制
//	var rlim syscall.Rlimit
//	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim)
//	if err != nil {
//		log.Fatal("获取socket限制失败: " + err.Error())
//	}
//	
//	log.Println("rlim.Cur:", strconv.FormatUint(rlim.Cur, 10))
//	rlim.Cur = 50000
//
//	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
//	if err != nil {
//		log.Fatal("设置socket限制失败: " + err.Error())
//	}
//	
//	log.Println("rlim.Cur:", strconv.FormatUint(rlim.Cur, 10))

}
func main() {
	InitSocket()
	var err error
		
	// 初始化
	clientMap = pushserver.NewClientMap()
	appMap = pushserver.NewAppMap()
	// 读取数据库应用
	db := pushserver.DefaultDataBase()
	db.LoadApp(appMap)
	
	// 开启检测客户端在线协程
	c := make(chan int)
	for i := 0; i < 1; i++ {
		go func() {
			index := <-c
			for j := 0; ;j++ {
				clientMap.Test(index)
				time.Sleep(30 * time.Second)
			}
		}()
		c<-i
	}
	
	// 创建http路由
	http.HandleFunc("/showclients", showclientsHandleFunc)
	http.HandleFunc("/", clientHandleFunc)
	
	// 启动http服务
	err = http.ListenAndServe(":8123", nil)
	if err != nil {
		log.Fatal("推送服务器启动失败, 原因:", err.Error())
	}
	
	log.Println("")
	log.Println("######### 推送服务器启动成功 #########")
	fmt.Println("")
}