terraform {
  required_providers {
    ciscoise = {
      version = "0.1.0-rc.3"
      source  = "hashicorp.com/edu/ciscoise"
    }
  }
}

provider "ciscoise" {
}

data "ciscoise_mnt_session_profiler_count" "example" {
  provider = ciscoise
}

output "ciscoise_mnt_session_profiler_count_example" {
  value = data.ciscoise_mnt_session_profiler_count.example.item
}
