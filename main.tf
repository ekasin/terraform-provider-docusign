terraform {
  required_providers {
    docusign = {
      version = "0.2"
      source  = "terraform/provider/docusign"
    }
  }
}

provider "docusign" {
  secretkey = "2eb389be-ef04-4f5e-aea2-b7d6b9efb1c2"
  integrationkey = "a8adb361-9ead-4069-ba34-102bb972ab17"
  refreshtoken   = "eyJ0eXAiOiJNVCIsImFsZyI6IlJTMjU2Iiwia2lkIjoiNjgxODVmZjEtNGU1MS00Y2U5LWFmMWMtNjg5ODEyMjAzMzE3In0.AQoAAAABAAgABwAAxVKv9iPZSAgAAEW3p4k72UgCANEePMOL2c1LjayKwZaWOZMVAAEAAAAYAAEAAAAFAAAADQAkAAAAYThhZGIzNjEtOWVhZC00MDY5LWJhMzQtMTAyYmI5NzJhYjE3IgAkAAAAYThhZGIzNjEtOWVhZC00MDY5LWJhMzQtMTAyYmI5NzJhYjE3MAAAFlepzhDZSDcAoSgUJQ2MeU2Q6Y5or-gk4A.xJCb6NrTgx7hNKKXPOPTFvGvdjpG0D4p-_H6yYV0_TcngccL5t8l1Dl5UQo8_ZmZL_cNoaGdAF6JJLx9QjpxPgBGwMBHDKhGbem4mjob5xKN-7t5bgSyDJpdFXLHfJOrTklxNkr995nbQX6a2C-vapIEDWFH7lbKpwtwzLBp1v7PfpJ1hN9kcK_dXwISWEFPTbo4j1ul2dCJlESS0IQbTAQZj9PB7LLOzzzRugMbcnbph3jwH1-Gknh1i-3VLvXAFlEun2TQ2PCPWOATdAP0hISwuevPkFORtvYQPgbtaQ_qRf4dcb3ve947JAfPOq5JQi-E8BUUjjfZiGTNvezEew"
  accountid = "b83a2769-7cd6-43c7-868e-21b58a562ce6"
}


/*
resource "docusign_user" "user1" {
  email = "ui17co14@iiitsurat.ac.in"
  firstname = "ekansh"
  lastname = "singh"
  jobtitle = "manager"
  company = "clevertap"
}
*/



data "docusign_user" "user3" {
  email = "cpp7coder@gmail.com"
}

output "user3" {
  value = data.docusign_user.user3
}

