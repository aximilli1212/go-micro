package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.writeJSON(w, http.StatusBadRequest, jsonResponse{
			Error:   true,
			Message: "Invalid request payload",
		})
		return
	}

	//validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid email or password"), http.StatusUnauthorized)
		return
	}

	//If we have the user in the DB , check the password
	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid email or password"), http.StatusUnauthorized)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Welcome %s", user.FirstName),
		Data:    user,
	}

	app.writeJSON(w, http.StatusOK, payload)
}
