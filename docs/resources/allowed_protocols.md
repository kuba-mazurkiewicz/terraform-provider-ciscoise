---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ciscoise_allowed_protocols Resource - terraform-provider-ciscoise"
subcategory: ""
description: |-
  
---

# ciscoise_allowed_protocols (Resource)





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

- **allow_chap** (Boolean)
- **allow_eap_fast** (Boolean)
- **allow_eap_md5** (Boolean)
- **allow_eap_tls** (Boolean)
- **allow_eap_ttls** (Boolean)
- **allow_leap** (Boolean)
- **allow_ms_chap_v1** (Boolean)
- **allow_ms_chap_v2** (Boolean)
- **allow_pap_ascii** (Boolean)
- **allow_peap** (Boolean)
- **allow_preferred_eap_protocol** (Boolean)
- **allow_teap** (Boolean)
- **allow_weak_ciphers_for_eap** (Boolean)
- **description** (String)
- **eap_fast** (Block List) (see [below for nested schema](#nestedblock--item--eap_fast))
- **eap_tls** (Block List) (see [below for nested schema](#nestedblock--item--eap_tls))
- **eap_tls_l_bit** (Boolean)
- **eap_ttls** (Block List) (see [below for nested schema](#nestedblock--item--eap_ttls))
- **id** (String) The ID of this resource.
- **name** (String)
- **peap** (Block List) (see [below for nested schema](#nestedblock--item--peap))
- **preferred_eap_protocol** (String)
- **process_host_lookup** (Boolean)
- **require_message_auth** (Boolean)
- **teap** (Block List) (see [below for nested schema](#nestedblock--item--teap))

Read-Only:

- **link** (List of Object) (see [below for nested schema](#nestedatt--item--link))

<a id="nestedblock--item--eap_fast"></a>
### Nested Schema for `item.eap_fast`

Optional:

- **allow_eap_fast_eap_gtc** (Boolean)
- **allow_eap_fast_eap_gtc_pwd_change** (Boolean)
- **allow_eap_fast_eap_gtc_pwd_change_retries** (Number)
- **allow_eap_fast_eap_ms_chap_v2** (Boolean)
- **allow_eap_fast_eap_ms_chap_v2_pwd_change** (Boolean)
- **allow_eap_fast_eap_ms_chap_v2_pwd_change_retries** (Number)
- **allow_eap_fast_eap_tls** (Boolean)
- **allow_eap_fast_eap_tls_auth_of_expired_certs** (Boolean)
- **eap_fast_dont_use_pacs_accept_client_cert** (Boolean)
- **eap_fast_dont_use_pacs_allow_machine_authentication** (Boolean)
- **eap_fast_enable_eap_chaining** (Boolean)
- **eap_fast_use_pacs** (Boolean)
- **eap_fast_use_pacs_accept_client_cert** (Boolean)
- **eap_fast_use_pacs_allow_anonym_provisioning** (Boolean)
- **eap_fast_use_pacs_allow_authen_provisioning** (Boolean)
- **eap_fast_use_pacs_allow_machine_authentication** (Boolean)
- **eap_fast_use_pacs_authorization_pac_ttl** (Number)
- **eap_fast_use_pacs_authorization_pac_ttl_units** (String)
- **eap_fast_use_pacs_machine_pac_ttl** (Number)
- **eap_fast_use_pacs_machine_pac_ttl_units** (String)
- **eap_fast_use_pacs_return_access_accept_after_authenticated_provisioning** (Boolean)
- **eap_fast_use_pacs_stateless_session_resume** (Boolean)
- **eap_fast_use_pacs_tunnel_pac_ttl** (Number)
- **eap_fast_use_pacs_tunnel_pac_ttl_units** (String)
- **eap_fast_use_pacs_use_proactive_pac_update_precentage** (Number)


<a id="nestedblock--item--eap_tls"></a>
### Nested Schema for `item.eap_tls`

Optional:

- **allow_eap_tls_auth_of_expired_certs** (Boolean)
- **eap_tls_enable_stateless_session_resume** (Boolean)
- **eap_tls_session_ticket_precentage** (Number)
- **eap_tls_session_ticket_ttl** (Number)
- **eap_tls_session_ticket_ttl_units** (String)


<a id="nestedblock--item--eap_ttls"></a>
### Nested Schema for `item.eap_ttls`

Optional:

- **eap_ttls_chap** (Boolean)
- **eap_ttls_eap_md5** (Boolean)
- **eap_ttls_eap_ms_chap_v2** (Boolean)
- **eap_ttls_eap_ms_chap_v2_pwd_change** (Boolean)
- **eap_ttls_eap_ms_chap_v2_pwd_change_retries** (Number)
- **eap_ttls_ms_chap_v1** (Boolean)
- **eap_ttls_ms_chap_v2** (Boolean)
- **eap_ttls_pap_ascii** (Boolean)


<a id="nestedblock--item--peap"></a>
### Nested Schema for `item.peap`

Optional:

- **allow_peap_eap_gtc** (Boolean)
- **allow_peap_eap_gtc_pwd_change** (Boolean)
- **allow_peap_eap_gtc_pwd_change_retries** (Number)
- **allow_peap_eap_ms_chap_v2** (Boolean)
- **allow_peap_eap_ms_chap_v2_pwd_change** (Boolean)
- **allow_peap_eap_ms_chap_v2_pwd_change_retries** (Number)
- **allow_peap_eap_tls** (Boolean)
- **allow_peap_eap_tls_auth_of_expired_certs** (Boolean)
- **allow_peap_v0** (Boolean)
- **require_cryptobinding** (Boolean)


<a id="nestedblock--item--teap"></a>
### Nested Schema for `item.teap`

Optional:

- **accept_client_cert_during_tunnel_est** (Boolean)
- **allow_downgrade_msk** (Boolean)
- **allow_teap_eap_ms_chap_v2** (Boolean)
- **allow_teap_eap_ms_chap_v2_pwd_change** (Boolean)
- **allow_teap_eap_ms_chap_v2_pwd_change_retries** (Number)
- **allow_teap_eap_tls** (Boolean)
- **allow_teap_eap_tls_auth_of_expired_certs** (Boolean)
- **enable_eap_chaining** (Boolean)


<a id="nestedatt--item--link"></a>
### Nested Schema for `item.link`

Read-Only:

- **href** (String)
- **rel** (String)
- **type** (String)

