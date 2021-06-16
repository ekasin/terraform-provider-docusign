package docusign

import (
	"strings"
	"fmt"
	"terraform-provider-docusign/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
 )

 func dataSourceUser() *schema.Resource {
	return &schema.Resource{
	  Read: dataSourceUserRead,
	  Schema: map[string]*schema.Schema{
		"email": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"firstname": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"lastname": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"jobtitle": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"company": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"permissionprofilename": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
	  },
	}
  }
  
  func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
  	userId := d.Get("email").(string)
	user, err := apiClient.GetUser(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("user %s does not exist" , userId)
		}
	}
	d.SetId(user.Email)
	d.Set("email", user.Email)
	d.Set("firstname", user.FirstName)
	d.Set("lastname", user.LastName)
	d.Set("jobtitle", user.JobTitle)
  	d.Set("company", user.Company)
	d.Set("permissionprofilename", user.PermissionProfileName)
	return nil
  }
