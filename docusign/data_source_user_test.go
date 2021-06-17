package docusign

import (
	"fmt"
	"testing"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.docusign_user.user1", "email", "demo@gmail.com"),
				),
			},
		},
	})
}

func testAccUserDataSourceConfig() string {
	return fmt.Sprintf(`	  
	resource "docusign_user" "user1" {
		email        = "demo@gmail.com"
		firstname   = "demoname"
		lastname    = "singh"
		jobtitle   = "manager"
  		company    = "democompany"
	  }
	data "docusign_user" "user1" {
		email = "demo@gmail.com"
	}
	`)
}