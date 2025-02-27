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
  for_each = {
    for parent_segment in data.td_parent_segments.all.parent_segments : parent_segment.id => parent_segment
  }

  name             = "1_TEST_TF"
  description      = "Test Description"
  parent_folder_id = each.value.parent_folder_id
}

resource "td_folder" "child" {
  for_each = {
    for parent_segment in data.td_parent_segments.all.parent_segments : parent_segment.id => parent_segment
  }
  name             = "1_TEST_TF_CHILD"
  description      = "Test Description"
  parent_folder_id = td_folder.test[each.key].id
}

output "all_parent_segments" {
  value = data.td_parent_segments.all
}
