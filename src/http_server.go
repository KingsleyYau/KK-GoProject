package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
)
/*
 * 发送返回
 */
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

/*
 * 2.1.Facebook注册及登录（https post）（New）
 */
func member_facebooklogin(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("2.1.Facebook注册及登录（https post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":0,\"data\":{\"manid\":\"CM28171208\",\"email\":\"samson.fan@qpidnetwork.com\",\"firstname\":\"samson\",\"lastname\":\"fan\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"sid\":\"dcbd871252c965435b64260880b8b66c\",\"reg_step\":\"1\",\"premit\":\"Y\",\"country\":\"GB\",\"telephone_verify\":0,\"telephone_cc\":\"GB\",\"telephone\":\"123456\",\"photosend\":\"1\",\"photoreceived\":\"1\",\"sessionid\":\"58u634i9mi31hqkqlqlranaopo\",\"permission\":{\"ladyprofile\":0,\"livechat\":0,\"admirer\":0,\"bpemf\":0}},\"errno\":\"MBCE64002\",\"errmsg\":\"Facebook邮箱已注册但未绑定（需要提交密码）\",\"ext\":\"\",\"errdata\":{\"email\":\"samsom@qpidnetwork.com\",\"firstname\":\"samsson\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\"}}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 2.3.获取验证码（http post）（New）
 */
func member_checkcode(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("2.3.获取验证码（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	fileName := "u182.png"
    if bufSend, err := ioutil.ReadFile(fileName); err == nil {
        log.Println(conn.RemoteAddr().String(), "返回文件:", fileName)		
		sendRespone(bf.Writer, []byte(bufSend))
		log.Println("\n")
    }	
}

/*
 * 2.4.登录（https post）（New）
 */
func member_logincheck(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("2.4.登录（https post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"manid\":\"CM28171208\",\"email\":\"samson.fan@qpidnetwork.com\",\"firstname\":\"samson\",\"lastname\":\"fan\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"sid\":\"dcbd871252c965435b64260880b8b66c\",\"reg_step\":\"1\",\"premit\":\"Y\",\"country\":\"GB\",\"telephone_verify\":0,\"telephone_cc\":\"GB\",\"telephone\":\"123456\",\"photosend\":\"1\",\"photoreceived\":\"1\",\"sessionid\":\"58u634i9mi31hqkqlqlranaopo\",\"permission\":{\"ladyprofile\":0,\"livechat\":0,\"admirer\":0,\"bpemf\":0}},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 2.6.手机获取认证短信（http post）/ 2.8.固定电话获取认证短信（http post）
 */
func member_sms_get(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("2.6.手机获取认证短信（http post）/ 2.8.固定电话获取认证短信（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"MBCE47003\",\"errmsg\":\"The code has been sent. Please wait patiently.\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 2.7.手机短信认证（http post）
 */
func member_sms_verify(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("2.7.手机短信认证（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 2.9.固定电话短信认证（http post）
 */
func sms_verify_phonetype_2(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("2.9.固定电话短信认证（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":0,\"data\":\"\",\"errno\":\"MBCE48001\",\"errmsg\":\"Your verification code is incorrect. Please enter again.\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}


/*
 * 3.1.查询个人信息（http post）
 */
func member_myprofile(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("3.1.查询个人信息（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"manid\":\"CM31822615\",\"birthday\":\"1980-01-01\",\"firstname\":\"test\",\"lastname\":\"fan\",\"gender\":\"M\",\"country\":\"GB\",\"marry\":\"0\",\"email\":\"test.fan@qn.com\",\"weight\":\"0\",\"height\":\"0\",\"smoke\":\"0\",\"drink\":\"0\",\"language\":\"0\",\"religion\":\"0\",\"education\":\"0\",\"profession\":\"0\",\"ethnicity\":\"0\",\"income\":\"0\",\"children\":\"0\",\"headline\":\"\",\"jj\":\"\",\"address1\":\"\",\"address2\":\"\",\"city\":\"\",\"province\":\"\",\"zipcode\":\"\",\"telephone\":\"\",\"fax\":\"\",\"alternate_email\":\"\",\"money\":\"0.00\",\"photo\":\"0\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"integral\":\"10\",\"rs_status\":\"2\",\"rs_content\":\"\",\"mobile_cc\":\"AF\",\"mobile\":\"12345678910\",\"landline_cc\":\"AF\",\"landline_ac\":\"011\",\"landline\":\"\"},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 3.2.修改个人信息（http post）
 */
func member_updatepro(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("3.2.修改个人信息（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"jj_result\":-1},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 3.5.上传头像（http post）
 */
func member_uploadphoto(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("3.5.上传头像（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"dataCount\":1,\"datalist\":[{\"attachnum\":\"1\",\"virtual_gifts\":0,\"womanid\":\"P311508\",\"id\":\"19371228\",\"readflag\":\"1\",\"rflag\":\"0\",\"fflag\":\"0\",\"pflag\":\"0\",\"firstname\":\"Visa\",\"lastname\":\"Li\",\"weight\":\"47\",\"height\":\"162\",\"country\":\"China\",\"province\":\"HuNan\",\"age\":25,\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"sendTime\":\"Jan 20, 2015\",\"intro\":\"test,intro\"}],\"pageIndex\":1,\"pageSize\":30},\"errno\":\"\",\"errmsg\":\"\",\"ext\":null}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 5.1.获取匹配女士条件（http post）
 */
func member_get_criteria(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("5.1.获取匹配女士条件（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"age1\":\"18\",\"age2\":\"99\",\"m_marry\":\"0\",\"m_children\":\"0\",\"m_education\":\"0\"},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 5.2.保存匹配女士条件（http post）
 */
func member_save_criteria(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("5.2.保存匹配女士条件（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 5.3.条件查询女士列表（http post）（New）
 */
func member_search(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("5.3.条件查询女士列表（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"dataCount\":1,\"datalist\":[{\"womanid\":\"P211211\",\"firstname\":\"HaiYan\",\"age\":50,\"weight\":\"52\",\"height\":\"164\",\"country\":\"China\",\"province\":\"Sichuan\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\"}],\"pageIndex\":1,\"pageSize\":30,\"maninfo\":{\"age1\":\"18\",\"age2\":\"99\",\"m_marry\":\"0\",\"m_country\":\"\",\"m_children\":\"0\",\"m_education\":\"0\"}},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 5.8.获取最近联系人列表（http post）（New）
 */
func member_recentcontact(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("5.8.获取最近联系人列表（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"datalist\":[{\"womanid\":\"GZA881\",\"firstname\":\"Jennifer\",\"age\":23,\"photoURL\":\"http://192.168.88.140/Share/213.png\",\"videocount\":2,\"isfavorite\":1,\"lasttime\":1428471899}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 5.9.查询女士标签列表（http post）（New）
 */
func member_signlist(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("5.9.查询女士标签列表（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"datalist\":[{\"signid\":\"12345\",\"name\":\"Cute\",\"color\":\"#FF0000\",\"issigned\":1}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 5.10.提交女士标签（http post）（New）
 */
func member_uploadsign(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("5.10.提交女士标签（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}


/*
 * 6.7.付费获取私密照片（http get）（New）
 * 6.8.获取对方私密照片（http get）（New）
 */
func livechat_setstatus(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("6.7.付费获取私密照片（http get）（New）/ 6.8.获取对方私密照片（http get）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	err = r.ParseForm()
	if err != nil {
        return
    }
    
    action := r.FormValue("action")
    log.Println(conn.RemoteAddr().String(), "action :", action)	

	if action == "man_get_photolist" {
	 	log.Println(conn.RemoteAddr().String(), "6.7.付费获取私密照片（http get）（New）")	
		bufSend := "<?xml version=\"1.0\" encoding=\"utf-8\"?><ROOT><result><status>1</status><errcode></errcode><errmsg></errmsg><message></message></result></ROOT>"
		log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
		sendRespone(bf.Writer, []byte(bufSend))
		log.Println("\n")
	} else if action == "load_private_photo" {
		log.Println(conn.RemoteAddr().String(), "6.8.获取对方私密照片（http get）（New）")	
		fileName := "u213.png"
	    if bufSend, err := ioutil.ReadFile(fileName); err == nil {
	        log.Println(conn.RemoteAddr().String(), "返回文件:", fileName)		
			sendRespone(bf.Writer, []byte(bufSend))
			log.Println("\n")
	    }
	}

}



/*
 * 7.1.查询收件箱列表（http post）
 */
func emf_inboxlist(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.1.查询收件箱列表（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"dataCount\":1,\"datalist\":[{\"attachnum\":\"1\",\"virtual_gifts\":0,\"womanid\":\"P311508\",\"id\":\"19371228\",\"readflag\":\"1\",\"rflag\":\"0\",\"fflag\":\"0\",\"pflag\":\"0\",\"firstname\":\"Visa\",\"lastname\":\"Li\",\"weight\":\"47\",\"height\":\"162\",\"country\":\"China\",\"province\":\"HuNan\",\"age\":25,\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"sendTime\":\"Jan 20, 2015\",\"intro\":\"test,intro\"}],\"pageIndex\":1,\"pageSize\":30},\"errno\":\"\",\"errmsg\":\"\",\"ext\":null}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.2.查询已收邮件详情（http post）（New）
 */
func emf_inboxmsg(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.2.查询已收邮件详情（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"id\":\"19371228\",\"womanid\":\"P311508\",\"firstname\":\"Visa\",\"lastname\":\"Li\",\"weight\":\"47\",\"height\":\"162\",\"country\":\"China\",\"province\":\"HuNan\",\"age\":25,\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"body\":\";';';';;'\u91cd\u8981\u63d0\u9192\uff1a<br />\r\n* \u8acb\u7559\u610f\uff0c\u9019\u662f\u5973\u58eb\u7d66sdfsa\u7684\u7b2c\u4e00\u5c01\u56de\u4fe1\u3002\u7232\u4e86\u80fd\u7d66\u7537\u58eb\u7559\u4e0b\u826f\u597d\u7684\u5370\u8c61\uff0c\u589e\u9032\u96d9\u65b9\u7684\u611f\u60c5\u4ea4\u6d41\uff0c\u8acb\u8a8d\u771f\u56de\u5fa9\u9019\u5c01\u4fe1\u3002 \u91cd\u8981\u63d0\u9192\uff1a<br />\r\n* \u8acb\u7559\u610f\uff0c\u9019\u662f\u5973\u58eb\u7d66sdfsa\u7684\u7b2c\u4e00\u5c01\u56de\u4fe1\u3002\u7232\u4e86\u80fd\u7d66\u7537\u58eb\u7559\u4e0b\u826f\u597d\u7684\u5370\u8c61\uff0c\u589e\u9032\u96d9\u65b9\u7684\u611f\u60c5\u4ea4\u6d41\uff0c\u8acb\u8a8d\u771f\u56de\u5fa9\u9019\u5c01\u4fe1\u3002 \u91cd\u8981\u63d0\u9192\uff1a<br />\r\n* \u8acb\u7559\u610f\uff0c\u9019\u662f\u5973\u58eb\u7d66sdfsa\u7684\u7b2c\u4e00\u5c01\u56de\u4fe1\u3002\u7232\u4e86\u80fd\u7d66\u7537\u58eb\u7559\u4e0b\u826f\u597d\u7684\u5370\u8c61\uff0c\u589e\u9032\u96d9\u65b9\u7684\u611f\u60c5\u4ea4\u6d41\uff0c\u8acb\u8a8d\u771f\u56de\u5fa9\u9019\u5c01\u4fe1\u3002 \",\"notetoman\":\"\",\"photosURL\":\"http://192.168.88.140/Share/u213.png\",\"vg_id\":\"VG0008\",\"sendTime\":\"Jan 20, 2015\",\"private_photos\":[{\"send_id\":\"1\",\"photo_id\":\"2001\",\"photo_fee\":1,\"photo_desc\":\"photo description\"}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.3.查询发件箱列表（http post）
 */
func emf_outboxlist(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.3.查询发件箱列表（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"dataCount\":9,\"datalist\":[{\"attachnum\":\"1\",\"virtual_gifts\":0,\"womanid\":\"P844666\",\"id\":\"19371274\",\"progress\":\"0\",\"firstname\":\"Summer\",\"lastname\":\"Yang\",\"weight\":\"48\",\"height\":\"160\",\"country\":\"China\",\"province\":\"HuNan\",\"age\":23,\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"sendTime\":\"Mar 11, 2015\",\"intro\":\"emf mail intro\"}],\"pageIndex\":1,\"pageSize\":15},\"errno\":\"\",\"errmsg\":\"\",\"ext\":null}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.4.查询已发邮件详细（http post）
 */
func emf_outboxmsg(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.4.查询已发邮件详细（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"id\":\"19371263\",\"womanid\":\"P2180824\",\"firstname\":\"Jasmine\",\"lastname\":\"Tang\",\"weight\":\"48\",\"height\":\"160\",\"country\":\"China\",\"province\":\"Guangdong\",\"age\":46,\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"photosURL\":\"http://192.168.88.140/Share/u213.png\",\"vg_id\":\"VG0008\",\"content\":\"test\",\"sendTime\":\"Feb 28, 2015\",\"private_photos\":[{\"send_id\":\"1\",\"photo_id\":\"2001\",\"photo_fee\":1,\"photo_desc\":\"photo description\"}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.5.查询收件箱某状态邮件数量（http post）
 */
func emf_msgtotal(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.5.查询收件箱某状态邮件数量（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"datalist\":[{\"womanid\":\"123456\",\"image\":\"http://192.168.88.140/Share/u213.png\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"firstname\":\"Junnifer\",\"country\":\"CN\",\"age\":18},{\"womanid\":\"123457\",\"image\":\"http://192.168.88.140/Share/u381.png\",\"photoURL\":\"http://192.168.88.140/Share/u381.png\",\"firstname\":\"Cherry\",\"country\":\"US\",\"age\":20}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.6.发送邮件（http post）
 */
func emf_sendmsg(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.6.发送邮件（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"messageid\":19371263,\"sendTime\":1425093349},\"errno\":\"\",\"errmsg\":\"\",\"errdata\":{\"money\":\"1.0\"},\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.8.上传邮件附件（http post）（New）
 */
func emf_uploadattach(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.8.上传邮件附件（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.15.男士付费获取EMF私密照片（http post）（New）
 */
func emf_inbox_photo_fee(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.15.男士付费获取EMF私密照片（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 7.16.获取对方或自己的EMF私密照片（http post）（New）
 */
func emf_private_photo_view(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("7.16.获取对方或自己的EMF私密照片（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	fileName := "u213.png"
    if bufSend, err := ioutil.ReadFile(fileName); err == nil {
        log.Println(conn.RemoteAddr().String(), "返回文件:", fileName)		
		sendRespone(bf.Writer, []byte(bufSend))
		log.Println("\n")
    }
}

/*
 * 9.1.查询浮窗广告（http post）（New）
 */
func advert_mainadvert(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("9.1.查询浮窗广告（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"advertId\":\"12345\",\"image\":\"http://192.168.88.140/Share/213.png\",\"width\":320,\"height\":240,\"adurl\":\"http://www.baidu.com\",\"opentype\":1,\"isshow\":1,\"valid\":1429687937},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 9.2.查询女士列表广告（http post）（New）
 */
func advert_womanlistadvert(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("9.2.查询女士列表广告（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"advertId\":\"12345\",\"image\":\"http://192.168.88.140/Share/u213.png\",\"width\":320,\"height\":240,\"adurl\":\"http://www.baidu.com\",\"opentype\":1,\"isshow\":1,\"valid\":1429687937},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 9.3.查询Push广告（http post）（New）
 */
func advert_pushadvert(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("9.3.查询Push广告（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"datalist\":[{\"pushId\":\"123456\",\"message\":\"test\",\"adurl\":\"http://www.baidu.com\",\"opentype\":1}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}


/*
 * 10.1.查询女士图片列表（http post）（New）
 */
func quickmatch_ladylist(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("10.1.查询女士图片列表（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"datalist\":[{\"womanid\":\"123456\",\"image\":\"http://192.168.88.140/Share/u213.png\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\",\"firstname\":\"Junnifer\",\"country\":\"CN\",\"age\":18},{\"womanid\":\"123457\",\"image\":\"http://192.168.88.140/Share/u381.png\",\"photoURL\":\"http://192.168.88.140/Share/u381.png\",\"firstname\":\"Cherry\",\"country\":\"US\",\"age\":20}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))
	
	log.Println("\n")
}

/*
 * 10.2.提交已标记的女士（http post）（New）
 */
func quickmatch_uploadlady(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("10.2.提交已标记的女士（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

/*
 * 10.3.查询已标记like的女士列表（http post）（New）
 */
func quickmatch_likeladylist(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("10.3.查询已标记like的女士列表（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"dataCount\":1,\"datalist\":[{\"womanid\":\"123456\",\"firstname\":\"Jasmine\",\"age\":46,\"country\":\"CN\",\"photoURL\":\"http://192.168.88.140/Share/u213.png\"}],\"pageIndex\":1,\"pageSize\":15},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

/*
 * 10.4.删除已标记like的女士（http post）（New）
 */
func quickmatch_removelikelady(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("10.4.删除已标记like的女士（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

/*
 * 11.1.获取Love Call列表接口（http post）（New）
 */
func lovecall_lovecalllist(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("11.1.获取Love Call列表接口（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"pageIndex\":1,\"pageSize\":20,\"dataCount\":100,\"datalist\":[{\"orderid\":\"123456\",\"womanid\":\"123456\",\"image\":\"http://192.168.88.140/Share/u213.png\",\"firstname\":\"Junnifer\",\"begintime\":1428471899,\"endtime\":1428471899,\"needtr\":1,\"isconfirm\":1,\"confirmmsg\":\"hi\",\"callid\":\"123456\",\"centerid\":\"123456\"}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

/*
 * 11.2.确定Love Call接口（http post）（New）
 */
func lovecall_confirmrequest(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("11.2.确定Love Call接口（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":1,\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}


/*
 * 12.2.男士会员统计（http post）
 */
func member_get_count(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("12.2.男士会员统计（http post）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"money\":\"979.00\",\"coupon\":2,\"integral\":1,\"regstep\":3,\"allowalbum\":1,\"admirer_ur\":\"10\",\"isonline\":1,\"isfavorite\":1,\"videofav\":1},\"errno\":{\"regstep\":null},\"errmsg\":{\"regstep\":null},\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

/*
 * 12.5.检查客户端更新（http post）（New）
 */
func other_versioncheck(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("12.5.检查客户端更新（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"versionCode\":281,\"versionName\":\"V2.6\",\"versionDesc\":\"\",\"url\":\"http://demo-m.chnlove.com/uploadfile/apk/qpidnetwork_man_v2.6_demo.apk\",\"app_ios_url\":\"https://itunes.apple.com/us/app/qpid-network/id570924081?ls=1&mt=8\",\"pubtime\":\"2012-05-29\",\"checktime\":1425112925},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

/*
 * 12.8.查询站点当前在线人数（http post）（New）
 */
func other_onlinecount(w http.ResponseWriter, r *http.Request) {
	log.Println("###########################")
	log.Println("12.8.查询站点当前在线人数（http post）（New）")
	log.Println(r.RemoteAddr, r.Proto, "的新连接", r.URL.Path)
	w.Header().Set("Connection", "close")
	
	// 械劫
    var conn net.Conn
	var bf *bufio.ReadWriter
	var err error
	
	if wh, ok := w.(http.Hijacker); ok {
		conn, bf, err = wh.Hijack() 
		if err != nil {
			log.Println(r.RemoteAddr, "械劫连接失败, 原因:", err.Error())
			return
		}
		 
		defer log.Println("关闭连接:", r.RemoteAddr, "\n")
		defer conn.Close()
	}

	// 读取 http body
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "连接已经断开, 原因:", err.Error())
	    return
	}
	r.Body.Close()
	log.Println(conn.RemoteAddr().String(), "Http请求 Body:\n", string(buf))
	
	bufSend := "{\"result\":1,\"data\":{\"datalist\":[{\"siteid\":\"1\",\"onlinecount\":10000}]},\"errno\":\"\",\"errmsg\":\"\",\"ext\":\"\"}"
	log.Println(conn.RemoteAddr().String(), "返回数据:\n", bufSend)			
	sendRespone(bf.Writer, []byte(bufSend))

	log.Println("\n")
}

func main() {
	var err error

	// 创建http路由
	// 2.认证登录
	http.HandleFunc("/member/facebooklogin", member_facebooklogin)
	http.HandleFunc("/member/checkcode", member_checkcode)
	http.HandleFunc("/member/logincheck", member_logincheck)
	
	http.HandleFunc("/member/sms_get", member_sms_get)
	http.HandleFunc("/member/sms_verify", member_sms_verify)
	http.HandleFunc("/sms_verify/phonetype/2/", sms_verify_phonetype_2)
			
	// 3.个人信息
	http.HandleFunc("/member/myprofile", member_myprofile)
	http.HandleFunc("/member/uploadphoto", member_uploadphoto)
	http.HandleFunc("/member/updatepro", member_updatepro)
		
	// 5.女士信息
	http.HandleFunc("/member/get_criteria", member_get_criteria)
	http.HandleFunc("/member/save_criteria", member_save_criteria)	
	http.HandleFunc("/member/search", member_search)
	http.HandleFunc("/member/recentcontact", member_recentcontact)
	http.HandleFunc("/member/signlist", member_signlist)
	http.HandleFunc("/member/uploadsign", member_uploadsign)
	
	// 6.Live Chat
	http.HandleFunc("/livechat/setstatus.php", livechat_setstatus)
	
	// 7.EMF
	http.HandleFunc("/emf/inboxlist", emf_inboxlist)
	http.HandleFunc("/emf/inboxmsg", emf_inboxmsg)
	http.HandleFunc("/emf/outboxlist", emf_outboxlist)
	http.HandleFunc("/emf/outboxmsg", emf_outboxmsg)
	http.HandleFunc("/emf/msgtotal", emf_msgtotal)
	http.HandleFunc("/emf/sendmsg", emf_sendmsg)
	http.HandleFunc("/emf/uploadattach", emf_uploadattach)
	http.HandleFunc("/emf/inbox_photo_fee", emf_inbox_photo_fee)
	http.HandleFunc("/emf/private_photo_view", emf_private_photo_view)
	
	// 9.广告
	http.HandleFunc("/advert/mainadvert", advert_mainadvert)
	http.HandleFunc("/advert/womanlistadvert", advert_womanlistadvert)
	http.HandleFunc("/advert/pushadvert", advert_pushadvert)
	
	// 10.Quick Match（New）
	http.HandleFunc("/quickmatch/ladylist", quickmatch_ladylist)
	http.HandleFunc("/quickmatch/likeladylist", quickmatch_likeladylist)
	http.HandleFunc("/quickmatch/uploadlady", quickmatch_uploadlady)
	http.HandleFunc("/quickmatch/removelikelady", quickmatch_removelikelady)
	
	// 11.Love Call（New）
	http.HandleFunc("/lovecall/lovecalllist", lovecall_lovecalllist)
	http.HandleFunc("/lovecall/confirmrequest", lovecall_confirmrequest)
	
	// 12.其它协议
	http.HandleFunc("/member/get_count", member_get_count)
	http.HandleFunc("/other/versioncheck", other_versioncheck)
	http.HandleFunc("/other/onlinecount", other_onlinecount)
	
	// 启动http服务
	err = http.ListenAndServe(":81", nil)
	if err != nil {
		log.Fatal("http test server 启动失败, 原因:", err.Error())
	}
	
	log.Println("")
	log.Println("######### http test server 启动成功 #########")
	fmt.Println("")
}