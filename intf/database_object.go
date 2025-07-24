package intf

import "github.com/hdget/common/protobuf"

type DatabaseObject[Object any, Condition any] interface {
	Create(props map[string]any) (int64, error)                                 // 创建对象
	Delete(id int64) error                                                      // 删除对象
	Edit(props map[string]any) error                                            // 编辑对象
	Get(id int64) (Object, error)                                               // 获取对象
	Count(filters map[string]string) (int64, error)                             // 统计对象
	List(filters map[string]string, list *protobuf.ListParam) ([]Object, error) // 列出对象
	GetQueryConditions(filters map[string]string) []Condition                   // 获取查询条件
}
