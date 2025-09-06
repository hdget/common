package types

import (
	"github.com/hdget/common/protobuf"
)

type BizSvcQueryResponse[BizObject any] struct {
	Total int64       `json:"total"`
	Items []BizObject `json:"items"`
}

type BizSvcOperation[BizObject any] interface {
	BizSvcCreate[BizObject]
	BizSvcRetrieve[BizObject]
	BizSvcUpdate[BizObject]
	BizSvcDelete
}

type BizSvcCreate[BizObject any] interface {
	Create(object BizObject) (int64, error) // 创建业务对象
}

type BizSvcRetrieve[BizObject any] interface {
	Get(id int64) (BizObject, error)                                                                    // 获取业务对象
	Query(filters map[string]string, list *protobuf.ListParam) (*BizSvcQueryResponse[BizObject], error) // 查询业务对象
}

type BizSvcUpdate[BizObject any] interface {
	Edit(BizObject) error // 编辑业务对象
}

type BizSvcDelete interface {
	Delete(id int64) error // 删除业务对象
}
