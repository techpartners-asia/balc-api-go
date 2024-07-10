package balcapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/techpartners-asia/balc-api-go/utils"
)

var (
	BalcLoan = utils.API{
		Url:    "/api?cust_id=",
		Method: http.MethodPost,
		Func:   "loanadv",
	}
	BalcLimit = utils.API{
		Url:    "/api?cust_id=",
		Method: http.MethodPost,
		Func:   "limitcheck",
	}
)

func (b *balc) httpRequest(body interface{}, api utils.API, customerId int) (response []byte, err error) {
	var requestByte []byte
	var requestBody *bytes.Reader
	if body == nil {
		requestBody = bytes.NewReader(nil)
	} else {
		requestByte, _ = json.Marshal(body)
		requestBody = bytes.NewReader(requestByte)
	}

	url := fmt.Sprintf(b.endpoint+api.Url+"%d", customerId)
	fmt.Println(url)
	req, _ := http.NewRequest(api.Method, url, requestBody)

	req.Header.Add("Content-Type", utils.HttpContent)
	req.Header.Add("Authorization", "Bearer "+b.token)
	req.Header.Add("func", api.Func)

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		response, _ = io.ReadAll(res.Body)
		errorString := string(response[:])
		return nil, errors.New(string(errorString))
	}
	defer res.Body.Close()
	response, _ = io.ReadAll(res.Body)
	return
}
