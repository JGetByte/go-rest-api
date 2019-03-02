package controllers

import (
	"encoding/json"
	"net/http"
)

type ProfilesController struct {
	resource string
}

func NewProfilesControllers() *ProfilesController {
	return &ProfilesController{
		resource: "proifles",
	}
}

type myProfile struct {
	Firstname  string
	Lastname   string
	Occupation string
	Departure  string
}

func (c *ProfilesController) GetMyPorifle(res http.ResponseWriter, req *http.Request) {
	p := myProfile{
		Firstname:  "Steve",
		Lastname:   "Rogers",
		Occupation: "Hero",
		Departure:  "Mavel",
	}
	json.NewEncoder(res).Encode(p)
}
