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

type ProviderName string

const (
	ProviderNameConfigViper      ProviderName = "config-viper"
	ProviderNameLoggerZerolog    ProviderName = "logger-zerolog"
	ProviderNameRedisRedigo      ProviderName = "redis-redigo"
	ProviderNameMysqlSqlBoiler   ProviderName = "mysql-sqlboiler"
	ProviderNameSqlite3SqlBoiler ProviderName = "sqlite3-sqlboiler"
	ProviderNameMqRabbitMq       ProviderName = "mq-rabbitmq"
	ProviderNameOssAliyun        ProviderName = "oss-aliyun"
)

// Capability 能力提供者
type Capability struct {
	Category ProviderCategory
	Name     ProviderName
	Module   fx.Option
}

// Provider 底层库能力提供者接口
type Provider interface {
	Init(args ...any) error // 初始化
}
