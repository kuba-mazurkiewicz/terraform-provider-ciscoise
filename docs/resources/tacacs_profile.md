---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_tacacs_profile Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_tacacs_profile (Resource)





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

- **description** (String)
- **id** (String) The ID of this resource.
- **name** (String)
- **session_attributes** (Block List) (see [below for nested schema](#nestedblock--item--session_attributes))

Read-Only:

- **link** (List of Object) (see [below for nested schema](#nestedatt--item--link))

<a id="nestedblock--item--session_attributes"></a>
### Nested Schema for `item.session_attributes`

Optional:

- **session_attribute_list** (Block List) (see [below for nested schema](#nestedblock--item--session_attributes--session_attribute_list))

<a id="nestedblock--item--session_attributes--session_attribute_list"></a>
### Nested Schema for `item.session_attributes.session_attribute_list`

Optional:

- **name** (String)
- **type** (String)
- **value** (String)



<a id="nestedatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

