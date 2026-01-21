package object

import "github.com/hdget/common/protobuf"

/* request */

type Identifier interface {
	int64 | int32 | int | string
}

type OperateObjectRequest[TObjectId Identifier] struct {
	Id TObjectId `json:"id"`
}

type BulkOperateObjectRequest[TObjectId Identifier] struct {
	Ids []TObjectId `json:"ids"`
}

type QueryObjectRequest struct {
	Filters map[string]string   `json:"filters,omitempty"`
	List    *protobuf.ListParam `json:"list,omitempty"`
}

/* response */

type CreateObjectResponse[TObjectId Identifier] struct {
	Id TObjectId `json:"id"`
}

type QueryObjectResponse[TBizObject any] struct {
	Total int64        `json:"total"`
	Items []TBizObject `json:"items"`
}
