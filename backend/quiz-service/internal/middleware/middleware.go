package middleware

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Пока просто пропускаем все запросы
		ctx := context.WithValue(r.Context(), "user_id", "11111111-1111-1111-1111-111111111111")
		ctx = context.WithValue(ctx, "role", "student")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
