package strtech

import (
	"os"
	"bufio"
	"fmt"
	"time"
	"encoding/json"
	"log"
)

type Logging struct {
	f *os.File
	w *bufio.Writer

	path string
}

func (ctx *Logging) Setup(path string) {
	ctx.path = path
}

/**
 */
func (ctx *Logging) createLogFile() error {
	if ctx.f != nil {
		return nil
	}

	if len(ctx.path) <= 0 {
		ctx.path = "."
	}

	t := time.Now()
	date := t.Format("2006-01-02")
	ft := fmt.Sprintf("%s/%s.log", ctx.path, date)
	fmt.Println("文件名:" + ft)

	ff := os.O_RDWR | os.O_CREATE | os.O_APPEND
	f, err := os.OpenFile(ft, ff, os.ModePerm)
	if err != nil {
		ft = fmt.Sprintf("%s/%s.log", ".", date)
		f, err = os.OpenFile(ft, ff, os.ModePerm)
		if err != nil {
			fmt.Printf("创建日志文件失败！！！")
			return err
		}
	}

	ctx.w = bufio.NewWriter(f)
	ctx.f = f

	return nil
}

func (ctx *Logging) Release() {
	ctx.f.Close()
	ctx.f = nil
	ctx.w = nil
}

func buildHead(level string) map[string]string {
	hn, err := os.Hostname()
	if err != nil {
		hn = "Unknown"
	}
	return map[string]string{"log.hostname": hn, "log.level": level}
}

func (ctx *Logging) log(level string, content string) {
	err := ctx.createLogFile()
	if err != nil {
		log.Fatal("创建log文件失败！")
		return
	}

	head := buildHead(level)
	head["log.content"] = content

	js, _ := json.Marshal(head)
	str := string(js) + string('\n')
	fmt.Println(content)
	ctx.w.WriteString(str)
	ctx.w.Flush()
}

func (ctx *Logging) Verbose(content string) {
	ctx.log("verbose", content)
}

func (ctx *Logging) Debug(content string) {
	ctx.log("debug", content)
}

func (ctx *Logging) Info(content string) {
	ctx.log("info", content)
}

func (ctx *Logging) Warn(content string) {
	ctx.log("warning", content)
}

func (ctx *Logging) Error(content string) {
	ctx.log("error", content)
}
