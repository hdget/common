package types

type DaprModuleKind int

const (
	DaprModuleKindUnknown DaprModuleKind = iota
	DaprModuleKindInvocation
	DaprModuleKindEvent
	DaprModuleKindDelayEvent
	DaprModuleKindHealth
)

type ParsedDaprHandler struct {
	ModuleKind  DaprModuleKind
	PkgPath     string            // 包路径
	Module      string            // 模块名
	Name        string            // 处理函数名
	Alias       string            // 处理函数别名
	Comments    []string          // 注释
	Annotations map[string]string // 注解
}
