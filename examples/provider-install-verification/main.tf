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
  name             = "1_TEST"
  description      = "Test Description"
  parent_folder_id = "956755"
}

resource "td_folder" "child" {
  name             = "0_CHILD_UPDATE"
  description      = "Child Description"
  parent_folder_id = td_folder.test.id
}

output "all_parent_segments" {
  value = data.td_parent_segments.all
}
