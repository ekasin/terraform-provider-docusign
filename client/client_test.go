package client

import(
	"github.com/stretchr/testify/assert"
	"testing"
	"log"
	"os"
)

func init() {
	os.Setenv("DOCUSIGN_TOKEN", "demovalue")
	os.Setenv("DOCUSIGN_ACCOUNT_ID","demovalue")
}

func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		seedData     map[string]User
		expectErr    bool
		expectedResp *User
	}{
		{
			testName: "user exists",
			itemName: "user@gmail.com",
			seedData: map[string]User{
				"user@gmail.com": {
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "DocuSign Sender",
				},
			},
			expectErr: false,
			expectedResp: &User{
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "",
			},
		},
		
		{
			testName:     "user does not exist",
			itemName:     "user@gmail.com",
			seedData:     nil,
			expectErr:    true,
			expectedResp: nil,
		},
		
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("DOCUSIGN_TOKEN"),os.Getenv("DOCUSIGN_ACCOUNT_ID"))

			item, err := client.GetUser(tc.itemName)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}

func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *User
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &User{
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "",
			},
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "item already exists",
			newItem: &User{
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "",
			},
			seedData: map[string]User{
				"item1": {
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "DocuSign Sender",
				},
			},
			expectErr: false,
		},
		
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("DOCUSIGN_TOKEN"),os.Getenv("DOCUSIGN_ACCOUNT_ID"))
			err := client.NewItem(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetUser(tc.newItem.Email)
			assert.NoError(t, err)
			assert.Equal(t, tc.newItem, item)
		})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedItem *User
		seedData    map[string]User
		expectErr   bool
	}{
		{
			testName: "item exists",
			updatedItem: &User{
				Email:   "user@gmail.com",
				FirstName: "demoname",
				LastName: "singh",
				JobTitle: "manager",
				Company: "democompany",
				PermissionProfileName: "",
			},
			seedData: map[string]User{
				"item1": {
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "",
				},
			},
			expectErr: false,
		},
		{
			testName: "item does not exist",
			updatedItem: &User{
					Email:   "demo@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "",
			},
			seedData:  nil,
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("DOCUSIGN_TOKEN"),os.Getenv("DOCUSIGN_ACCOUNT_ID"))
			err := client.UpdateItem(tc.updatedItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetUser(tc.updatedItem.Email)
			assert.NoError(t, err)
			assert.Equal(t, tc.updatedItem, item)
		})
	}
}

func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemName  string
		seedData  map[string]User
		expectErr bool
	}{
		{
			testName: "user exists",
			itemName: "user@gmail.com",
			seedData: map[string]User{
				"user1": {
					Email:   "user@gmail.com",
					FirstName: "demoname",
					LastName: "singh",
					JobTitle: "manager",
					Company: "democompany",
					PermissionProfileName: "",
				},
			},
			expectErr: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient(os.Getenv("DOCUSIGN_TOKEN"),os.Getenv("DOCUSIGN_ACCOUNT_ID"))
			err := client.DeleteItem(tc.itemName)
			if tc.expectErr {
				log.Println("[DELETE ERROR]: ", err)
				assert.Error(t, err)
				return
			}
			_, err = client.GetUser(tc.itemName)
			log.Println("[DELETE ERROR]: ", err)
			assert.Error(t, err)
		})
	}
}
