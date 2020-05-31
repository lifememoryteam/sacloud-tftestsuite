resource "sakuracloud_server" "validserver02" {
  name   = "validserver01"
  core   = 100
  memory = 100
}

resource "sakuracloud_server" "unknownserver02" {
  name   = "unknownserver01"
  core   = 1
  memory = 2
}