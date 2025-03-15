package intf

import "github.com/hdget/common/types"

// Provider 底层库能力提供者接口
type Provider interface {
	GetCapability() types.Capability // 获取能力
}
