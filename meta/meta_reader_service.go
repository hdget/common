package meta

import (
	"context"
)

type serviceCtxReaderImpl struct {
	*metaCtxReaderImpl
}

func FromServiceContext(ctx context.Context) MetaReader {
	return &serviceCtxReaderImpl{
		metaCtxReaderImpl: &metaCtxReaderImpl{
			metas: GetMetaFromContext(ctx),
		},
	}
}
