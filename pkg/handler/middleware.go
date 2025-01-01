package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			WriteError(w, 401, fmt.Errorf("empty auth header"))
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			WriteError(w, 401, fmt.Errorf("invalid auth header"))
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			WriteError(w, 401, fmt.Errorf("invalid auth header"))
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserId(w http.ResponseWriter, r *http.Request) (int, error) {
	id := r.Context().Value(userCtx)
	if id == nil {
		WriteError(w, 500, fmt.Errorf("user id not found"))
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		WriteError(w, 500, fmt.Errorf("user id not found"))
		return 0, errors.New("invalid user id")
	}

	return idInt, nil
}
