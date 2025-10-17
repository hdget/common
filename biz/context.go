package biz

import (
	"context"
	"strconv"
	"strings"

	"github.com/hdget/common/meta"
	"google.golang.org/grpc/metadata"
)

type Context interface {
	MetaData() meta.MetaData // 元数据
	Transactor() transactor  // 数据库Trasactor相关
	Tid() int64              // 获取租户ID
	Uid() int64              // 获取用户ID
	AppId() string           // 获取应用ID
	RoleIds() []int64        // 获取角色ID列表
}

type contextImpl struct {
	metadata   meta.MetaData
	transactor transactor
	tid        int64   // 缓存租户ID提升效率
	uid        int64   // 缓存用户ID提升效率
	appId      string  // 缓存应用ID提升效率
	roleIds    []int64 // 缓存角色列表提升效率
}

func NewContext(kvs ...string) Context {
	return &contextImpl{
		metadata:   meta.New(kvs...),
		transactor: newTransactor(),
	}
}

// NewFromIncomingGrpcContext GRPC context => biz.Context
func NewFromIncomingGrpcContext(ctx context.Context) Context {
	c := &contextImpl{
		metadata:   meta.New(),
		transactor: newTransactor(),
	}

	// 尝试从grpc context中获取meta信息
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return c
	}

	for key, values := range md {
		switch key {
		case meta.KeyTid, meta.KeyUid: // int64
			val, _ := strconv.ParseInt(values[0], 10, 64)
			c.metadata.Set(key, val)
		case meta.KeyRoleIds:
			strIds := strings.Split(values[0], ",")
			val := make([]int64, len(strIds))
			for i, s := range strIds {
				val[i], _ = strconv.ParseInt(s, 10, 64)
			}
			c.metadata.Set(key, val)
		default:
			c.metadata.Set(key, values[0])
		}
	}
	return c
}

// NewOutgoingGrpcContext biz.Context => GRPC context
func NewOutgoingGrpcContext(ctx Context) context.Context {
	return metadata.NewOutgoingContext(context.Background(), ctx.MetaData().AsGRPCMetaData())
}

// MetaData 获取元数据
func (c *contextImpl) MetaData() meta.MetaData {
	return c.metadata
}

// Transactor 获取DB Transactor
func (c *contextImpl) Transactor() transactor {
	return c.transactor
}

func (c *contextImpl) Tid() int64 {
	if c.tid == 0 {
		c.tid = c.metadata.GetInt64(meta.KeyTid)
	}
	return c.tid
}

func (c *contextImpl) Uid() int64 {
	if c.uid == 0 {
		c.uid = c.metadata.GetInt64(meta.KeyUid)
	}
	return c.uid
}

func (c *contextImpl) AppId() string {
	if c.appId == "" {
		c.appId = c.metadata.GetString(meta.KeyAppId)
	}
	return c.appId
}

func (c *contextImpl) RoleIds() []int64 {
	if len(c.roleIds) == 0 {
		c.roleIds = c.metadata.GetInt64Slice(meta.KeyRoleIds)
	}
	return c.roleIds
}
