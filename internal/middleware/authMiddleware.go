package middleware

import (
	"cloud-disk/utils"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("登录时效，或者未登录"))
			return
		}
		token, err := utils.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserId", string(rune(token.Id)))
		r.Header.Set("UserIdentity", token.Identity)
		r.Header.Set("UserName", token.Name)
		r.Header.Set("UserName", token.Name)
		next(w, r)
	}
}
