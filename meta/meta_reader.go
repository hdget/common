package meta

import (
	"github.com/hdget/utils/convert"
	"github.com/spf13/cast"
	"strings"
)

type MetaReader interface {
	GetStringSlice(key string) []string
	GetInt64Slice(key string) []int64
	GetString(key string) string
	GetInt64(key string) int64
}

type metaReaderImpl struct {
	metas map[string]string
}

func (c *metaReaderImpl) GetStringSlice(key string) []string {
	v, exists := c.metas[key]
	if !exists {
		return nil
	}
	return strings.Split(v, ",")
}

func (c *metaReaderImpl) GetInt64Slice(key string) []int64 {
	v, exists := c.metas[key]
	if !exists {
		return nil
	}
	return convert.CsvToInt64s(v)
}

func (c *metaReaderImpl) GetString(key string) string {
	v, exists := c.metas[key]
	if !exists {
		return ""
	}
	return v
}

func (c *metaReaderImpl) GetInt64(key string) int64 {
	v, exists := c.metas[key]
	if !exists {
		return 0
	}
	return cast.ToInt64(v)
}
