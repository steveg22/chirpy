package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/steveg22/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUsersUpgrade(w http.ResponseWriter, r *http.Request) {

	type requestBody struct {
		Event string `json:"event"`
		Data  struct {
			UserId uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find api key", err)
		return
	}
	if apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "API key is invalid", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	body := requestBody{}
	err = decoder.Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Body", err)
		return
	}

	if body.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, struct{}{})
		return
	}

	err = cfg.db.UpgradeMembership(r.Context(), body.Data.UserId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, struct{}{})

}
