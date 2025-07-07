package intf

type HookFunction func() error

type AppServer interface {
	Start() error                               // 开始
	HookPreStart(fns ...HookFunction) AppServer // 添加启动前的钩子函数
	HookPreStop(fns ...HookFunction) AppServer  // 添加停止前的钩子函数
}
