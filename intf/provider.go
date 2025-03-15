package intf

import "github.com/hdget/common/types"

// Provider 底层库能力提供者接口
type Provider interface {
	GetCategory() types.ProviderCategory // 获取能力提供者分类
	GetName() string                     // 获取能力提供者名字
}
