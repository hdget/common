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
