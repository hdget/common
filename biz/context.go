package biz

import (
	"context"

	"github.com/hdget/common/constant"
	"github.com/spf13/cast"
	"google.golang.org/grpc/metadata"
)

type Context interface {
	Meta() MetaReader
	SetTx(tx any)
	GetTx() any
	Tid() int64
	Uid() int64
	AppId() string
}

type contextImpl struct {
	metadata map[string]string
	tx       any
	appId    string
	tid      int64 // 租户ID, 高频对象缓存
	uid      int64 // 当前用户ID, 高频对象缓存
}

func NewContext(metaKvs ...string) Context {
	return &contextImpl{
		metadata: toMap(metaKvs...),
	}
}

func NewFromIncomingGrpcContext(ctx context.Context) Context {
	c := &contextImpl{}

	// 尝试从grpc context中获取meta信息
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return c
	}

	if len(md) > 0 {
		c.metadata = make(map[string]string)
		for key, value := range md {
			c.metadata[key] = value[0]
		}
	}

	return c
}

func (c *contextImpl) Meta() MetaReader {
	return &metaReaderImpl{
		metaMap: c.metadata,
	}
}

func (c *contextImpl) SetTx(tx any) {
	c.tx = tx
}

func (c *contextImpl) GetTx() any {
	return c.tx
}

func (c *contextImpl) Tid() int64 {
	if c.tid > 0 {
		return c.tid
	}
	return c.Meta().GetInt64(constant.MetaKeyTid)
}

func (c *contextImpl) Uid() int64 {
	if c.uid > 0 {
		return c.uid
	}
	return c.Meta().GetInt64(constant.MetaKeyUid)
}

func (c *contextImpl) AppId() string {
	if c.appId != "" {
		return c.appId
	}
	return c.Meta().GetString(constant.MetaKeyAppId)
}

type debugContextImpl struct {
	*contextImpl
}

func NewDebugContext(tid int64, kvs ...string) Context {
	c := &debugContextImpl{
		contextImpl: &contextImpl{
			metadata: map[string]string{
				constant.MetaKeyTid: cast.ToString(tid),
			},
		},
	}

	for key, value := range toMap(kvs...) {
		c.metadata[key] = value
	}

	return c
}

func NewOutgoingGrpcContext(ctx Context) context.Context {
	bizMetaKvs := ctx.Meta().GetAll()

	md := make(map[string][]string, len(bizMetaKvs))
	for k, v := range bizMetaKvs {
		md[k] = []string{v}
	}
	return metadata.NewOutgoingContext(context.Background(), md)
}

func toMap(kvs ...string) map[string]string {
	if len(kvs) == 0 || len(kvs)%2 == 1 {
		return nil
	}

	result := make(map[string]string, len(kvs)/2)
	for i := 0; i < len(kvs); i += 2 {
		result[kvs[i]] = kvs[i+1]
	}
	return result
}
