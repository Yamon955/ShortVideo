package main

import (
	"trpc.group/trpc-go/trpc-go/codec"
)

const (
	SerializationTypeFormData = 131 // trpc 框架 formdata 格式序列化对应的键值
)

// FormDataSerialization form-data 格式数据序列化处理器
type FormDataSerialization struct {
	codec.JSONPBSerialization
}

// Unmarshal multipart/form-data 类型解码方法, 重写 Unmarshal 方法，对于 form-data 数据无需反序列化
func (s *FormDataSerialization) Unmarshal(in []byte, body interface{}) error {

	return nil
}
