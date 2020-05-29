resource "sakuracloud_server" "validserver01" {
  name   = "validserver01"
  core   = 2
  memory = 4
}

resource "sakuracloud_server" "unknownserver01" {
  name   = "unknownserver01"
  core   = 100
  memory = 4
}