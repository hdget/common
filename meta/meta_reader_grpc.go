package meta

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type grpcCtxReaderImpl struct {
	*metaCtxReaderImpl
}

func FromGrpcContext(ctx context.Context) MetaReader {
	// 尝试从grpc context中获取meta信息
	var metaMap map[string]string
	if md, exists := metadata.FromIncomingContext(ctx); exists {
		metaMap = make(map[string]string)
		for key, values := range md {
			if len(values) > 0 {
				metaMap[key] = values[0]
			}
		}
	}

	return &grpcCtxReaderImpl{
		metaCtxReaderImpl: &metaCtxReaderImpl{
			metaMap: metaMap,
		},
	}
}
