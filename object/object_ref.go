package object

import (
	"github.com/hdget/common/protobuf"
)

/* request */

type CreateRefObjectRequest[TObjectId Identifier, TBizObject any] struct {
	Id   TObjectId  `json:"id"`
	Item TBizObject `json:"item"`
}

type EditRefObjectRequest[TObjectId Identifier, TBizObject any] struct {
	Id   TObjectId  `json:"id"`
	Item TBizObject `json:"item"`
}

type DeleteRefObjectRequest[TObjectId Identifier] struct {
	Id     TObjectId `json:"id"`
	ItemId TObjectId `json:"itemId"`
}

type GetRefObjectRequest[TObjectId Identifier] struct {
	Id     TObjectId `json:"id"`
	ItemId TObjectId `json:"itemId"`
}

type QueryRefObjectRequest[TObjectId Identifier] struct {
	Id      TObjectId           `json:"id"`
	Filters map[string]string   `json:"filters,omitempty"`
	List    *protobuf.ListParam `json:"list,omitempty"`
}

/* response */

type QueryRefObjectResponse[TBizObject any] struct {
	Total int64        `json:"total"`
	Items []TBizObject `json:"items"`
}
