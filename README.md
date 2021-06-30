
This terraform provider allows to perform Create ,Read ,Update, Delete and Import docusign User(s). 


## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Docusign](https://developers.docusign.com/) Developers account
* [Docusign Rest API](https://developers.docusign.com/docs/esign-rest-api/reference/) 


## Application Setup
 ***This provider can be successfully tested on any dcosign developer account.*** <br><br>

1. Create a docusign account with your required subscription (PRO Plan/Business Account). (https://www.docusign.com/products-and-pricing)<br>
2. Sign in to the docusign account (https://developers.docusign.com/)<br>
3. Go to [Apps and Keys](https://admindemo.docusign.com/apps-and-keys)<br>
4. Click on `Add App and integration key`. For our purpose we need to create an Auth Token. <br>
5. Follow this [Generate Auth Token](https://developers.docusign.com/platform/auth/authcode/authcode-get-token/) website to make an app. <br>
This app will provide us with the token and accountid which will be needed to configure our provider and make request. <br>


## Building the Providers 
1. Clone the repository  to $GOPATH/src/github.com/docusign/terraform-provider-docusign <br>
2. Add the Refresh token generted in  Docusign App to respective fields as shown in example usage <br>
3. Run the following command :
 ```golang
go mod init terraform-provider-docusign
go mod tidy
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Managing plugins for terraform
1. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
Command: 
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/docusign/0.2.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/terraform/provider/docusign/0.2.0/windows_amd64`<br>

2. Run `go build -o terraform-provider-docusign.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to appropriate location.
 ```
 move terraform-provider-docusign.exe %APPDATA%\terraform.d\plugins\terraform\provider\docusign\0.2.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Working with terraform

#### Create User
1. Add the user email, first name, last name, job_title, company in the respective field in resourcse block[refer example usage](#example-usage)
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an account activation mail has been sent to the user.
5. Activate the account using the link provided in the mail.

#### Update the user
Update the data of the user in the resourcse block and apply using `terraform apply`

#### Read the User Data
Add data and output blocks and run `terraform plan` to read user data


#### Delete the user
Delete the resource block of the particular user file and run `terraform apply` or run `terraform destroy`.

#### Import a User Data
1. Write manually a resource configuration block for the User(refer resourcse block), to which the imported object will be mapped.
2. Run the command `terraform import docusign_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.


## Example Usage <a id="example-usage"></a>
```terraform
terraform {
  required_providers {
    docusign = {
      version = "0.2"
      source  = "terraform/provider/docusign"
    }
  }
}

provider "docusign" {
  secretkey = "rwetwuytrqweyr"
  integrationkey = "gdfjhsdfjdf"
  refresh token = "dgfshdgfsdh"
  accountid = "dhfhglsdhfljsdh"
}

resource "docusign_user" "user1" {
   email      = "user@gmail.com"
   firstname = "usersfirstname"
   lastname  = "userslastname"
   jobtitle  = "engineer"
   company = "clevertap"
}

data "docusign_user" "user1" {
  email = "user@gmail.com"
}

output "user1" {
  value = data.docusign_user.user1
}
```

## Argument Reference

* `secretkey`(Required, String)       - The Docusing secret Key from created application (Set DOCUSIGN_SECRET_KEY environment variable)
* `integrationkey`(Required, String)  - The Docusing integration Key from created application (Set DOCUSIGN_INTEGRATION_KEY environment variable)
* `refreshtoken`(Required, String)    - refresh token for generation of new access token.(Set DOCUSIGN_REFRESH_TOKEN environment variable)
* `accountid`(Required, String)       - Account id of the Admin. (Set DOCUSIGN_ACCOUNT_ID environment variable)
* `first_name`(Required, String)      - First name of the User.
* `last_name`(Required, String)       - Last Name / Family Name / Surname of the User.
* `job_title`(Optional, String)       - Job title of the particular user.
* `company`(Optional, String)         - Company of the particular user.
* `permissionprofilename`(computed, String)  - profile of the particular user ie.(DS Admin,DS Sender,DS Viewer)Can be managed only through docusign dashboard.

