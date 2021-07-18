package main

import (
	"encoding/json"
	"net/url"
	"os"
	"time"

	duoapi "github.com/duosecurity/duo_api_golang"
)

type DuoAuthResponse struct {
	Response struct {
		Result    string `json:"result"`
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	} `json:"response"`
	Stat string `json:"stat"`
}

func main() {
	ikey := os.Getenv("INTEGRATION_KEY")
	skey := os.Getenv("SECRET_KEY")
	host := os.Getenv("API_HOSTNAME")
	userAgent := "duocheck"

	api := duoapi.NewDuoApi(
		ikey,
		skey,
		host,
		userAgent,
		duoapi.SetTimeout(30*time.Second),
	)

	username := os.Args[1]

	values := url.Values{}
	values.Set("username", username)
	values.Set("factor", "push")
	values.Set("device", "auto")

	_, body, err := api.SignedCall("POST", "/auth/v2/auth", values, duoapi.UseTimeout)
	if err != nil {
		print("error: " + err.Error())
		os.Exit(2)
	}

	var resp DuoAuthResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		print("error: " + err.Error())
		os.Exit(2)
	}

	if resp.Response.Result != "allow" {
		os.Exit(1)
	}
	os.Exit(0)
}
