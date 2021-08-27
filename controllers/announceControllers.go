package controllers

import (
	"avtoru/models"
	u "avtoru/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var GetAnnById = func(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok || len(id[0]) < 1 {
		u.Respond(w, u.Message(false, "ID is missing in URL string"))
	}
	id_conv, _ := strconv.ParseUint(id[0], 10, 32)
	announcementModel := models.GetAnnById(uint(id_conv))
	resp := u.Message(true, "Success")
	resp["data"] = announcementModel
	u.Respond(w, resp)
}

var GetAnns = func(w http.ResponseWriter, r *http.Request) {
	announcements := models.GetAnns()
	resp := map[string]interface{}{"data": announcements}
	//for index, announce := range announcements {
	//	resp[string(index)] = map[string]interface{}{"id": announce.ID}
	//}

	u.Respond(w, resp)
}

var AddAn = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	announcement := &models.Announce{}
	announcement.AccountID = accountId
	err := json.NewDecoder(r.Body).Decode(announcement) //decode the request body into struct and failed if any error occur
	if err != nil {
		fmt.Println(err.Error())
		u.Respond(w, u.Message(false, "Required field filled incorrectly"))
		return
	}

	announcement = announcement.AddAn() //(r.Context().Value("AccountId"))

	resp := map[string]interface{}{"Message": "The announcement was published successfully"}

	u.Respond(w, resp)
}

var DelAn = func(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok || len(id[0]) < 1 {
		u.Respond(w, u.Message(false, "ID is missing in URL string"))
	}
	announcement := &models.Announce{}
	id_conv, _ := strconv.ParseUint(id[0], 10, 32)
	announcementModel := models.GetAnnById(uint(id_conv))
	compare := announcementModel.AccountID
	accountId := r.Context().Value("user").(uint) //
	if accountId != compare {
		u.Respond(w, u.Message(false, "Access denied"))
		return
	}
	//id, ok := r.URL.Query()["id"]
	//if !ok || len(id[0]) < 1 {
	//	u.Respond(w, u.Message(false, "ID is missing in URL string"))
	//}
	//id_conv, _ := strconv.ParseUint(id[0], 10, 64)
	result := announcement.DelAn(announcement.ID)

	resp := map[string]interface{}{"Message": "The announcement was deleted", "result": result}
	u.Respond(w, resp)
}
