package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
JWT claims struct
*/
type Token struct {
	UserEmail string
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `gorm:"-"`
}

//Validate incoming user details...
func (account *Account) Validate() (int, error) {

	if !strings.Contains(account.Email, "@") {
		return http.StatusBadRequest, errors.New("email address is required")
	}

	if len(account.Password) < 6 {
		return http.StatusBadRequest, errors.New("password is required")
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, errors.New("connection error")
	}
	if temp.Email != "" {
		return http.StatusBadRequest, errors.New("email address is used by another user")
	}

	return http.StatusOK, nil
}

func (account *Account) Create() (int, error) {

	if status, err := account.Validate(); err != nil {
		return status, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return http.StatusInternalServerError, errors.New("failed to connect to the server")
	}

	//Create new JWT token for the newly registered account
	//ttl := time.Now().Add(5 * time.Minute).Unix()
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
	//	account.Email, jwt.StandardClaims{
	//		ExpiresAt: ttl,
	//		IssuedAt:  time.Now().Unix(),
	//	},
	//})
	//
	//tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	//account.Token = tokenString

	//account.Password = "" //delete password

	return http.StatusOK, nil
}

func Login(email, password string) (string, error, int) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("email address not found"), http.StatusNotFound

		}
		return "", errors.New("connection error. please retry"), http.StatusInternalServerError
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			return "", errors.New("invalid login credentials. please try again"), http.StatusBadRequest
		}
		return "", errors.New("internal server error"), http.StatusInternalServerError
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	//tk := &Token{UserEmail: account.Email}
	ttl := time.Now().Add(30 * time.Second).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		account.Email, jwt.StandardClaims{
			ExpiresAt: ttl,
			IssuedAt:  time.Now().Unix(),
		},
	})
	//token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return tokenString, nil, http.StatusOK
}

func GetAccountById(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
