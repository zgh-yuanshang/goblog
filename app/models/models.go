package models

import (
	apply "goblog/app"
	"goblog/pkg/types"
)

var (
	app    = apply.App
	router = app.Router
	db     = app.DB
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64
}

// GetStringID 获取 ID 的字符串格式
func (b BaseModel) GetStringID() string {
	return types.Uint64ToString(b.ID)
}
