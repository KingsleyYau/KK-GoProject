package pushserver

import (
	"encoding/json"
//	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Message struct {
	Content 	string	`json:"content"`
	RecvTime	string
}
type MessageList struct {
	messageList []Message	`json:"messages"`
	lock *sync.RWMutex
}
func NewMessageList() *MessageList {
	return &MessageList {
        lock: new(sync.RWMutex),
        messageList:   make([]Message, 0, 1024),
    } 
}
func (ml *MessageList)Push(msg Message) {
	ml.lock.Lock()
    defer ml.lock.Unlock()
    
    ml.messageList = append(ml.messageList, msg)
}
func (ml *MessageList)PushList(msgs []Message) {
	ml.lock.Lock()
    defer ml.lock.Unlock()
    
    ml.messageList = append(ml.messageList, msgs...)
}
func (ml *MessageList) Pop() (Message, bool) {
	ml.lock.Lock()
    defer ml.lock.Unlock()
    var msg Message
    if len(ml.messageList) > 0 {
    	msg = ml.messageList[0]
    	ml.messageList = ml.messageList[1:len(ml.messageList)]
    	return msg, true
    }
    return msg, false
}
func (ml *MessageList) Count() int {
	ml.lock.RLock()
    defer ml.lock.RUnlock()
    
    return len(ml.messageList)
}
func (ml *MessageList) Data() []byte {
	ml.lock.Lock()
    defer ml.lock.Unlock()
    if data, err := json.Marshal(ml.messageList); err == nil {
    	return data
    }
    return nil
}

type Client struct {
	TokenId		string		`json:"tokenid"`
	AppVer		string		`json:"appver"`
	Model		string		`json:"model"`
	System		string		`json:"system"`
	Address		string		`json:"address"`	// 最后一次登录ip端口
	Online		string		`json:"online"`		// 最后一次登录时间
	Offline		string		`json:"offline"`	// 最后一次注销时间
	
	ml *MessageList
}
func NewClient(tokenid string) *Client {
	return &Client {
		TokenId:	tokenid,
        ml:   		NewMessageList(),
    } 
}
func (c *Client) InsertClientRecord() {
	db := DefaultDataBase()
	db.InsertClientRecord(c)
}
func (c *Client) UpdateClientInfo() {
	db := DefaultDataBase()
	db.UpdateClientInfo(c)
//	dir := "data/" + c.TokenId
//	os.Mkdir(dir, 0766)
//	fileName := dir + "/info.json"
//	fs, err := os.OpenFile(fileName, os.O_CREATE | os.O_TRUNC | os.O_RDWR, 0766)
//	if err != nil {
//		log.Println("打开客户端信息文件失败")
//		return
//	}
//	defer fs.Close()
//	
//	data, err := json.Marshal(c);
//	if err != nil {
//		log.Println("生成客户端信息失败")
//		return
//	}
//	
//	_, err = fs.Write(data)
//	if err != nil {
//		log.Println("更新客户端信息文件失败:", err.Error())
//		return
//	}
}
func (c *Client) DeleteMessage() {
	dir := "data/" + c.TokenId
	os.Mkdir(dir, 0766)
	fileName := dir + "/message.json"
	os.Remove(fileName)
}
func (c *Client) SaveMessage() {
	db := DefaultDataBase()
	db.SaveClientMessage(c)
//	dir := "data/" + c.TokenId
//	os.Mkdir(dir, 0766)
//	fileName := dir + "/message.json"
//	if data := c.ml.Data(); data != nil {
//		ioutil.WriteFile(fileName, data, 0766)
//	}
}
func (c *Client) LoadMessage() {
	db := DefaultDataBase()
	db.LoadClientMessage(c)
//	dir := "data/" + c.TokenId
//	os.Mkdir(dir, 0766)
//	fileName := dir + "/message.json"
//    if data, err := ioutil.ReadFile(fileName); err == nil {
//    	var msgs []Message
//    	json.Unmarshal(data, &msgs)
//    	c.ml.PushList(msgs)
//    }
//    c.DeleteMessage()
}
func (c *Client) GetMessageList() *MessageList {
	return c.ml
}


