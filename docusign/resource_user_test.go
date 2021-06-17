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
					resource.TestCheckResourceAttr("docusign_user.user1", "email", "demo@gmail.com"),
					resource.TestCheckResourceAttr("docusign_user.user1", "firstname", "demoname"),
					resource.TestCheckResourceAttr("docusign_user.user1", "lastname", "singh"),
					resource.TestCheckResourceAttr("docusign_user.user1", "jobtitle", "manager"),
					resource.TestCheckResourceAttr("docusign_user.user1", "company", "demovalue"),
				),
			},
		},
	})
}

func testAccCheckItemBasic() string {
	return fmt.Sprintf(`
resource "docusign_user" "user1" {
  email        = "demo@gmail.com"
  firstname   = "demoname"
  lastname    = "singh"
  jobtitle   = "manager"
  company    = "demovalue"
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
					resource.TestCheckResourceAttr("docusign_user.user1", "email", "demo@gmail.com"),
					resource.TestCheckResourceAttr("docusign_user.user1", "firstname", "demoname"),
					resource.TestCheckResourceAttr("docusign_user.user1", "lastname", "singh"),
					resource.TestCheckResourceAttr("docusign_user.user1", "jobtitle", "manager"),
					resource.TestCheckResourceAttr("docusign_user.user1", "company", "demovalue"),	
				),
			},
			{
				Config: testAccCheckItemUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("docusign_user.user1", "email", "demo@gmail.com"),
					resource.TestCheckResourceAttr("docusign_user.user1", "firstname", "demoname"),
					resource.TestCheckResourceAttr("docusign_user.user1", "lastname", "singh"),
					resource.TestCheckResourceAttr("docusign_user.user1", "jobtitle", "manager"),
					resource.TestCheckResourceAttr("docusign_user.user1", "company", "demovalue"),
				),
			},
		},
	})
}

func testAccCheckItemUpdatePre() string {
	return fmt.Sprintf(`
resource "docusign_user" "user1" {
	email        = "demo@gmail.com"
	firstname   = "demoname"
	lastname    = "singh"
	jobtitle   = "manager"
	company    = "demovalue"
}
`)
}

func testAccCheckItemUpdatePost() string {
	return fmt.Sprintf(`
resource "docusign_user" "user1" {
	email        = "demo@gmail.com"
	firstname   = "demoname"
	lastname    = "singh"
	jobtitle   = "manager"
	company    = "demovalue"
}
`)
}


