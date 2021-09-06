package controllers

import (
	"avtoru/models"
	u "avtoru/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Unable to decode request body", http.StatusInternalServerError)
		return
	}

	status, err := account.Create()
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	} //Create account
	u.Respond(w, "")
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		http.Error(w, "Unable to decode request body", http.StatusInternalServerError)
		return
	}

	token, err, status := models.Login(account.Email, account.Password)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	u.Respond(w, token)
}

var GetAccount = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	accountModel := models.GetAccountById(accountId)
	resp := map[string]interface{}{
		"id":         accountModel.ID,
		"created_at": accountModel.CreatedAt,
		"updated_at": accountModel.UpdatedAt,
	}
	u.Respond(w, resp)
}
