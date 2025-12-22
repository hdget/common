package io

import "github.com/hdget/common/protobuf"

/* request */

type CreateRefObjectRequest[BizObject any] struct {
	Id   int64     `json:"id"`
	Item BizObject `json:"item"`
}

type EditRefObjectRequest[BizObject any] struct {
	Id   int64     `json:"id"`
	Item BizObject `json:"item"`
}

type DeleteRefObjectRequest struct {
	Id     int64 `json:"id"`
	ItemId int64 `json:"itemId"`
}

type GetRefObjectRequest struct {
	Id     int64 `json:"id"`
	ItemId int64 `json:"itemId"`
}

type QueryRefObjectRequest struct {
	Id      int64               `json:"id"`
	Filters map[string]any      `json:"filters,omitempty"`
	List    *protobuf.ListParam `json:"list,omitempty"`
}

/* response */

type QueryRefObjectResponse[BizObject any] struct {
	Total int64       `json:"total"`
	Items []BizObject `json:"items"`
}
