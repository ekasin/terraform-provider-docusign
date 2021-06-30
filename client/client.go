package client

import(
	"io"
	"fmt"
	"log"
	"time"
	"bytes"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Client struct {
	authToken  string
	accountId  string
	httpClient *http.Client
}

type User struct {
	Email   string   `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	JobTitle  string  `json:"jobTitle"`
	Company  string  `json:"company"`
	PermissionProfileName string `json:"permissionProfileName"`
}

type Country struct {
	Users []struct {
		UserName                   string    `json:"userName"`
		UserID                     string    `json:"userId"`
		UserType                   string    `json:"userType"`
		IsAdmin                    string    `json:"isAdmin"`
		UserStatus                 string    `json:"userStatus"`
		URI                        string    `json:"uri"`
		Email                      string    `json:"email"`
		CreatedDateTime            time.Time `json:"createdDateTime"`
		UserAddedToAccountDateTime time.Time `json:"userAddedToAccountDateTime"`
		FirstName                  string    `json:"firstName"`
		LastName                   string    `json:"lastName"`
		JobTitle                   string    `json:"jobTitle"`
		Company                    string    `json:"company"`
		PermissionProfileID        string    `json:"permissionProfileId"`
		PermissionProfileName      string    `json:"permissionProfileName"`
	} `json:"users"`
	ResultSetSize string `json:"resultSetSize"`
	TotalSetSize  string `json:"totalSetSize"`
	StartPosition string `json:"startPosition"`
	EndPosition   string `json:"endPosition"`
}

var (
    Errors = make(map[int]string)
)

func init() {
	Errors[400] = "Bad Request, StatusCode = 400"
	Errors[404] = "User Does Not Exist , StatusCode = 404"
	Errors[409] = "User Already Exist, StatusCode = 409"
	Errors[401] = "Unautharized Access, StatusCode = 401"
	Errors[429] = "User Has Sent Too Many Request, StatusCode = 429"
}

func NewClient(token string,accountid string) *Client {
	return &Client{
		authToken:  token,
		accountId:  accountid,
		httpClient: &http.Client{},
	}
}

func (c *Client) NewItem(item *User) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		log.Println("[CREATE ERROR]: ", err)
		return err
	}
	_, err = c.httpRequest("POST", buf, item)
	if err != nil {
		log.Println("[CREATE ERROR]: ", err)
		return err
	}
	return nil
}

func (c *Client) httpRequest(method string, body bytes.Buffer, item *User) (closer io.ReadCloser, err error) {
	data := fmt.Sprintf("{\"newUsers\":[{\"email\":\"%s\",\"userName\":\"%s\",\"firstName\":\"%s\",\"lastName\":\"%s\",\"jobTitle\":\"%s\",\"company\":\"%s\"}]}", item.Email, item.FirstName+" "+item.LastName,item.FirstName,item.LastName,item.JobTitle,item.Company)
	payload := strings.NewReader(data)
	req, err := http.NewRequest(method, fmt.Sprintf("https://demo.docusign.net/restapi/v2.1/accounts/%s/users",c.accountId), payload)
	authtoken := "Bearer "+ c.authToken
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return resp.Body, nil
    } else {
		return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
    }
}

func (c *Client) DeleteItem(email string) error {
	_, err := c.deletehttpRequest(fmt.Sprintf("%s", email), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) deletehttpRequest(email, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	name,err := userIdFunc(email,c.authToken,c.accountId)
	if (err != nil){
		return nil, err
	}
	data := fmt.Sprintf("{\"users\":[{\"userId\":\"%s\"}]}", name)
	payload := strings.NewReader(data)
	
	req, err := http.NewRequest(method, fmt.Sprintf("https://demo.docusign.net/restapi/v2.1/accounts/%s/users",c.accountId), payload)
	if err != nil {
		return nil, err
	}
	authtoken := "Bearer "+ c.authToken
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) UpdateItem(item *User) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.updatehttpRequest(fmt.Sprintf("%s", item.Email), "PUT", buf,item)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) updatehttpRequest(path,method string, body bytes.Buffer, item *User) (closer io.ReadCloser, err error) {
	data := fmt.Sprintf("{\"email\":\"%s\",\"userName\":\"%s\",\"firstName\":\"%s\",\"lastName\":\"%s\",\"jobTitle\":\"%s\",\"company\":\"%s\"}",item.Email,item.FirstName+" "+item.LastName,item.FirstName,item.LastName,item.JobTitle,item.Company)
	payload := strings.NewReader(data)
	name,err := userIdFunc(item.Email,c.authToken,c.accountId)
	if (err != nil){
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("https://demo.docusign.net/restapi/v2.1/accounts/%s/users/%s",c.accountId,name), payload)
	authtoken := "Bearer "+ c.authToken
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return resp.Body, nil
    } else {
		return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
    }
}

func (c *Client) GetUser(email string) (*User, error) {
	body, err := c.gethttpRequest(email, "GET", bytes.Buffer{})
	if err != nil{
		return nil, err
	}
	item := &User{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) gethttpRequest(email string, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	name,err := userIdFunc(email,c.authToken,c.accountId)
	if (err != nil){
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("https://demo.docusign.net/restapi/v2.1/accounts/%s/users/%s",c.accountId,name), &body)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	authtoken := "Bearer "+ c.authToken
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
		}
		return nil, fmt.Errorf("Error : %v ", Errors[resp.StatusCode])
	}
	return resp.Body, nil
}

func userIdFunc(email string,token string,accId string) (str string , err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://demo.docusign.net/restapi/v2.1/accounts/%s/users?email=%s",accId,email), nil)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return "",err
	}
	authtoken := "Bearer "+ token
	req.Header.Add("Authorization", authtoken)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return "", err
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		var country1 Country
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal([]byte(body), &country1)
		if err != nil {
			return "",err
		}
		if country1.Users == nil {
			return "",err
		}
		if len(country1.Users)==0 {
        	return "",err
    	} else {
        	return country1.Users[0].UserID ,nil
    	}
    } else {
		log.Println("Broken Request")
		return "", fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
    }
	
}

func (c *Client) IsRetry(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "429")==true {
			return true
		}
	}
	return false
}