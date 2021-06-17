package docusign

import(
	"os"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
	"log"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	os.Setenv("DOCUSIGN_SECRET_KEY", "demovalue")
	os.Setenv("DOCUSIGN_INTEGRATION_KEY", "demovalue")
	os.Setenv("DOCUSIGN_REFRESH_TOKEN", "demovalue")
	os.Setenv("DOCUSIGN_ACCOUNT_ID","demovalue")
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"docusign": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		log.Println("[ERROR]: ",err)
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T)  {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("DOCUSIGN_SECRET_KEY"); v == "" {
		t.Fatal("DOCUSIGN_SECRET_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("DOCUSIGN_INTEGRATION_KEY"); v == "" {
		t.Fatal("DOCUSIGN_INTEGRATION_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("DOCUSIGN_REFRESH_TOKEN"); v == "" {
		t.Fatal("DOCUSIGN_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("DOCUSIGN_ACCOUNT_ID"); v == "" {
		t.Fatal("DOCUSIGN_ACCOUNTID must be set for acceptance tests")
	}
}