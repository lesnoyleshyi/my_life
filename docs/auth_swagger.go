package docs

import (
	"my_life/internal/domain"
)

//  swagger:route POST /sign-in auth idOfSignIn
//  Receives JSON with login and password in body, responds with auth token.
//  responses:
//	    201: resp201
//      400: resp400
//      500: resp500

//  Server returns JSON with status "success" or "error" and relevant data (token or error description).
//  swagger:response resp201
type signInResponseWriter201 struct {
	// in:body
	Response domain.Response
}

//  Returns 400 error when can't find "name" or/and "password" values.
//  swagger:response resp400
type signInResponseWriter400 struct {
	// in:body
	Response domain.Response
}

//  Returns when faces some internal errors.
//  swagger:response resp500
type signInResponseWriter500 struct {
	// in:body
	Response domain.Response
}

// swagger:parameters idOfSignIn
type signInParamsWrapper struct {
	// Server wants username and password in HTTP body.
	// in:body
	Data domain.UsernamePasswd
}
