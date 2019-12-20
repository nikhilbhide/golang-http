package docs

import "github.com/nik/JWTDemo/model"

// swagger:route POST /login login-tag idOfLoginEndpoint
// Verifies user login credentials
// responses:
//   200: LoginResponse

// This text will appear as description of your response body.
// swagger:response LoginResponse
type LoginResponseWrapper struct {
	// in:body
	Body model.Login
}

// swagger:parameters idOfFoobarEndpoint
type LoginRequestWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body model.Login
}