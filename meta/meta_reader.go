package meta

import (
	"strings"

	"github.com/hdget/utils/convert"
	"github.com/spf13/cast"
)

type MetaReader interface {
	GetString(key string) string
	GetInt64(key string) int64
	GetStringSlice(key string) []string
	GetInt64Slice(key string) []int64
	/* shortcut */
	GetTid() int64
	GetUid() int64
	GetAppId() string
	GetTsn() string
	GetClient() string
}

type metaCtxReaderImpl struct {
	metaMap map[string]string
}

func (r *metaCtxReaderImpl) GetTid() int64 {
	return r.GetInt64(KeyTid)
}

func (r *metaCtxReaderImpl) GetUid() int64 {
	return r.GetInt64(KeyUid)
}

func (r *metaCtxReaderImpl) GetAppId() string {
	return r.GetString(KeyAppId)
}

func (r *metaCtxReaderImpl) GetTsn() string {
	return r.GetString(KeyTsn)
}

func (r *metaCtxReaderImpl) GetClient() string {
	return r.GetString(KeyClient)
}

func (r *metaCtxReaderImpl) GetStringSlice(key string) []string {
	if v, exists := r.metaMap[key]; exists {
		return strings.Split(v, ",")
	}
	return nil
}

func (r *metaCtxReaderImpl) GetInt64Slice(key string) []int64 {
	if v, exists := r.metaMap[key]; exists {
		return convert.CsvToInt64s(v)
	}
	return nil
}

func (r *metaCtxReaderImpl) GetString(key string) string {
	if v, exists := r.metaMap[key]; exists {
		return v
	}
	return ""
}

func (r *metaCtxReaderImpl) GetInt64(key string) int64 {
	if v, exists := r.metaMap[key]; exists {
		return cast.ToInt64(v)
	}
	return 0
}
