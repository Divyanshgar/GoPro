package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Divyanshgar/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Error parsing JSON")
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Url:       params.URL,
	})
	if err != nil {
		respondWithError(w, 400, "Error creating feed")
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiconfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not fetch feeds %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))
}
