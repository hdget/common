package types

import (
	"github.com/hdget/common/biz"
	"github.com/hdget/common/protobuf"
)

type AllowedKey interface {
	int64 | int32 | int | string
}

type ServiceOperation[TKey AllowedKey, TBizObject any] interface {
	ServiceCreate[TBizObject]
	ServiceRetrieve[TKey, TBizObject]
	ServiceUpdate[TBizObject]
	ServiceDelete[TKey]
}

type ServiceCreate[TBizObject any] interface {
	Create(ctx biz.Context, object TBizObject) (int64, error) // 创建业务对象
}

type ServiceRetrieve[TKey AllowedKey, TBizObject any] interface {
	Get(ctx biz.Context, key TKey) (TBizObject, error)                                                          // 获取业务对象
	Query(ctx biz.Context, filters map[string]string, list ...*protobuf.ListParam) (int64, []TBizObject, error) // 查询业务对象
}

type ServiceUpdate[TBizObject any] interface {
	Edit(ctx biz.Context, bizObject TBizObject) error // 编辑业务对象
}

type ServiceDelete[TKey AllowedKey] interface {
	Delete(ctx biz.Context, key TKey) error // 删除业务对象
}
