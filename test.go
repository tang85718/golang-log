package main

import (
	"time"
	"os"
	"log"
	"encoding/json"
	"bufio"
)

//func main() {
//	host, err := os.Hostname()
//	if err != nil {
//		fmt.Println(err)
//		host = "Baba"
//	}
//
//	msg := fmt.Sprintf("{ \"host\":\"%s\", \"contex\":\"fuck some things\" }", host)
//
//	for {
//		l := logstash.New("logstash", 5045, 5)
//		conn, err := l.Connect()
//		if err != nil {
//			fmt.Println(err)
//			time.Sleep(time.Second)
//			continue
//		}
//
//		err = l.Writeln(msg)
//		if err != nil {
//			fmt.Println(err)
//			time.Sleep(time.Second)
//			continue
//		}
//
//		conn.Close()
//		time.Sleep(time.Second)
//	}
//}

func main() {
	flag := os.O_RDWR | os.O_CREATE | os.O_APPEND
	f, err := os.OpenFile("/opt/log/json.log", flag, 0777)
	if err != nil {
		log.Printf("创建日志文件失败！！！")
		return
	}

	w := bufio.NewWriter(f)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Baba"
	}
	index := 0
	for {
		d := map[string]string{"name":hostname, "content":"fuck log system"}
		js, _:= json.Marshal(d)
		str := string(js) + string('\n')
		log.Println(str)
		w.WriteString(str)
		w.Flush()

		time.Sleep(time.Second)
		index ++
	}

	f.Close()
}
