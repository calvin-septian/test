package helper

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
	"training/entity"

	jwt "github.com/golang-jwt/jwt/v4"
)

// Global app context
type applicationContext struct {
	mssql *SQLServer
}

//Context type helper
var (
	Context applicationContext
	jwtKey  = []byte("test_jwt")
)

func ConnectDB() {
	sql := NewSQLConnection()
	Context.mssql = sql
}

func CloseConnectionDB() {
	Context.mssql.LocalDB.Close()
}

func requestHttp(url, method string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	var netClient = &http.Client{}
	resp, err := netClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	return data, nil
}

func GenerateJWT(Username string) (string, error) {
	tokenMethod := jwt.SigningMethodHS256
	claims := entity.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
		Username: Username,
	}
	token := jwt.NewWithClaims(tokenMethod, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
