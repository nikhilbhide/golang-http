package registry

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nik/JWTDemo/bo"
	"github.com/nik/JWTDemo/model"
	"github.com/nik/JWTDemo/utility"
	"io/ioutil"
	"net/http"
)

// NewArticleHandler will initialize the articles/ resources endpoint
func NewLoginHandler(router *mux.Router, s bo.LoginUseCase) {
	handler := &LoginHandler{
		userCase: s,
	}

	router.HandleFunc("/login", handler.LoginUser)
}

//http handler with json decoder to create instance of model
func SignUpHandlerV1(w http.ResponseWriter, r *http.Request) {
	//declare the instance variable and use json decoder to store the value into the same
	var signUpInstance model.Login
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
	signUpInstance := &model.Login{
	}

	//unmarshal json into the struct instance
	json.Unmarshal((body), signUpInstance)

	//print the instance attributes
	println("Email-V2:" ,signUpInstance.Email)
	println("UserID-V2",signUpInstance.UserID)

	fmt.Fprint(w, "POST done")
}
// Login ...
type LoginHandler struct {
	userCase bo.LoginUseCase
}

func SignUpHandlerV3(w http.ResponseWriter, r *http.Request) {
	//extract and store the form attributes into variables
	email := r.FormValue("email")
	userName := r.FormValue("userName")

	//instantiate the signup struct instance
	signUpInstance := &model.Login{
		UserID:userName,
		Email:email,
	}

	//print the instance attributes
	println("Email-V3:" ,signUpInstance.Email)
	println("Username-V3:", signUpInstance.UserID)

	fmt.Fprint(w, "POST done")
}

func (s *LoginHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	//invoke the login workflow
	login, err:= s.userCase.Login(nil,model.Login{})
	//check for error
	if(err.Message=="") {
		//in case of successful creation, respond with created
		utility.ResponseJSON(w, http.StatusCreated, login)
	} else {
		utility.ResponseJSON(w, http.StatusBadRequest, err)
	}
}