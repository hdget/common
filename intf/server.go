package intf

type HookPoint int

const (
	HookPointUnknown     HookPoint = iota
	HookPointBeforeStart           // 启动前执行
	HookPointBeforeStop            // 停止前执行
)

type HookFunction func() error

type AppServer interface {
	Start() error                                           // 开始
	Stop(forced ...bool) error                              // 默认为优雅关闭, 如果forced为true, 则强制关闭
	AddHook(point HookPoint, fns ...HookFunction) AppServer // 添加钩子函数
}
