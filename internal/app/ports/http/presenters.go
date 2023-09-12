package http

import "github.com/samarec1812/segmentation-service/internal/app/entity/segment"

type createSegmentRequest struct {
	Slug string `json:"slug"`
}

type deleteSegmentRequest struct {
	Slug string `json:"slug"`
}

type addUserToSegmentRequest struct {
	ID          int64    `json:"id"`
	AddSlugs    []string `json:"add_slugs"`
	RemoveSlugs []string `json:"remove_slugs"`
}

type userSegmentListRequest struct {
	ID int64 `json:"id"`
}

type userSegmentListResponse struct {
	ID       int64    `json:"id"`
	Segments []string `json:"segments"`
}

func SegmentErrorResponse(err error) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = nil
	resp["error"] = err

	return resp
}

func UserSegmentListResponse(res userSegmentListResponse) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = res
	resp["error"] = nil

	return resp
}

func SegmentDeleteResponse() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = nil
	resp["error"] = nil

	return resp
}

func SegmentSuccessResponse(sg segment.Segment) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["data"] = sg
	resp["error"] = nil

	return resp
}
