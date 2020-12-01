package service

import (
	"fmt"
	"time"
)

func Hello(name string) string {
	return fmt.Sprintf("hello %s, %s", name, time.Now().Format("2006-01-02 15:04:05"))
}