/*
 *	在线客户端列表
 */
type ClientMap struct {
	itemMap map[string]*Client // 根据token插入消息队列
	lock *sync.RWMutex
}	
func NewClientMap() *ClientMap {
    return &ClientMap {
        lock: new(sync.RWMutex),
        itemMap:   make(map[string]*Client),
    }
}
func (m *ClientMap) Test(i int) {
    m.lock.Lock()
    defer m.lock.Unlock()
    log.Println("########## " + "当前在线客户端:" + strconv.Itoa(len(m.itemMap)) + " #########")
}
func (m *ClientMap) Items() (map[string]*Client) {
	m.lock.RLock()
    defer m.lock.RUnlock()

    return m.itemMap
}
func (m *ClientMap) Count() int {
	m.lock.RLock()
    defer m.lock.RUnlock()
    
    return len(m.itemMap)
}
func (m *ClientMap) SendMessage(keys[]string, msg Message) {
    m.lock.Lock()
    defer m.lock.Unlock()
	for _, key := range keys {
		c := m.itemMap[key]
		if c != nil {
			log.Println("客户端在线, 插入消息到客户端[" + key + "]的消息队列")
			c.ml.Push(msg)
		} else {
			log.Println("客户端不在线, 插入消息到客户端[" + key + "]的数据库记录")
	    	// 保存未发送消息
	    	c := NewClient(key)
	    	c.ml.Push(msg)
	    	c.SaveMessage()
		}
	}
}
func (m *ClientMap) Get(key string) (*Client, bool) {
    m.lock.RLock()
    defer m.lock.RUnlock()
    var c *Client
    if v, ok := m.itemMap[key]; ok {
        return v, true
    }
    return c, false
}
func (m *ClientMap) Set(key string, c *Client) bool {
    m.lock.Lock()
    defer m.lock.Unlock()
    if v, ok := m.itemMap[key]; ok && v != nil {
    	// 已经在线
    	log.Println("客户端[" + key + "]已经在线")
    	return false
    } else {
    	// 新上线, 读取文件
    	log.Println("客户端[" + key + "]新上线")
    	c.Online = time.Now().Format("2006-01-02 15:04:05")
    	c.UpdateClientInfo()
    	c.LoadMessage()
    	m.itemMap[key] = c
    }
    return true
}
func (m *ClientMap) Check(key string) bool {
    m.lock.RLock()
    defer m.lock.RUnlock()
    if _, ok := m.itemMap[key]; !ok {
        return false
    }
    return true
}
func (m *ClientMap) Delete(key string) {
    m.lock.Lock()
    defer m.lock.Unlock()
  
    if v, ok := m.itemMap[key]; ok {
    	// 保存未发送消息到文件
    	v.SaveMessage()
    	// 更新下线时间
    	v.Offline = time.Now().Format("2006-01-02 15:04:05")
    	v.UpdateClientInfo()
    	// 生成在线记录
    	v.InsertClientRecord()
    }
    // 移除在线列表
    delete(m.itemMap, key)
}


/*
 *	应用列表		
 */
type AppMap struct {
	itemMap map[string]string // 根据appid插入appkey
	lock *sync.RWMutex
}
func NewAppMap() *AppMap {
    return &AppMap {
        lock: new(sync.RWMutex),
        itemMap:   make(map[string]string),
    }
}
func (m *AppMap) Get(key string) (string, bool) {
    m.lock.RLock()
    defer m.lock.RUnlock()
    if v, ok := m.itemMap[key]; ok {
        return v, true
    }
    return "", false
}
func (m *AppMap) Set(key string, value string) {
    m.lock.Lock()
    defer m.lock.Unlock()
    m.itemMap[key] = value
}
