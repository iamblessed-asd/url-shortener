package auth

import (
	"context"
)

type contextKey string

const userIDKey = contextKey("userID")

// Сохраняет userID в context
func WithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// Извлекает userID из context
func GetUserID(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(userIDKey).(int)
	return id, ok
}
