package handler

import (
	"github.com/heroyf/cnccamp/go/lesson12/service0/core"
)

type Handler interface {
	GetHandleFunc() core.HandlerFunc
}

var (
	convertorRegistry map[string]Handler
)

// Register 注册处理器
func Register(name string, handler Handler) {
	convertorRegistry[name] = handler
}

// GetHandler 获取转换器
func GetHandler(name string) Handler {
	return convertorRegistry[name]
}

func init() {
	convertorRegistry = make(map[string]Handler)
}
