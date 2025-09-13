package meta

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type grpcMetaReaderImpl struct {
	*metaReaderImpl
}

func FromGrpcContext(ctx context.Context) MetaReader {
	// 尝试从grpc context中获取meta信息
	if md, exists := metadata.FromIncomingContext(ctx); exists {
		metas := make(map[string]string)
		for key, values := range md {
			if len(values) > 0 {
				metas[key] = values[0]
			}
		}
		return &grpcMetaReaderImpl{
			metaReaderImpl: &metaReaderImpl{
				metas: metas,
			},
		}
	}

	return &grpcMetaReaderImpl{
		metaReaderImpl: &metaReaderImpl{
			metas: make(map[string]string),
		},
	}
}
