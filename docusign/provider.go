package docusign

import (
	"terraform-provider-docusign/client"
	tkn "terraform-provider-docusign/token"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secretkey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DOCUSIGN_SECRET_KEY", ""),
			},
			"integrationkey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DOCUSIGN_INTEGRATION_KEY", ""),
			},
			"refreshtoken": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DOCUSIGN_REFRESH_TOKEN", ""),
			},
			"accountid": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("DOCUSIGN_ACCOUNT_ID", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"docusign_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"docusign_user": dataSourceUser(),
		},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	token := tkn.GenerateToken(d.Get("secretkey").(string),d.Get("integrationkey").(string),d.Get("refreshtoken").(string))
	accountid := d.Get("accountid").(string)
	return client.NewClient(token,accountid), nil

}