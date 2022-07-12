package helper

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Global app context
type applicationContext struct {
	mssql *SQLServer
}

//Context type helper
var Context applicationContext

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
