package app

import (
	controller "MICROSERVICES/mvc1/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
