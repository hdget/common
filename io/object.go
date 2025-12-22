package io

import "github.com/hdget/common/protobuf"

/* request */

type GetObjectRequest struct {
	Id   int64  `json:"id,omitempty"`
	Sn   string `json:"sns,omitempty"`
	Name string `json:"name,omitempty"`
}

type BulkGetObjectRequest struct {
	Ids   []int64  `json:"ids,omitempty"`
	Sns   []string `json:"sns,omitempty"`
	Names []string `json:"names,omitempty"`
}

type ApproveObjectRequest struct {
	Id   int64  `json:"id,omitempty"`
	Sn   string `json:"sns,omitempty"`
	Name string `json:"name,omitempty"`
}

type DeleteObjectRequest struct {
	Id   int64  `json:"id,omitempty"`
	Sn   string `json:"sns,omitempty"`
	Name string `json:"name,omitempty"`
}

type QueryObjectRequest struct {
	Filters map[string]any      `json:"filters,omitempty"`
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
