terraform {
  required_providers {
    td = {
      source  = "hellofresh.io/mcp/td"
    }
  }
}

provider "td" {}

data "td_parent_segments" "all" {}

output "all_parent_segments" {
  value = data.td_parent_segments.all
}
