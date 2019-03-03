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
	firstname  string
	lastname   string
	occupation string
	departure  string
}

func (c *ProfilesController) GetMyPorifle(res http.ResponseWriter, req *http.Request) {
	p := myProfile{
		firstname:  "Steve",
		lastname:   "Rogers",
		occupation: "Hero",
		departure:  "Mavel",
	}
	json.NewEncoder(res).Encode(p)
}
