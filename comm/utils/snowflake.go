package utils

import (
	"errors"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake *sonyflake.Sonyflake
	startTime = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
)

// Init 使用当前机器 ID 初始化雪花算法生成器
func Init(machineID uint16) error {
	// 配置 Sonyflake
	settings := sonyflake.Settings{
		MachineID: func() (uint16, error) {
			return machineID, nil
		},
		StartTime: startTime,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	if sonyFlake == nil {
		return errors.New("sonyflake not created")
	}
	return nil
}

// GenID 雪花算法生成唯一 ID
func GenID() uint64 {
	id, _ := sonyFlake.NextID()
	return id
}
