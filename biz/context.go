package biz

import (
	"context"
	"strconv"
	"sync"

	"google.golang.org/grpc/metadata"
)

type Context interface {
	WithValue(kvs ...any) Context // 添加键值
	Value(key any) (any, bool)    // 获取键值
	MD() map[string][]string      // 转换成GRPC metadata类型
	Tid() int64                   // 获取租户ID
	Uid() int64                   // 获取用户ID
	Usn() string                  // 获取用户SN
	AppId() string                // 获取应用ID
}

type contextImpl struct {
	kvStore sync.Map
}

const (
	ContextKeyAppId         = "hd-app-id"
	ContextKeyClient        = "hd-client"
	ContextKeyRelease       = "hd-release"
	ContextKeyTsn           = "hd-tsn"      // tenant sn
	ContextKeyTid           = "hd-tid"      // tenant id
	ContextKeyUid           = "hd-uid"      // user id
	ContextKeyUsn           = "hd-usn"      // user sn
	ContextKeyRoleIds       = "hd-role-ids" // role ids
	ContextKeyCaller        = "dapr-caller-app-id"
	ContextKeyDbTransaction = "hd-db-tx"
)

func NewContext(kvs ...any) Context {
	ctx := &contextImpl{
		kvStore: sync.Map{},
	}
	if len(kvs) > 0 && len(kvs)%2 == 0 {
		for i := 0; i < len(kvs); i += 2 {
			ctx.kvStore.Store(kvs[i], kvs[i+1])
		}
	}
	return ctx
}

// WithTidContext adds tid
func WithTidContext(parent Context, tid int64) Context {
	return parent.WithValue(ContextKeyTid, tid)
}

// WithTxContext adds db executor
func WithTxContext(parent Context, tx any) Context {
	return parent.WithValue(ContextKeyDbTransaction, tx)
}

func NewFromIncomingGrpcContext(ctx context.Context) Context {
	c := &contextImpl{
		kvStore: sync.Map{},
	}

	// 尝试从grpc context中获取meta信息
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return c
	}

	for key, value := range md {
		switch key {
		case ContextKeyTid, ContextKeyUid: // int64
			v, _ := strconv.ParseInt(value[0], 10, 64)
			c.kvStore.Store(key, v)
		default:
			c.kvStore.Store(key, value[0])
		}
	}
	return c
}

func NewOutgoingGrpcContext(ctx Context) context.Context {
	return metadata.NewOutgoingContext(context.Background(), ctx.MD())
}

func (c *contextImpl) WithValue(kvs ...any) Context {
	if len(kvs) > 0 && len(kvs)%2 == 0 {
		for i := 0; i < len(kvs); i += 2 {
			c.kvStore.Store(kvs[i], kvs[i+1])
		}
	}
	return c
}

func (c *contextImpl) Value(key any) (any, bool) {
	return c.kvStore.Load(key)
}

func (c *contextImpl) MD() map[string][]string {
	md := make(map[string][]string)
	c.kvStore.Range(func(key, value any) bool {
		if k, ok := key.(string); ok {
			switch v := value.(type) {
			case string:
				md[k] = []string{v}
			case int64:
				md[k] = []string{strconv.FormatInt(v, 10)}
			default:
				// 对于未处理类型，使用fallback方案
				md[k] = []string{""}
			}
		}
		return true
	})
	return md
}

func (c *contextImpl) Tid() int64 {
	if v, exists := c.kvStore.Load(ContextKeyTid); exists {
		if tid, ok := v.(int64); ok {
			return tid
		}
	}
	return 0
}

func (c *contextImpl) Uid() int64 {
	if v, exists := c.kvStore.Load(ContextKeyUid); exists {
		if uid, ok := v.(int64); ok {
			return uid
		}
	}
	return 0
}

func (c *contextImpl) AppId() string {
	if v, exists := c.kvStore.Load(ContextKeyAppId); exists {
		if appId, ok := v.(string); ok {
			return appId
		}
	}
	return ""
}

func (c *contextImpl) Usn() string {
	if v, exists := c.kvStore.Load(ContextKeyUsn); exists {
		if appId, ok := v.(string); ok {
			return appId
		}
	}
	return ""
}
