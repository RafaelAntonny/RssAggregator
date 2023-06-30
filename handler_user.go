package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/RafaelAntonny/RSS_AGGREGATOR/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))

}

func (apiCfg *apiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to get posts: %v", err))
		return
	}

	respondWithJson(w, 200, databasePostToPosts(posts))
}
