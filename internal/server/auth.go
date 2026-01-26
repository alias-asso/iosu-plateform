package server

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/mail"
	"slices"
	"time"

	"github.com/alias-asso/iosu/internal/config"
	"github.com/alias-asso/iosu/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateActivationCode() string {
	return randSeq(32)
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateUsername(username string) bool {
	return len(username) <= 16
}

type UserRequest struct {
	username string
	email    string
}

func (s *Server) createDeactivatedAccount(ctx context.Context, userInfos UserRequest) error {
	user := database.User{Username: userInfos.username, Email: userInfos.email, Activated: false, Admin: false}

	activationCode := database.ActivationCode{Code: generateActivationCode(), Expiration: time.Now(), User: user}
	err := gorm.G[database.ActivationCode](s.db).Create(ctx, &activationCode)
	return err
}

// route handler
func (s *Server) getLogin(w http.ResponseWriter, r *http.Request) {
}

// route handler
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
	// TODO: replace with ok template
	w.Write([]byte("ok"))
}

// route handler
func (s *Server) postRegisterAccount(w http.ResponseWriter, r *http.Request) {

}

// route handler
func (s *Server) postBatchCreateAccounts(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("accounts")
	if err != nil {
		http.Error(w, "Error retrieving file.", http.StatusUnsupportedMediaType)
		log.Println(err)
		return
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvHeader, err := csvReader.Read()
	if err != nil {
		http.Error(w, "Error reading CSV header.", http.StatusUnsupportedMediaType)
		log.Println(err)
		return
	}
	var headers = [2]string{"username", "email"}
	var indexes = make([]int, len(headers))
	for i, h := range headers {
		index := slices.Index(csvHeader, h)
		if index == -1 {
			http.Error(w, "Invalid CSV header.", http.StatusUnsupportedMediaType)
			return
		}
		indexes[i] = index
	}
	lines, err := csvReader.ReadAll()

	if err != nil {
		http.Error(w, "Error reading CSV.", http.StatusUnsupportedMediaType)
		log.Println(err)
		return
	}
	users := make([]UserRequest, len(lines))
	for i, line := range lines {
		username := line[indexes[0]]
		if !validateUsername(username) {
			http.Error(w, fmt.Sprintf("Invalid username on line %d.", i), http.StatusUnsupportedMediaType)
			return
		}
		email := line[indexes[1]]
		if !validateEmail(email) {
			http.Error(w, fmt.Sprintf("Invalid email on line %d.", i), http.StatusUnsupportedMediaType)
			return
		}
		users[i] = UserRequest{
			username: username,
			email:    email,
		}
	}
	for _, user := range users {
		err := s.createDeactivatedAccount(context.Background(), user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to create account %s.", user.username), http.StatusUnsupportedMediaType)
			log.Println(err)
			return
		}
	}
	// TODO: replace with ok template
	w.Write([]byte("ok"))
}
