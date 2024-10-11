package middleware

import (
	"fmt"
	"log/slog"

	"github.com/ogen-go/ogen/middleware"
)

func Logging() middleware.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		slog.DebugContext(
			req.Context, fmt.Sprintf("[http] <-- %s", req.OperationName),
			slog.String("path", req.Raw.URL.String()),
			slog.Any("body", req.Body),
		)

		resp, err := next(req)
		if err != nil {
			slog.DebugContext(
				req.Context, fmt.Sprintf("FAIL %s", req.OperationName),
				slog.String("message", err.Error()),
			)
		} else {
			slog.DebugContext(
				req.Context, fmt.Sprintf("[http] --> %s", req.OperationName),
				slog.Any("body", resp.Type),
			)
		}

		return resp, err
	}
}
