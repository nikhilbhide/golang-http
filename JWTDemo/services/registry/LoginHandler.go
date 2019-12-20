package registry

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nik/JWTDemo/bo"
	"github.com/nik/JWTDemo/model"
	"github.com/nik/JWTDemo/token"
	"github.com/nik/JWTDemo/utility"
	"github.com/urfave/negroni"
	"io/ioutil"
	"net/http"

)

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
	//populate the login type
	var login *model.Login
	json.NewDecoder(r.Body).Decode(&login)

	//invoke the login workflow
	login, err:= s.userCase.Login(nil,login)

	if(err!=nil) {
		utility.ResponseJSON(w, http.StatusBadRequest, err)
	} else {
		//check for error
		tokenString, error := token.GenerateToken(login)
		if (error == nil) {
			token := model.Token{Token: tokenString}
			utility.ResponseJSON(w, http.StatusCreated, token)
		}
	}
}

func (s *LoginHandler) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	//populate the login type
	fmt.Println("GET params were:", r.URL.Query())
	params := r.URL.Query()
	if (len(params) != 1) {
		utility.ResponseJSON(w, http.StatusBadRequest, "Bad URL")
	} else {
		email := params.Get("email")
		if (email == "") {
			utility.ResponseJSON(w, http.StatusBadRequest, "Email is missing")
		} else {
			login, error := s.userCase.FetchUserByEmail(email)
			if (error != nil) {
			} else {
				utility.ResponseJSON(w, http.StatusOK, login)
			}
		}
	}
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewLoginHandler(router *mux.Router,s bo.LoginUseCase) (func(w http.ResponseWriter, r *http.Request)) {
	handler := &LoginHandler{
		userCase: s,
	}

	router.HandleFunc("/login", handler.LoginUser)
	router.Handle("/fetchuser", negroni.New(
		negroni.HandlerFunc(token.GetJWTMiddleware().HandlerWithNext),
		negroni.WrapFunc(handler.getUserByEmail),
	))

	return handler.LoginUser
}