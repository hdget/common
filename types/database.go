package types

import (
	"github.com/hdget/common/protobuf"
)

// Database

type DbOperation[TKey AllowedKey, TBizObject any, TModelObject any, TCondition any] interface {
	DbCreate[TBizObject, TModelObject]
	DbRetrieve[TKey, TModelObject, TCondition]
	DbUpdate[TModelObject]
	DbDelete[TKey]
}

// DbCreate 创建操作:C
type DbCreate[BizObject any, ModelObject any] interface {
	Create(bizObj BizObject) (ModelObject, error) // 创建对象
}

// DbRetrieve 读取操作:R
type DbRetrieve[TKey AllowedKey, TModelObject any, TCondition any] interface {
	Get(key TKey) (TModelObject, error)                                                  // 获取对象
	Count(filters map[string]string) (int64, error)                                      // 统计对象
	List(filters map[string]string, list ...*protobuf.ListParam) ([]TModelObject, error) // 列出对象, list不传的时候获取所有对象
	GetQueryConditions(filters map[string]string) []TCondition                           // 获取查询条件
}

// DbUpdate 更新：U
type DbUpdate[TModelObject any] interface {
	Update(modelObj TModelObject) error // 更新某个字段
}

type DbEdit[TBizObject any] interface {
	Edit(bizObj TBizObject) error // 编辑对象
}

// DbDelete 删除
type DbDelete[TKey any] interface {
	Delete(key TKey) error // 删除对象
}

// DbBulkRetrieve 批量读取
type DbBulkRetrieve[TKey AllowedKey, ModelObject any] interface {
	BulkGet(keys []TKey) (map[TKey]ModelObject, error) // 批量获取对象
}

/* 关联数据表 */

type RefDbOperation[TKey AllowedKey, TRefBizObject any, TRefModelObject any, TCondition any] interface {
	RefDbCreate[TKey, TRefBizObject, TRefModelObject]
	RefDbRetrieve[TKey, TRefModelObject, TCondition]
	RefDbUpdate[TRefModelObject]
	RefDbDelete[TKey]
}

// RefDbCreate 创建关联对象操作:C
type RefDbCreate[TKey AllowedKey, RefBizObject any, RefModelObject any] interface {
	Create(key TKey, refBizObj RefBizObject) (RefModelObject, error) // 创建关联对象DAO
}

// RefDbRetrieve 读取关联对象操作:R
type RefDbRetrieve[TKey AllowedKey, RefModelObject any, Condition any] interface {
	Get(key, itemKey TKey) (RefModelObject, error)                                                         // 获取关联对象DAO
	Count(key TKey, refObjFilters map[string]string) (int64, error)                                        // 统计关联对象DAO
	List(key TKey, refObjFilters map[string]string, list ...*protobuf.ListParam) ([]RefModelObject, error) // 列出关联对象DAO
	GetQueryConditions(key TKey, refObjFilters map[string]string) []Condition                              // 获取关联对象DAO
}

// RefDbUpdate 更新关联对象：U
type RefDbUpdate[RefModelObject any] interface {
	Update(refModelObj RefModelObject) error // 更新数据库关联对象
}

type RefDbEdit[TKey AllowedKey, RefBizObject any] interface {
	Edit(key TKey, refBizObj RefBizObject) error // 编辑数据库关联对象DAO
}

// RefDbDelete 删除关联对象
type RefDbDelete[TKey AllowedKey] interface {
	Delete(key TKey, itemKey TKey) error // 删除关联对象DAO
}

// RefDbBulkRetrieve 批量读取关联对象
type RefDbBulkRetrieve[TKey AllowedKey, ModelObject any] interface {
	BulkGet(key TKey, itemKeys []TKey) (map[TKey]ModelObject, error) // 批量获取对象
}
