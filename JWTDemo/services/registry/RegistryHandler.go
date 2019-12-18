package registry

import (
	"encoding/json"
	"fmt"
	"github.com/nik/JWTDemo/model"
	"io/ioutil"
	"net/http"
)

//http handler with json decoder to create instance of model
func SignUpHandlerV1(w http.ResponseWriter, r *http.Request) {
	//declare the instance variable and use json decoder to store the value into the same
	var signUpInstance model.Signup
	json.NewDecoder(r.Body).Decode(&signUpInstance)

	//print the instance attributes
	println("Email-V1:" ,signUpInstance.Email)
	println("UserID-V1:",signUpInstance.UserID)

	fmt.Fprint(w, "POST done")
}

//http handler by json unmarshal method to create instance of model
func SignUpHandlerV2(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	//instantiate the struct instance
	signUpInstance := &model.Signup{
	}

	//unmarshal json into the struct instance
	json.Unmarshal((body), signUpInstance)

	//print the instance attributes
	println("Email-V2:" ,signUpInstance.Email)
	println("UserID-V2",signUpInstance.UserID)

	fmt.Fprint(w, "POST done")
}

func SignUpHandlerV3(w http.ResponseWriter, r *http.Request) {
	//extract and store the form attributes into variables
	email := r.FormValue("email")
	userName := r.FormValue("userName")

	//instantiate the signup struct instance
	signUpInstance := &model.Signup{
		UserID:userName,
		Email:email,
	}

	//print the instance attributes
	println("Email-V3:" ,signUpInstance.Email)
	println("Username-V3:", signUpInstance.UserID)

	fmt.Fprint(w, "POST done")
}