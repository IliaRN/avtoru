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
		http.Error(w, "ID is missing in URL string", http.StatusBadRequest)
		return
	}
	id_conv, _ := strconv.ParseUint(id, 10, 32)
	announcementModel := models.GetAnnById(uint(id_conv))
	compare := announcementModel.AccountID
	if compare <= 0 {
		http.Error(w, "Announcement not found", http.StatusNotFound)
		return
	}

	u.Respond(w, announcementModel)
}

var decoder = schema.NewDecoder()

var GetAnns = func(w http.ResponseWriter, r *http.Request) {
	var filterStruct f.FilterStruct
	err := decoder.Decode(&filterStruct, r.URL.Query())
	if err != nil {
		http.Error(w, "Error in GET parameters", http.StatusBadRequest)
		return
	}

	announcements := models.GetAnns(filterStruct)
	u.Respond(w, announcements)
}

var AddAn = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	announcement := &models.Announce{}
	auto := &models.Auto{}
	err := json.NewDecoder(r.Body).Decode(announcement) //decode the request body into struct and failed if any error occur
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Required field filled incorrectly", http.StatusBadRequest)
		return
	}
	announcement.AccountID = accountId

	if announcement.Name == "" {
		http.Error(w, "Name is required field", http.StatusBadRequest)
		return
	}
	if announcement.Price <= 0 {
		http.Error(w, "Price is required field", http.StatusBadRequest)
		return
	}
	if announcement.Description == "" {
		log.Println("Description is empty but it's ok")
	}

	if announcement.Auto.Year <= 0 {
		http.Error(w, "Year is required field", http.StatusBadRequest)
		return
	} else {
		auto.Year = announcement.Auto.Year
	}
	if announcement.Auto.Mileage <= 0 {
		http.Error(w, "Mileage is required field", http.StatusBadRequest)
		return
	} else {
		auto.Mileage = announcement.Auto.Mileage
	}
	if announcement.Auto.ModelID <= 0 {
		http.Error(w, "ModelID is required field", http.StatusBadRequest)
		return
	} else {
		auto.ModelID = announcement.Auto.ModelID
	}
	auto.Categories = announcement.Auto.Categories
	auto = auto.AddAuto()
	announcement.Auto.ID = auto.ID
	announcement = announcement.AddAn() //

	u.Respond(w, "")
}

var DelAn = func(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]
	if !ok || len(id[0]) < 1 {
		http.Error(w, "ID is missing in URL string", http.StatusBadRequest)
		return
	}

	id_conv, _ := strconv.ParseUint(id[0], 10, 32)
	announcementModel := models.GetAnnById(uint(id_conv))
	compare := announcementModel.AccountID
	accountId := r.Context().Value("user").(uint)
	log.Println(r.Context()) //
	if compare <= 0 {
		http.Error(w, "Announcement not found", http.StatusNotFound)
		return
	}
	if accountId != compare {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	announcementModel.DelAn(announcementModel.ID)

	u.Respond(w, "")
}

var UpdAn = func(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value("user").(uint)
	announcement := &models.Announce{}
	err := json.NewDecoder(r.Body).Decode(announcement)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Required field filled incorrectly", http.StatusBadRequest)
		return
	}
	annToUpd := models.GetAnnById(announcement.ID)
	if annToUpd.ID <= 0 {
		http.Error(w, "Announcement not found", http.StatusNotFound)
		return
	}
	if accountId != annToUpd.AccountID {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}
	annToUpd.Description = announcement.Description
	annToUpd.Name = announcement.Name
	annToUpd.Price = announcement.Price
	announcement.UpdAn(*annToUpd)
	u.Respond(w, "")

}
