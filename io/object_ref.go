package io

import "github.com/hdget/common/protobuf"

/* request */

type OperateRefObjectRequest[BizObject any] struct {
	Id   int64     `json:"id"`
	Item BizObject `json:"item"`
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
