package pushserver

import (
	"database/sql"
	_ "mysql" // 仅为了调用mysql包里面的init()
	"log"
)

var defaultDataBase *PushDataBase

func DefaultDataBase() *PushDataBase {
	if defaultDataBase == nil {
		defaultDataBase = NewPushDataBase()
	}
	return defaultDataBase
} 

type PushDataBase struct {
	username	string
	password	string
	host		string
	dbpool		chan *sql.DB
	maxConn		int
}
func (d *PushDataBase) getConn() (*sql.DB) {
	var conn *sql.DB = nil
	if len(d.dbpool) > 0 {
		conn = <-d.dbpool
	} else {
		conn = d.newConn()
	}
	return conn
}
func (d *PushDataBase) putConn(conn *sql.DB) {
	if len(d.dbpool) == d.maxConn {
		conn.Close()
		return
	}
	d.dbpool<-conn
}
func (d *PushDataBase) newConn() (*sql.DB) {
	conn, err := sql.Open("mysql", "root:!QAZ2wsx@tcp(localhost:3306)/push")
	if err != nil {
        log.Println(err.Error())
        return nil
    }
	return conn
}
func NewPushDataBase() *PushDataBase {
	return &PushDataBase{
			username:	"root",
			password:	"!QAZ2wsx",
			host:		"localhost:3306",
			dbpool:		make(chan *sql.DB, 20),
			maxConn:	20,
			}
}
func (d *PushDataBase) LoadClientMessage(c *Client) bool  {
	conn := d.getConn()
	if conn == nil {
		return false
	}
	
	// 查询离线消息
	rows, err := conn.Query("select message from push_message where tokenid='" + c.TokenId + "'") 
	if err != nil {
        log.Println(err.Error())
        return false
    }
    defer rows.Close()
    
    for rows.Next() { 
    	var msg Message
    	rErr := rows.Scan(&msg.Content)
        if rErr != nil {
        	log.Println("读取离线消息失败:", rErr.Error())
        	return false
        }
        c.ml.Push(msg)
    }
    
    // 清除离线消息
    sqlString := "delete from push_message where tokenid = '" + c.TokenId + "'"
	log.Println("sql:", sqlString)
	_, err = conn.Exec(sqlString)
	if err != nil {
		log.Println("清除离线消息失败:", err.Error())
		return false
	}
    return true
}
func (d *PushDataBase) SaveClientMessage(c *Client) bool {
	conn := d.getConn()
	if conn == nil {
		return false
	}
	
	for c.ml.Count() > 0 {
		msg, b := c.ml.Pop()
		if b {
			sqlString := "insert into push_message (tokenid, message, type, recvtime) values('" + c.TokenId + "','" + msg.Content + "'," + "1" + ",'" + msg.RecvTime + "')"
			log.Println("sql:", sqlString)
			_, err := conn.Exec(sqlString)
			if err != nil {
				log.Println("插入离线消息失败:", err.Error())
				return false
			}
		} else {
			return false
		}
	}
	return true
}
func (d *PushDataBase) UpdateClientInfo(c *Client) bool {
	conn := d.getConn()
	if conn == nil {
		return false
	}
	
	_, err := conn.Exec("insert into client_info (tokenid, appver, model, system, address, online, offline) values('" + c.TokenId + "','" + c.AppVer + "','" + c.Model + "','" + c.System + "','" + c.Address + "','" + c.Online + "','" + c.Offline + "')")
	if err != nil {
//		log.Println("插入客户端失败:", err.Error())
		
		// 插入失败, 尝试updapte
		sqlString := "update client_info set appver = '" + c.AppVer + "'," + "model = '" + c.Model + "'," + "system = '" + c.System + "'," + "address = '" + c.Address + "'," + "online = '" + c.Online + "',"+ "offline = '" + c.Offline + "'"
		log.Println("sql:", sqlString)
		_, err := conn.Exec(sqlString) 
		if err != nil {
			panic(err.Error())
			log.Println("更新客户端失败:", err.Error())
			return false
		}
	}
	
	return true
}
func (d *PushDataBase) InsertClientRecord(c *Client) bool {
	conn := d.getConn()
	if conn == nil {
		return false
	}
	
	_, err := conn.Exec("insert into client_online_record (tokenid, address, online, offline) values('" + c.TokenId + "','" + c.Address + "','" + c.Online + "','" + c.Offline + "')")
	if err != nil {
		log.Println("插入客户端在线记录失败:", err.Error())
		return false
	}
	
	return true
}
func (d *PushDataBase) LoadApp(m *AppMap) bool {
	conn := d.getConn()
	if conn == nil {
		return false
	}
	
	// 查询应用
	rows, err := conn.Query("select appid, appkey from push_identificatoin") 
	if err != nil {
        log.Println(err.Error())
        return false
    }
    defer rows.Close()
    
    for rows.Next() { 
    	var appid string
    	var appkey string
    	rErr := rows.Scan(&appid, &appkey)
        if rErr != nil {
        	log.Println("查询应用失败:", rErr.Error())
        	return false
        }
        m.Set(appid, appkey)
    }
    
    return true
}