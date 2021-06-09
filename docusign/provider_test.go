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
	os.Setenv("DOCUSIGN_ACCOUNTID","")
	os.Setenv("DOCUSIGN_TOKEN", "")
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
	if v := os.Getenv("DOCUSIGN_ACCOUNTID"); v == "" {
		t.Fatal("DOCUSIGN_ACCOUNTID must be set for acceptance tests")
	}

	if v := os.Getenv("DOCUSIGN_TOKEN"); v == "" {
		t.Fatal("DOCUSIGN_TOKEN must be set for acceptance tests")
	}
}