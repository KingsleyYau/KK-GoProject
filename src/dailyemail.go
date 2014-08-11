package main 

import (
//	iconv "github.com/djimenez/iconv-go"
	"encoding/base64"
	"encoding/json"
//	"fmt"
	"flag"
	"io/ioutil"
	"net/smtp"
	"log"
	"os"
	"strings"
	"time"
//	"mime/multipart"
//	"mime"
)
const DailyEmailBoundary = "DailyEmailBoundary"

type DailyEmail struct {
	username string
	password string
	host string
	
	subject string
	toAddress []string
	ccAddress []string
	mailAddress []string
	
	content string
	attachment string
	
}

func (de *DailyEmail)LoadConfig() {
	log.Println("开始读取配置...")
	var err error
	
	file , err := os.Open("dailyemailconfig.json")
	if err != nil {
		log.Fatal("打开配置文件失败")
	}
	defer file.Close()
	
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("读取配置文件失败")
	}
	
	var config map[string]interface{}
	err = json.Unmarshal(buf, &config)
    if err != nil {
        log.Fatal("解析配置文件失败")
    }
    
    if value, ok := config["username"].(string); ok {
    	de.username = value
    } else {
    	log.Fatal("解析邮箱名失败")
    }
    if value, ok := config["password"].(string); ok {
    	de.password = value
    } else {
    	log.Fatal("解析用户密码失败")
    }
    if value, ok := config["host"].(string); ok {
    	de.host = value
    } else {
    	log.Fatal("解析smtp服务器失败")
    }
    if value, ok := config["subject"].(string); ok {
    	de.subject = value
    } else {
    	log.Fatal("解析标题失败")
    }
    if value, ok := config["toAddress"].(string); ok && value != "" {
    	de.toAddress = strings.Split(value, ",")
    } else {
    	log.Fatal("解析收件人失败")
    }
    
    if value, ok := config["ccAddress"].(string); ok && value != "" {
    	de.ccAddress = strings.Split(value, ",")
    }
    if value, ok := config["content"].(string); ok {
    	de.content = value
    	de.content = strings.Replace(de.content, "\\r\\n", "\r\n", -1)
    }
    if value, ok := config["attachment"].(string); ok {
    	de.attachment = value
    }
    
    log.Println("读取配置成功!")
}
func (de *DailyEmail)ToAddress() (string) {
	if len(de.toAddress) == 0 {
		return ""
	}
	toAddress := "To: "
	for _, value := range de.toAddress {
		toAddress += "<" + value + ">,"
	}
	toAddress += "\r\n"
	return toAddress
}
func (de *DailyEmail)CcAddress() (string) {	
	if len(de.ccAddress) == 0 {
		return ""
	}
	ccAddress := "Cc: "
	for _, value := range de.ccAddress {
		ccAddress += "<" + value + ">,"
	}
	ccAddress += "\r\n"
	return ccAddress
}
func (de *DailyEmail)Subject() (string) {
	subject := "subject: " + de.subject + "\r\n"
	return subject
}
func (de *DailyEmail)Header() (string) {
	header := "MIME-Version: 1.0\r\n"
	header += "Content-Type: multipart/mixed;charset=utf-8;"
	header +=" boundary=" + DailyEmailBoundary + "\r\n"
	header += "\r\n"
	return header
}
func (de *DailyEmail)Attachment() (string) {
	file , err := os.Open(de.attachment)
	if err != nil {
		log.Fatal("打开附件失败, 错误", err.Error())
	}
	defer file.Close()
	
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("读取附件失败, 错误", err.Error())
	}

//	fileName := "[" + time.Now().Format("2006-01-02") + "]" + de.attachment
	fileName := de.subject + "[" + time.Now().Format("2006-01-02") + "]" + ".xls"
	
	attachment := "--" + DailyEmailBoundary + "\r\n"
	attachment += "Content-Type: application/vnd.ms-excel;charset=utf-8;\r\n" 
//	attachment += "name=\"" + fileName + "\"" +"\r\n"
	attachment += "Content-Disposition: attachment;" 
	attachment += "filename=\"" + fileName + "\"" + "\r\n"
	
	attachment += "Content-Transfer-Encoding: base64\r\n\r\n"
	attachment += base64.StdEncoding.EncodeToString([]byte(buf))
	attachment += "\r\n"
	return attachment
}

func (de *DailyEmail)Content() (string) {
	msg := "--" + DailyEmailBoundary + "\r\n"
	msg += "Content-Type: text/plain;charset=utf-8;\r\n" 
	msg += "Content-Transfer-Encoding: base64\r\n\r\n"
	msg += base64.StdEncoding.EncodeToString([]byte(de.content))
	msg += "\r\n"
	return msg
}
func (de *DailyEmail)Send() {
	body := "" 
	body += de.ToAddress()
	body += de.CcAddress()
	body += de.Subject()
    body += de.Header()         
	body += de.Content()
	log.Println("邮件内容:\r\n", body)
	
	body += de.Attachment()
	body += "\r\n" + "--" + DailyEmailBoundary + "--"
	
	log.Println("开始发送邮件...")

	auth := smtp.PlainAuth(
				"",
                de.username,
                de.password,
                de.host,
                )
                
    err := smtp.SendMail(
    			de.host + ":25",
                auth,
                de.username,
                append(de.toAddress, de.ccAddress...),
                []byte(body),
        )
    if err != nil {
    	log.Fatal(err)
    } 
	log.Println("发送完成!")
}

var ArgHelp = flag.Bool("help", false, "帮助")

func main() {
	log.Println("")
	log.Println("############ 自动发送工作日报 v1.0 ############")
	log.Println("# Author:	Kingsley Yau")
	log.Println("# Date:	2014-07-28")
	log.Println("############################################")
	log.Println("")
	
	flag.Parse()
	
	if *ArgHelp {
		log.Fatal("当前目录下dailyemailconfig为配置文件, 格式为json, 填写好再次行程序即可")
	} 
	
	email := DailyEmail{}
	email.LoadConfig()
	email.Send()
}

