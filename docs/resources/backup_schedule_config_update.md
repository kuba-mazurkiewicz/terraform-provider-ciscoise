---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_backup_schedule_config_update Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  It performs update operation on Backup And Restore.
  - Update the Schedule of the configuration backup on the ISE node as per the input parameters. This resource
  only helps in editing the schedule.
---

# ciscoise_backup_schedule_config_update (Resource)

It performs update operation on Backup And Restore.
- Update the Schedule of the configuration backup on the ISE node as per the input parameters. This resource
only helps in editing the schedule.

~>Warning: This resource does not represent a real-world entity in Cisco ISE, therefore changing or deleting this resource on its own has no immediate effect. Instead, it is a task part of a Cisco ISE workflow. It is executed in ISE without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "ciscoise_backup_schedule_config_update" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    backup_description    = "string"
    backup_encryption_key = "string"
    backup_name           = "string"
    end_date              = "string"
    frequency             = "string"
    month_day             = "string"
    repository_name       = "string"
    start_date            = "string"
    status                = "string"
    time                  = "string"
    week_day              = "string"
  }
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

- **backup_description** (String) Description of the backup.
- **backup_encryption_key** (String) The encyption key for the backed up file. Encryption key must satisfy the following criteria - Contains at least one uppercase letter [A-Z], Contains at least one lowercase letter [a-z], Contains at least one digit [0-9], Contain only [A-Z][a-z][0-9]_#, Has at least 8 characters, Has not more than 15 characters, Must not contain 'CcIiSsCco', Must not begin with
- **backup_name** (String) The backup file will get saved with this name.
- **end_date** (String) End date of the scheduled backup job. Allowed format MM/DD/YYYY. End date is not required in case of ONE_TIME frequency.
- **frequency** (String)
- **month_day** (String) Day of month you want backup to be performed on when scheduled frequency is MONTHLY. Allowed values - from 1 to 28.
- **repository_name** (String) Name of the configured repository where the generated backup file will get copied.
- **start_date** (String) Start date for scheduling the backup job. Allowed format MM/DD/YYYY.
- **status** (String)
- **time** (String) Time at which backup job get scheduled. example- 12:00 AM
- **week_day** (String)


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- **link** (List of Object) (see [below for nested schema](#nestedobjatt--item--link))
- **message** (String)

<a id="nestedobjatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

