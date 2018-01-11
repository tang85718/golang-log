package main

import (
	"./strtech"
	"time"
	"math/rand"
)

func main() {

	log := new(strtech.Logging)
	log.Setup("")

	content := [...]string{
		"成就天下无双之力",
		"流放之王",
		"暗黑破坏神迪亚波罗",
		"剑仙李逍遥",
		"喜剧之王",
		"破烂指针",
		"僵尸妹妹爱大吊",
		"僵尸哥哥爱美丽",
		"无厘小仙女"}

	for {
		size := len(content)
		r := rand.Intn(size)
		str := content[r]
		log.Debug(str)
		time.Sleep(time.Second)
	}
}
