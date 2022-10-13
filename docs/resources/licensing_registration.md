---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_licensing_registration Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It manages DEREGISTER, REGISTER, RENEW and UPDATE operations on License - registration information.
  License Configure registration information.
---

# ciscoise_licensing_registration (Resource)

It manages DEREGISTER, REGISTER, RENEW and UPDATE operations on License - registration information.

- License Configure registration information.

## Example Usage

```terraform
resource "ciscoise_licensing_registration" "example" {
  provider = ciscoise
  parameters {
    connection_type    = "string" # "HTTP_DIRECT", "PROXY", "SSM_ONPREM_SERVER", "TRANSPORT_GATEWAY"
    registration_type  = "string" # "DEREGISTER", "REGISTER", "RENEW", "UPDATE"
    ssm_on_prem_server = "string"
    tier               = ["string"] # "ADVANTAGE", "DEVICEADMIN", "ESSENTIAL", "PREMIER", "VM"
    token              = "string"
  }
}

output "ciscoise_licensing_registration_example" {
  value = ciscoise_licensing_registration.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parameters** (Block List, Min: 1, Max: 1) it is a container o ISE API parameters. (see [below for nested schema](#nestedblock--parameters))

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))
- **last_updated** (String) Unix timestamp records the last time that the resource was updated.

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- **connection_type** (String)
- **registration_type** (String)
- **ssm_on_prem_server** (String) If connection type is selected as SSM_ONPREM_SERVER, then  IP address or the hostname (or FQDN) of the SSM On-Prem server Host.
- **tier** (List of String)
- **token** (String) token


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **connection_type** (String)
- **registration_state** (String)
- **ssm_on_prem_server** (String)
- **tier** (List of String)

## Import

Import is supported using the following syntax:

```shell
terraform import ciscoise_licensing_registration.example "connection_type:=string\registration_type:=string\ssm_on_prem_server:=string\tier:=string"
```