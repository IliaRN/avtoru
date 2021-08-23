package controllers

import (
	"avtoru/models"
	u "avtoru/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var GetAnn = func(w http.ResponseWriter, r *http.Request) {
	announcements := models.GetAnn()
	resp := map[string]interface{}{"data": announcements}
	//for index, announce := range announcements {
	//	resp[string(index)] = map[string]interface{}{"id": announce.ID}
	//}

	u.Respond(w, resp)
}

var AddAn = func(w http.ResponseWriter, r *http.Request) {
	announcement := &models.Announce{}
	err := json.NewDecoder(r.Body).Decode(announcement) //decode the request body into struct and failed if any error occur
	if err != nil {
		fmt.Println(err.Error())
		u.Respond(w, u.Message(false, "Required field filled incorrectly"))
		return
	}
	announcement = announcement.AddAn()//(r.Context().Value("AccountId"))

	resp := map[string]interface{}{"Message": "The announcement was published successfully"}

	u.Respond(w, resp)
}




