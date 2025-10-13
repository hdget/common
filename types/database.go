package types

import (
	"github.com/hdget/common/protobuf"
)

// Database

type DbOperation[BizObject any, ModelObject any, Condition any] interface {
	DbCreate[BizObject, ModelObject]
	DbRetrieve[ModelObject, Condition]
	DbUpdate[BizObject, ModelObject]
	DbDelete
}

// DbCreate 创建操作:C
type DbCreate[BizObject any, ModelObject any] interface {
	Create(bizObj BizObject) (ModelObject, error) // 创建对象
}

// DbRetrieve 读取操作:R
type DbRetrieve[ModelObject any, Condition any] interface {
	Get(id int64) (ModelObject, error)                                                  // 获取对象
	Find(filters map[string]string) (ModelObject, error)                                // 查找某个对象
	Count(filters map[string]string) (int64, error)                                     // 统计对象
	List(filters map[string]string, list ...*protobuf.ListParam) ([]ModelObject, error) // 列出对象, list不传的时候获取所有对象
	GetQueryConditions(filters map[string]string) []Condition                           // 获取查询条件
}

// DbUpdate 更新：U
type DbUpdate[BizObject any, ModelObject any] interface {
	Edit(bizObj BizObject) error       // 编辑对象
	Update(modelObj ModelObject) error // 更新某个字段
}

// DbDelete 删除
type DbDelete interface {
	Delete(id int64) error // 删除对象
}

// DbBulkRetrieve 批量读取
type DbBulkRetrieve[ModelObject any] interface {
	BulkGet(ids []int64) (map[int64]ModelObject, error) // 批量获取对象
}

/* 关联数据表 */

type RefDbOperation[RefBizObject any, RefModelObject any, Condition any] interface {
	RefDbCreate[RefBizObject, RefModelObject]
	RefDbRetrieve[RefModelObject, Condition]
	RefDbUpdate[RefBizObject, RefModelObject]
	RefDbDelete
}

// RefDbCreate 创建关联对象操作:C
type RefDbCreate[RefBizObject any, RefModelObject any] interface {
	Create(id int64, refBizObj RefBizObject) (RefModelObject, error) // 创建关联对象DAO
}

// RefDbRetrieve 读取关联对象操作:R
type RefDbRetrieve[RefModelObject any, Condition any] interface {
	Get(id, refId int64) (RefModelObject, error)                                                           // 获取关联对象DAO
	Find(id int64, refObjFilters map[string]string) (RefModelObject, error)                                // 查找某个对象
	Count(id int64, refObjFilters map[string]string) (int64, error)                                        // 统计关联对象DAO
	List(id int64, refObjFilters map[string]string, list ...*protobuf.ListParam) ([]RefModelObject, error) // 列出关联对象DAO
	GetQueryConditions(id int64, refObjFilters map[string]string) []Condition                              // 获取关联对象DAO
}

// RefDbUpdate 更新关联对象：U
type RefDbUpdate[RefBizObject any, RefModelObject any] interface {
	Edit(id int64, refBizObj RefBizObject) error // 编辑数据库关联对象DAO
	Update(refModelObj RefModelObject) error     // 更新数据库关联对象
}

// RefDbDelete 删除关联对象
type RefDbDelete interface {
	Delete(id int64, refId int64) error // 删除关联对象DAO
}

// RefDbBulkRetrieve 批量读取关联对象
type RefDbBulkRetrieve[ModelObject any] interface {
	BulkGet(id int64, refIds []int64) (map[int64]ModelObject, error) // 批量获取对象
}
