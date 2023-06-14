terraform {
  required_providers {
    ciscoise = {
      version = "0.6.19-beta"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

resource "ciscoise_pxgrid_account_activate" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    description = "string"
  }
}