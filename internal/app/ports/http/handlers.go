package http

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samarec1812/segmentation-service/internal/app/entity/segment"
	"golang.org/x/exp/slog"
	"net/http"

	"github.com/go-chi/render"
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
			render.JSON(w, r, err)

			return
		}

		err := a.CreateSegment(reqBody.Slug)
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, SegmentErrorResponse(err))

			return
		}
		var sg segment.Segment
		render.JSON(w, r, SegmentSuccessResponse(sg))
	}
}
