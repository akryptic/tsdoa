package utils

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EmitEmptyEvent(ctx context.Context, key string) {
	runtime.EventsEmit(ctx, key, nil)
}

func EmitEvent[T any](ctx context.Context, key string, data T) {
	runtime.EventsEmit(ctx, key, data)
}
