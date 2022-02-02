---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_licensing_eval_license Data Source - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs read operation on Licensing.
  Get registration information
---

# ciscoise_licensing_eval_license (Data Source)

It performs read operation on Licensing.

- Get registration information

## Example Usage

```terraform
data "ciscoise_licensing_eval_license" "example" {
  provider = ciscoise
}

output "ciscoise_licensing_eval_license_example" {
  value = data.ciscoise_licensing_eval_license.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **item** (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **days_remaining** (Number)

