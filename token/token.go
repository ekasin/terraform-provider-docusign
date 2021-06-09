package token

import (
	"encoding/base64"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"log"
)

func GenerateToken(secretKey string, integrationKey string,refreshToken string) string {
	msg := integrationKey + ":" + secretKey
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	data := url.Values{}
	data.Set("grant_type","refresh_token")
    data.Set("refresh_token", refreshToken)
	req, err := http.NewRequest("POST","https://account-d.docusign.com/oauth/token", strings.NewReader(data.Encode()))
	authtoken := "Basic "+ encoded
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println("[ERROR]: ",err)
		
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[ERROR]: ",err)
	}
	var temp map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &temp)
    if err != nil {
        panic(err)
    }
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		tokenString := temp["access_token"].(string)
		return tokenString
    } else {
		log.Println("refresh token expired")
		return ""	
    }
}