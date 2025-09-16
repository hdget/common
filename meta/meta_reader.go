package meta

import (
	"github.com/spf13/cast"
)

type MetaReader interface {
	GetTid() int64
	GetUid() int64
	GetAppId() string
	GetTsn() string
}

type metaCtxReaderImpl struct {
	metas map[string]string
}

func (r *metaCtxReaderImpl) GetTid() int64 {
	return r.getInt64(KeyTid)
}

func (r *metaCtxReaderImpl) GetUid() int64 {
	return r.getInt64(KeyUid)
}

func (r *metaCtxReaderImpl) GetAppId() string {
	return r.getString(KeyAppId)
}

func (r *metaCtxReaderImpl) GetTsn() string {
	return r.getString(KeyTsn)
}

//
//func (r *metaCtxReaderImpl) getStringSlice(key string) []string {
//	v, exists := r.metas[key]
//	if !exists {
//		return nil
//	}
//	return strings.Split(v, ",")
//}
//
//func (r *metaCtxReaderImpl) getInt64Slice(key string) []int64 {
//	v, exists := r.metas[key]
//	if !exists {
//		return nil
//	}
//	return convert.CsvToInt64s(v)
//}

func (r *metaCtxReaderImpl) getString(key string) string {
	v, exists := r.metas[key]
	if !exists {
		return ""
	}
	return v
}

func (r *metaCtxReaderImpl) getInt64(key string) int64 {
	v, exists := r.metas[key]
	if !exists {
		return 0
	}
	return cast.ToInt64(v)
}
