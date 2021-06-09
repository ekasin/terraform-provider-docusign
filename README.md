# Terraform Docusign Provider

This terraform provider allows to perform Create ,Read ,Update, Delete and Import docusign User(s). 


## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Docusign](https://developers.docusign.com/) Developers account 


## Docusign Setup
 :heavy_exclamation_mark:  [IMPORTANT] : This provider can be successfully tested on any dcosign developer account. <br><br>

1. Create a docusign account with your required subscription (PRO Plan/Business Account). (https://www.docusign.com/products-and-pricing)<br>
2. Sign in to the docusign account (https://developers.docusign.com/)<br>
3. Go to [Apps and Keys](https://admindemo.docusign.com/apps-and-keys)<br>
4. Click on `Add App and integration key`. For our purpose we need to create an Auth Token. <br>
5. Follow this [Generate Auth Token](https://developers.docusign.com/platform/auth/authcode/authcode-get-token/) website to make an app. <br>
This app will provide us with the token and accountid which will be needed to configure our provider and make request. <br>


## Initialise Docusign Provider in local machine 
1. Clone the repository  to $GOPATH/src/github.com/docusign/terraform-provider-docusign <br>
2. Add the Refresh token generted in  Docusign App to respective fields in `main.tf` <br>
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
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/docusign/0.2.0/windows_amd64`<br>

2. Run `go build -o terraform-provider-docusign.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to appropriate location.
 ```
 move terraform-provider-docusign.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\docusign\0.2.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


[OR]

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Working with terraform

#### Create User
1. Add the user email, first name, last name, job_title, company in the respective field in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an account activation mail has been sent to the user.
5. Activate the account using the link provided in the mail.

#### Update the user
Update the data of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User Data
Add data and output blocks in the `main.tf` file and run `terraform plan` to read user data

#### Activate/Deactivate the user
No API for activate,deactive.

#### Delete the user
Delete the resource block of the particular user from `main.tf` file and run `terraform apply`.

#### Import a User Data
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped.
2. Run the command `terraform import docusign_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` . This command will give combined test result for the execution or errors if any failure occur.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


## Example Usage
```terraform
terraform {
  required_providers {
    docusign = {
      version = "0.2"
      source  = "hashicorp.com/edu/docusign"
    }
  }
}

provider "docusign" {
  secretkey = ""
  integrationkey = ""
  refresh token = ""
  accountid = ""
}

resource "docusign_user" "user1" {
   email      = "[EMAIL_ID]"
   firstname = "[USER_FIRST_NAME]"
   lastname  = "[USER_LAST_NAME]"
   jobtitle  = "[USER_JOB_TITLE]"
   company = "[USER_COMPANY]"
}

data "docusign_user" "user1" {
  email = "[EMAIL_ID]"
}

output "user1" {
  value = data.docusign_user.user1
}
```


## Argument Reference

* `secretkey`(Required,string)     - The Docusing secret Key from created application
* `integrationkey`(Required,string)  - The Docusing integration Key from created application
* `refreshtoken`(Required,string)      - refresh token for generation of new access token.
* `accountid`(Required,string)      - Account id of the Admin.
* `first_name`(Required,string) - First name of the User.
* `last_name`(Required,string)  - Last Name / Family Name / Surname of the User.
* `job_title`(Optional,string)         - Job title of the particular user.
* `company`(Optional,string)         - Company of the particular user.
* `permissionprofilename`(Computed,string)  - profile of the particular user ie.(DS Admin,DS Sender,DS Viewer)Can be managed only through docusign dashboard.

