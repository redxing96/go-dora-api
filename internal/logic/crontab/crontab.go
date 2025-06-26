package crontab

import (
	"context"
	"fmt"
	"go-dora-api/internal/model"
	"go-dora-api/internal/service"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sCrontab struct {
		cronObj *gcron.Cron
	}
)

func New() *sCrontab {
	c := gcron.New()
	obj := &sCrontab{
		cronObj: c,
	}
	return obj
}

func init() {
	service.RegisterCrontab(New())
}

// 初始化数据
func (s *sCrontab) InitData(ctx context.Context) error {
	// 添加一个每秒执行一次的任务
	_, err := s.cronObj.AddSingleton(ctx, "0/1 * * * * *", func(ctx context.Context) {
		// 打印当前时间
		fmt.Println(gtime.Now().Format("Y-m-d H:i:s.u"))
	}, "SecondsTimer")
	// 如果添加任务失败，则返回错误
	if err != nil {
		return err
	}
	// 添加任务成功，则返回nil
	return nil
}

// 启动sCrontab的定时任务
func (s *sCrontab) Start(ctx context.Context, name string) {
	// 启动定时任务
	s.cronObj.Start(name)
}

// 停止定时任务
func (s *sCrontab) Stop(ctx context.Context, name string) {
	// 调用cronObj的Stop方法，停止指定名称的定时任务
	s.cronObj.Stop(name)
}

// Remove函数用于从sCrontab中移除指定的定时任务
func (s *sCrontab) Remove(ctx context.Context, names []string) {
	// 遍历names中的每一个定时任务名称
	for _, name := range names {
		// 从s.cronObj中移除指定的定时任务
		s.cronObj.Remove(name)
	}
}

// 停止所有定时任务
func (s *sCrontab) StopAll(ctx context.Context) {
	// 停止定时任务
	s.cronObj.Stop()
}

// 根据传入的name参数，在s.cronObj中搜索对应的entry，如果找不到，则返回nil，否则返回一个model.CrontabSearchOutput结构体，包含entry的Name、Status、IsSingleton和RegisterTime属性
func (s *sCrontab) Search(ctx context.Context, name string) *model.CrontabSearchOutput {
	// 在s.cronObj中搜索name对应的entry
	entry := s.cronObj.Search(name)
	// 如果找不到，则返回nil
	if entry == nil {
		return nil
	}
	// 返回一个model.CrontabSearchOutput结构体，包含entry的Name、Status、IsSingleton和RegisterTime属性
	return &model.CrontabSearchOutput{
		Name:        entry.Name,
		Status:      entry.Status(),
		IsSingleton: entry.IsSingleton(),
		Time:        entry.RegisterTime,
	}
}

// All函数用于获取所有定时任务的信息
func (s *sCrontab) All(ctx context.Context) (output []model.CrontabSearchOutput) {
	// 初始化输出切片
	output = make([]model.CrontabSearchOutput, 0)
	// 获取所有定时任务
	entries := s.cronObj.Entries()
	// 遍历所有定时任务
	for _, entry := range entries {
		// 将定时任务的信息添加到输出切片中
		output = append(output, model.CrontabSearchOutput{
			Name:        entry.Name,
			Status:      entry.Status(),
			IsSingleton: entry.IsSingleton(),
			Time:        entry.RegisterTime,
		})
	}
	// 返回输出切片
	return
}
