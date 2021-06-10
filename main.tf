terraform {
  required_providers {
    docusign = {
      version = "0.2"
      source  = "terraform/provider/docusign"
    }
  }
}

provider "docusign" {
  secretkey = ""
  integrationkey = ""
  refreshtoken   = ""
  accountid = ""
}

resource "docusign_user" "user2" {
  email = ""
  firstname = ""
  lastname = ""
  jobtitle = ""
  company = ""
}

data "docusign_user" "user2" {
  email = ""
}

output "user3" {
  value = data.docusign_user.user3
}
