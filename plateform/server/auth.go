package server

import (
	"log"
	"net/http"
	"time"

	"github.com/alias-asso/iosu/plateform/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func generateJWT(username string, config *config.Config) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JwtKey))

}

func (s *Server) getLogin(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) postLogin(w http.ResponseWriter, r *http.Request) {

	// TODO: verify username/password
	token, err := generateJWT(r.FormValue("username"), s.cfg)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(time.Hour * 4),
	})

	// w.Header().Set("HX-Redirect", "/admin")
	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
