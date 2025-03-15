package intf

type ConfigProvider interface {
	Provider
	Unmarshal(configVar any, key ...string) error
}
