// package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"url-shortener/internal/auth"
// )
// func JWTAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
// 		claims, err := auth.ParseJWT(tokenStr)
// 		if err != nil {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}
// 		ctx := r.Context()
// 		ctx = auth.WithUserID(ctx, claims.UserID)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

package middleware

import (
	"net/http"
	"url-shortener/internal/auth"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		claims, err := auth.ValidateJWT(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := auth.WithUserID(r.Context(), claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
