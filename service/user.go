package service

import (
	"fmt"
	"time"
)

func Hello(name string) string {
	if name == "" {
		return "嘿，亲爱的，你还没告诉我你的名字呢"
	}
	return fmt.Sprintf("hello %s, %s", name, time.Now().Format("2006-01-02 15:04:05"))
}
