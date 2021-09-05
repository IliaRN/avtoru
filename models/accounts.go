package models

import (
	u "avtoru/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	IsAdmin  bool
}

//Validate incoming user details...
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() map[string]interface{} {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	ttl := time.Now().Add(5 * time.Hour).Unix()
	//tk := &Token{
	//	UserEmail: account.Email,
	//	ExpiresAt: ttl,
	//}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		account.Email, jwt.StandardClaims{
			ExpiresAt: ttl,
			IssuedAt:  time.Now().Unix(),
		},
	})
	//token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}

func Login(email, password string) map[string]interface{} {

	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
			return u.Message(false, "Invalid login credentials. Please try again")
		}
		return u.Message(false, "Internal server error")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	//tk := &Token{UserEmail: account.Email}
	ttl := time.Now().Add(5 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Token{
		account.Email, jwt.StandardClaims{
			ExpiresAt: ttl,
			IssuedAt:  time.Now().Unix(),
		},
	})
	//token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
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

func GetAccountByCredentials(email, password string) *Account {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil
	}

	return account
}
