package intf

import "github.com/hdget/common/protobuf"

type DataAccessObject[businessObject any, modelObject any, Condition any] interface {
	Create(businessObject) (int64, error)                                            // 创建对象
	Edit(businessObject) error                                                       // 编辑对象
	Delete(id int64) error                                                           // 删除对象
	Get(id int64) (modelObject, error)                                               // 获取对象
	BulkGet(ids []int64) (map[int64]modelObject, error)                              // 批量获取对象
	Count(filters map[string]string) (int64, error)                                  // 统计对象
	List(filters map[string]string, list *protobuf.ListParam) ([]modelObject, error) // 列出对象
	GetQueryConditions(filters map[string]string) []Condition                        // 获取查询条件
}
