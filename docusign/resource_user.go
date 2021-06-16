package docusign

import (
	"terraform-provider-docusign/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"fmt"
	"regexp"
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validateEmail(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value := v.(string)
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !(emailRegex.MatchString(value)) {
		errs = append(errs, fmt.Errorf("Expected EmailId is not valid  %s", k))
		return warns, errs
	}
	return
}

func validateValue(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	if value == "" {
		errs = append(errs, fmt.Errorf("value not to be empty"))
		return warns, errs
	}
	return
}

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceUserImporter,
		},
		Schema: map[string]*schema.Schema{

			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validateEmail,
			},
			"firstname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validateValue,
			},
			"lastname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validateValue,
			},
			"jobtitle": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validateValue,
			},
			"company": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validateValue,
			},
			"permissionprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}

}

func resourceUserCreate(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	user := client.User{
		Email:   d.Get("email").(string),
		FirstName: d.Get("firstname").(string),
		LastName: d.Get("lastname").(string),
		JobTitle: d.Get("jobtitle").(string),
		Company: d.Get("company").(string),

	}
	err := apiClient.NewItem(&user)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return diag.FromErr(err)
	}
	d.SetId(user.Email)
	resourceUserRead(ctx,d,m)
	return diags
}

func resourceUserRead(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	user, err := apiClient.GetUser(userId)
	if err != nil {
		return diags
	}
	if len(user.Email) > 0{
		d.SetId(user.Email)
		d.Set("email", user.Email)
		d.Set("firstname", user.FirstName)
		d.Set("lastname", user.LastName)
		d.Set("jobtitle", user.JobTitle)
		d.Set("company", user.Company)
		d.Set("permissionprofilename", user.PermissionProfileName)
	}
	return diags
}

func resourceUserUpdate(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var _ diag.Diagnostics
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})

		return diags
	}
	user := client.User{
		Email:   d.Get("email").(string),
		FirstName: d.Get("firstname").(string),
		LastName: d.Get("lastname").(string),
		JobTitle: d.Get("jobtitle").(string),
		Company: d.Get("company").(string),
	}
	err := apiClient.UpdateItem(&user)
	if err != nil {
		log.Printf("[Error] Error updating user :%s", err)
		return diag.FromErr(err)
	}
	return resourceUserRead(ctx,d,m)
}

func resourceUserDelete(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	err := apiClient.DeleteItem(userId)
	if err != nil {
		log.Printf("[Error] Error deleting user :%s", err)
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

func resourceUserImporter(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	apiClient := m.(*client.Client)
	userId := d.Id()
	user, err := apiClient.GetUser(userId)
	if err != nil {
		return nil,err
	}
	if len(user.Email) > 0{
		d.SetId(user.Email)
		d.Set("email", user.Email)
		d.Set("firstname", user.FirstName)
		d.Set("lastname", user.LastName)
		d.Set("jobtitle", user.JobTitle)
		d.Set("company", user.Company)
		d.Set("permissionprofilename", user.PermissionProfileName)
	}
	return []*schema.ResourceData{d}, nil
}