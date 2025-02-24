terraform {
  required_providers {
    td = {
      source  = "hellofresh.io/mcp/td"
    }
  }
}

provider "td" {}

data "td_folders" "example" {}
