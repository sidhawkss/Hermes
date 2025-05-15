package middleware

import (
	"net/http"
	"crypto/subtle"
	"crypto/sha256"
)

var correctUser string = "ENVIRONMENTUSER";
var correctPass string = "ENVIRONMENTPASSWORD";

func Authentication(next http.HandlerFunc) http.HandlerFunc{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		user, pass, ok := r.BasicAuth();
		if(!ok){
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "", http.StatusUnauthorized)
		}
		
		userHash := sha256.Sum256([]byte(user));
		passHash := sha256.Sum256([]byte(pass));
		expectedUser := sha256.Sum256([]byte(correctUser));
		expectedPass := sha256.Sum256([]byte(correctPass));
		
		if(subtle.ConstantTimeCompare(expectedUser[:],userHash[:]) == 1 && subtle.ConstantTimeCompare(expectedPass[:],passHash[:]) == 1){
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "", http.StatusOK) // :)
		};

	})
}