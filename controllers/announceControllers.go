package controllers

import (
	f "avtoru/helpers"
	"avtoru/models"
	u "avtoru/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"strconv"
)

var GetAnnById = func(w http.ResponseWriter, r *http.Request) {
	//id, ok := r.URL.Query()["id"]
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID is missing in URL string", 422)
		return
	}
	id_conv, _ := strconv.ParseUint(id, 10, 32)
	announcementModel := models.GetAnnById(uint(id_conv))
	resp := u.Message(true, "Success")
	resp["data"] = announcementModel
	u.Respond(w, resp)
}

var decoder = schema.NewDecoder()

var GetAnns = func(w http.ResponseWriter, r *http.Request) {
	var filterStruct f.FilterStruct
	err := decoder.Decode(&filterStruct, r.URL.Query())
	if err != nil {
		u.Respond(w, u.Message(false, "Error in GET parameters"))
	}

	announcements := models.GetAnns(filterStruct)
	resp := map[string]interface{}{"data": announcements}
	u.Respond(w, resp)
}

var AddAn = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	announcement := &models.Announce{}
	announcement.AccountID = accountId
	auto := &models.Auto{}
	err := json.NewDecoder(r.Body).Decode(announcement) //decode the request body into struct and failed if any error occur
	if err != nil {
		fmt.Println(err.Error())
		u.Respond(w, u.Message(false, "Required field filled incorrectly"))
		return
	}
	announcement.AccountID = accountId

	if announcement.Name == "" {
		http.Error(w, "Name is required field", 422)
		return
	}
	if announcement.Price == 0 {
		http.Error(w, "Price is required field", 422)
		return
	}
	if announcement.Description == "" {
		log.Println("Description is empty but it's ok")
	}

	if announcement.Auto.Year == 0 {
		http.Error(w, "Year is required field", 422)
		return
	} else {
		auto.Year = announcement.Auto.Year
	}
	if announcement.Auto.Mileage == 0 {
		http.Error(w, "Mileage is required field", 422)
		return
	} else {
		auto.Mileage = announcement.Auto.Mileage
	}
	if announcement.Auto.ModelID == 0 {
		http.Error(w, "ModelID is required field", 422)
		return
	} else {
		auto.ModelID = announcement.Auto.ModelID
	}
	auto.Categories = announcement.Auto.Categories
	auto = auto.AddAuto()
	announcement.Auto.ID = auto.ID
	announcement = announcement.AddAn() //

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
	accountId := r.Context().Value("user").(uint)
	log.Println(r.Context()) //
	if accountId != compare {
		u.Respond(w, u.Message(false, "Access denied"))
		return
	}

	result := announcement.DelAn(announcement.ID)

	resp := map[string]interface{}{"Message": "The announcement was deleted", "result": result}
	u.Respond(w, resp)
}

var UpdAn = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	announcement := &models.Announce{}
	err := json.NewDecoder(r.Body).Decode(announcement)
	if err != nil {
		fmt.Println(err.Error())
		u.Respond(w, u.Message(false, "Required field filled incorrectly"))
		return
	}
	annToUpd := models.GetAnnById(announcement.ID)
	if accountId != annToUpd.AccountID {
		u.Respond(w, u.Message(false, "Access denied"))
		return
	}
	annToUpd.Description = announcement.Description
	annToUpd.Name = announcement.Name
	annToUpd.Price = announcement.Price
	result := announcement.UpdAn(*annToUpd)
	resp := map[string]interface{}{"Message": "The announcement was updated", "result": result}
	u.Respond(w, resp)

}
