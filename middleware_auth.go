package main

import (
	"fmt"
	"net/http"

	"github.com/Divyanshgar/rssagg/internal/auth"
	"github.com/Divyanshgar/rssagg/internal/database"
)

type authedhandler func(http.ResponseWriter, *http.Request, database.User)

func(apicfg *apiconfig) middlewareAuth(handler authedhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}	
	user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apikey)
	if err != nil { 
		respondWithError(w, 402, fmt.Sprintf("Could not find user with the given API key %v", err))
		return
	}
	handler(w, r, user)
	}
}