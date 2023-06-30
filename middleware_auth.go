package main

import (
	"fmt"
	"net/http"

	"github.com/RafaelAntonny/RSS_AGGREGATOR/internal/auth"
	"github.com/RafaelAntonny/RSS_AGGREGATOR/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user %v", err))
			return
		}

		handler(w, r, user)
	}
}
