package io

import "github.com/hdget/common/protobuf"

/* request */

type OperateObjectRequest struct {
	Id   int64  `json:"id,omitempty"`
	Sn   string `json:"sns,omitempty"`
	Name string `json:"name,omitempty"`
}

type BulkOperateObjectRequest struct {
	Ids   []int64  `json:"ids,omitempty"`
	Sns   []string `json:"sns,omitempty"`
	Names []string `json:"names,omitempty"`
}

type QueryObjectRequest struct {
	Filters map[string]string   `json:"filters,omitempty"`
	List    *protobuf.ListParam `json:"list,omitempty"`
}

/* response */

type CreateObjectResponse struct {
	Id int64 `json:"id"`
}

type QueryObjectResponse[BizObject any] struct {
	Total int64       `json:"total"`
	Items []BizObject `json:"items"`
}
