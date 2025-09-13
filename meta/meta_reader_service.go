package meta

import (
	"context"
	"github.com/hdget/common/servicectx"
)

type serviceMetaReaderImpl struct {
	*metaReaderImpl
}

func FromServiceContext(ctx context.Context) MetaReader {
	return &serviceMetaReaderImpl{
		metaReaderImpl: &metaReaderImpl{
			metas: servicectx.GetMeta(ctx),
		},
	}
}
