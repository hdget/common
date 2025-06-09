package intf

type ConfigProvider interface {
	Provider
	Unmarshal(configVar any, key ...string) error // 读取配置到变量configVar
}
