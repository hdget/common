package biz

import (
	"strings"

	"github.com/hdget/utils/convert"
	"github.com/spf13/cast"
)

type MetaReader interface {
	GetAll() map[string]string
	GetString(key string) string
	GetInt64(key string) int64
	GetStringSlice(key string) []string
	GetInt64Slice(key string) []int64
	///* shortcut */
	//GetTid() int64
	//GetUid() int64
	//GetAppId() string
	//GetTsn() string
	//GetClient() string
}

type metaReaderImpl struct {
	metaMap map[string]string
}

//
//type serviceCtxReaderImpl struct {
//	*metaReaderImpl
//}
//
//func FromServiceContext(ctx context.Context) MetaReader {
//	return &serviceCtxReaderImpl{
//		metaReaderImpl: &metaReaderImpl{
//			metaMap: GetMetaFromContext(ctx),
//		},
//	}
//}
//
//func (r *metaReaderImpl) GetTid() int64 {
//	return r.GetInt64(KeyTid)
//}
//
//func (r *metaReaderImpl) GetUid() int64 {
//	return r.GetInt64(KeyUid)
//}
//
//func (r *metaReaderImpl) GetAppId() string {
//	return r.GetString(KeyAppId)
//}
//
//func (r *metaReaderImpl) GetTsn() string {
//	return r.GetString(KeyTsn)
//}
//
//func (r *metaReaderImpl) GetClient() string {
//	return r.GetString(KeyClient)
//}

func (r *metaReaderImpl) GetAll() map[string]string {
	return r.metaMap
}

func (r *metaReaderImpl) GetStringSlice(key string) []string {
	if v, exists := r.metaMap[key]; exists {
		return strings.Split(v, ",")
	}
	return nil
}

func (r *metaReaderImpl) GetInt64Slice(key string) []int64 {
	if v, exists := r.metaMap[key]; exists {
		return convert.CsvToInt64s(v)
	}
	return nil
}

func (r *metaReaderImpl) GetString(key string) string {
	if v, exists := r.metaMap[key]; exists {
		return v
	}
	return ""
}

func (r *metaReaderImpl) GetInt64(key string) int64 {
	if v, exists := r.metaMap[key]; exists {
		return cast.ToInt64(v)
	}
	return 0
}
