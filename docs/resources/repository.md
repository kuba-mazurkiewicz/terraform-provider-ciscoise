---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_repository Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_repository (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **item** (Block List) (see [below for nested schema](#nestedblock--item))

### Read-Only

- **last_updated** (String)

<a id="nestedblock--item"></a>
### Nested Schema for `item`

Optional:

- **enable_pki** (Boolean)
- **name** (String)
- **password** (String, Sensitive)
- **path** (String)
- **protocol** (String)
- **server_name** (String)
- **user_name** (String)

