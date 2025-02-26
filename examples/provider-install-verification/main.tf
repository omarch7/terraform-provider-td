terraform {
  required_providers {
    td = {
      source = "hellofresh.io/mcp/td"
    }
  }
}

provider "td" {}

data "td_parent_segments" "all" {}

resource "td_folder" "test" {
  name             = "1_TEST_UPDATE"
  description      = "Test Description" 
  parent_folder_id = "956755"
}

output "all_parent_segments" {
  value = data.td_parent_segments.all
}
