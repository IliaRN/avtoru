package app

import (
	"avtoru/models"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"POST:/api/user/new", "POST:/api/user/login"} //List of endpoints that doesn't require auth
		requestPath := r.Method + ":" + r.URL.Path                        //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			http.Error(w, "Token is missing", http.StatusUnauthorized)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			http.Error(w, "Wrong token format", http.StatusUnauthorized)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			http.Error(w, "Malformed auth token", http.StatusUnauthorized)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			http.Error(w, "Malformed auth token", http.StatusUnauthorized)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token

		acc := &models.Account{}
		models.GetDB().Find(acc, "email = ?", tk.UserEmail)
		if len(acc.Email) == 0 {
			http.Error(w, "Account not found", http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "user", acc.ID)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
