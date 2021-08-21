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
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

var GetAccount = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	accountModel := models.GetAccountById(accountId)
	resp := u.Message(true, "Success")
	resp["data"] = map[string]interface{}{
		"id": accountModel.ID,
		"created_at": accountModel.CreatedAt,
		"updated_at": accountModel.UpdatedAt,
	}
	u.Respond(w, resp)
}
