package main

import (
	"fmt"
	"net/http"

	"github.com/JGetByte/go-rest-api/src/handlers/api/controllers"
)

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Welcome to my home")
}

func main() {
	fmt.Println(" -- start api --")

	http.HandleFunc("/", home)

	var controller = controllers.NewProfilesControllers()
	http.HandleFunc("/myProfile", controller.GetMyPorifle)

	http.ListenAndServe(":8080", nil)
}
