package crontask

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/threading"
)

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) process() {
	cn := cron.New()

	cn.AddFunc("@every 5m", func() {
		fmt.Println("cron working")
	})

	cn.Start()
}

func (t *Task) Start() {
	threading.GoSafe(t.process)
}
