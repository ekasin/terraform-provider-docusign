package docusign

import(
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("docusign_user.user1", "email", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "firstname", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "lastname", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "jobtitle", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "company", ""),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "docusign_user" "user1" {
  email        = ""
  firstname   = ""
  lastname    = ""
  jobtitle   = ""
  company    = ""
}
`)
}


func TestAccItem_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("docusign_user.user1", "email", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "firstname", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "lastname", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "jobtitle", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "company", ""),	
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("docusign_user.user1", "email", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "firstname", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "lastname", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "jobtitle", ""),
					resource.TestCheckResourceAttr("docusign_user.user1", "company", ""),
				),
			},
		},
	})
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
resource "docusign_user" "user1" {
	email        = ""
	firstname   = ""
	lastname    = ""
	jobtitle   = ""
	company    = ""
}
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
resource "docusign_user" "user1" {
	email        = ""
	firstname   = ""
	lastname    = ""
	jobtitle   = ""
	company    = ""
}
`)
}


