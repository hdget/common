package types

import "go.uber.org/fx"

type ProviderCategory int

const (
	ProviderCategoryUnknown ProviderCategory = iota
	ProviderCategoryConfig
	ProviderCategoryLogger
	ProviderCategoryDb
	ProviderCategoryRedis
	ProviderCategoryMq
	ProviderCategoryOss
)

// Capability 能力提供者
type Capability struct {
	Category ProviderCategory
	Name     string
	Module   fx.Option
}
