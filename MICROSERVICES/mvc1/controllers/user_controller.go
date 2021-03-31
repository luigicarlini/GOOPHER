package controller

import (
	"MICROSERVICES/mvc1/services"
	"MICROSERVICES/mvc1/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	UserIdParam := req.URL.Query().Get("user_id")
	log.Printf("about to process user_id %v", UserIdParam)
	UserId, err := strconv.ParseInt(UserIdParam, 10, 64)
	if err != nil {
		// resp.WriteHeader(http.StatusNotFound)
		// resp.Write([]byte("user_id must be a number"))
		//Just return the bed request to the client
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(UserId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		//it isn’t necessary to dereference structures to access their fields. The previous listing is preferable
		//to writing resp.WriteHeader((*apiErr).StatusCode)
		resp.Write(jsonValue)
		//handle the error and return to the client
		return
	}
	//return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
