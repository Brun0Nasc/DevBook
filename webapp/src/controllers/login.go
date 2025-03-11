package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// DoLogin
func DoLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
	}

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.APIError{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.APIError{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	var authData models.AuthData
	if err = json.NewDecoder(response.Body).Decode(&authData); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	if err = cookies.Save(w, authData.ID, authData.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.APIError{Error: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
