package http

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"

	"github.com/samarec1812/segmentation-service/internal/app/entity/segment"
	"github.com/samarec1812/segmentation-service/internal/app/service"
)

func createSegment(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.segment.create"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody createSegmentRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}

		err := a.CreateSegment(r.Context(), reqBody.Slug)
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}
		var sg segment.Segment
		render.JSON(w, r, SegmentSuccessResponse(sg))
	}
}

func deleteSegment(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.segment.delete"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody deleteSegmentRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}

		err := a.RemoveSegment(r.Context(), reqBody.Slug)
		if err != nil {
			log.Error("error with delete", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}

		render.JSON(w, r, SegmentDeleteResponse())
	}
}

func addUser(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user-segment.addUser"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody addUserToSegmentRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}

		err := a.AddUserToSegment(r.Context(), reqBody.UserID, reqBody.AddSlugs, reqBody.RemoveSlugs)
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}
		var sg segment.Segment
		render.JSON(w, r, SegmentSuccessResponse(sg))
	}
}

func getUserSegments(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.user-segment.getUserSegments"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody userSegmentListRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}

		segments, err := a.GetSegments(r.Context(), reqBody.UserID)
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}

		render.JSON(w, r, UserSegmentListResponse(userSegmentListResponse{
			reqBody.UserID,
			segments,
		}))
	}
}
