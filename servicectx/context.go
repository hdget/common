package servicectx

import (
	"context"

	"github.com/hdget/common/constant"
	"google.golang.org/grpc/metadata"
)

var (
	awareMetaKeys = []string{
		constant.MetaKeyTid,
		constant.MetaKeyUid,
		constant.MetaKeyAppId,
		constant.MetaKeyTsn,
	}
)

func New(ctx context.Context) context.Context {
	// 尝试从grpc context中获取meta信息
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return context.Background()
	}

	metas := make(map[string]string)
	for _, key := range awareMetaKeys {
		if values := md.Get(key); len(values) > 0 {
			metas[key] = values[0]
		}
	}
	return addMeta(context.Background(), metas)
}

func NewOutgoingGrpcContext(ctx context.Context) context.Context {
	metas := GetMeta(ctx)
	if len(metas) == 0 {
		return context.Background()
	}
	
	md := make(map[string][]string, len(metas))
	for k, v := range metas {
		md[k] = []string{v}
	}
	return metadata.NewOutgoingContext(ctx, md)
}
