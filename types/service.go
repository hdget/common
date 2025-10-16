package types

import (
	"github.com/hdget/common/biz"
	"github.com/hdget/common/protobuf"
)

type ServiceQueryResponse[BizObject any] struct {
	Total int64       `json:"total"`
	Items []BizObject `json:"items"`
}

type ServiceOperation[BizObject any] interface {
	ServiceCreate[BizObject]
	ServiceRetrieve[BizObject]
	ServiceUpdate[BizObject]
	ServiceDelete
}

type ServiceCreate[BizObject any] interface {
	Create(ctx biz.Context, object BizObject) (int64, error) // 创建业务对象
}

type ServiceRetrieve[BizObject any] interface {
	Get(ctx biz.Context, id int64) (BizObject, error)                                                       // 获取业务对象
	Query(ctx biz.Context, filters map[string]string, list *protobuf.ListParam) (int64, []BizObject, error) // 查询业务对象
}

type ServiceUpdate[BizObject any] interface {
	Edit(ctx biz.Context, bizObject BizObject) error // 编辑业务对象
}

type ServiceDelete interface {
	Delete(ctx biz.Context, id int64) error // 删除业务对象
}
