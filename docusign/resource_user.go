package docusign

import (
	"fmt"
	"time"
	"regexp"
	"context"
	"strings"
	"terraform-provider-docusign/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		if err = apiClient.NewItem(&user); err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	if err != nil {
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
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		user, err := apiClient.GetUser(userId)
		if err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
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
		return nil
	})
	if retryErr!=nil {
		if strings.Contains(retryErr.Error(), "User Does Not Exist")==true {
			d.SetId("")
			return diags
		}
		return diag.FromErr(retryErr)
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
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		if err = apiClient.UpdateItem(&user); err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceUserRead(ctx,d,m)
}

func resourceUserDelete(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	userId := d.Id()
	var err error
	retryErr := resource.Retry(2*time.Minute, func() *resource.RetryError {
		if err = apiClient.DeleteItem(userId); err != nil {
			if apiClient.IsRetry(err) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if retryErr != nil {
		time.Sleep(2 * time.Second)
		return diag.FromErr(retryErr)
	}
	if err != nil {
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