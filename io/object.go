package io

import "github.com/hdget/common/protobuf"

/* request */

type GetObjectRequest struct {
	Id int64 `json:"id"`
}

type BulkGetObjectRequest struct {
	Ids []int64 `json:"ids"`
}

type ApproveObjectRequest struct {
	Id int64 `json:"id"`
}

type DeleteObjectRequest struct {
	Id int64 `json:"id"`
}

type QueryObjectRequest struct {
	Filters map[string]any      `json:"filters"`
	List    *protobuf.ListParam `json:"list"`
}

/* response */

type CreateObjectResponse struct {
	Id int64 `json:"id"`
}

type QueryObjectResponse[BizObject any] struct {
	Total int64       `json:"total"`
	Items []BizObject `json:"items"`
}
