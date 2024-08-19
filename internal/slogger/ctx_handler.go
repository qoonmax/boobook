package slogger

import (
	"boobook/internal/http/contextkey"
	"context"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if rId := ctx.Value(contextkey.CtxKeyRequestID); rId != nil {
		r.AddAttrs(slog.Attr{
			Key:   string(contextkey.CtxKeyRequestID),
			Value: slog.StringValue(rId.(string)),
		})
	}

	return h.Handler.Handle(ctx, r)
}

func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{Handler: h.Handler.WithGroup(name)}
}
