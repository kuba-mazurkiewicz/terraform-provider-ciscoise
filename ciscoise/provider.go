package ciscoise

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

// Provider definition of schema(configuration), resources(CRUD) operations and dataSources(query)
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ISE_BASE_URL", nil),
				Description: "Identity Services Engine base URL, FQDN or IP. If not set, it uses the ISE_BASE_URL environment variable.",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ISE_USERNAME", nil),
				Description: "Identity Services Engine username to authenticate. If not set, it uses the ISE_USERNAME environment variable.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ISE_PASSWORD", nil),
				Description: "Identity Services Engine password to authenticate. If not set, it uses the ISE_PASSWORD environment variable.",
			},
			"debug": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_DEBUG", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
				Description:  "Flag for Identity Services Engine to enable debugging. If not set, it uses the ISE_DEBUG environment variable; defaults to `false`.",
			},
			"ssl_verify": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_SSL_VERIFY", "true"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
				Description:  "Flag to enable or disable SSL certificate verification. If not set, it uses the ISE_SSL_VERIFY environment variable; defaults to `true`.",
			},
			"use_api_gateway": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_USE_API_GATEWAY", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
				Description:  "Flag to enable or disable the usage of the ISE's API Gateway. If not set, it uses the ISE_USE_API_GATEWAY environment variable; defaults to `false`.",
			},
			"use_csrf_token": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_USE_CSRF_TOKEN", "false"),
				ValidateFunc: validateStringHasValueFunc([]string{"true", "false"}),
				Description:  "Flag to enable or disable the usage of the X-CSRF-Token header. If not set, it uses the ISE_USE_CSRF_TOKEN environment varible; defaults to `false`.",
			},
			"single_request_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc("ISE_SINGLE_REQUEST_TIMEOUT", 60),
				ValidateFunc: validateIntegerGeqThan(0),
				Description:  "Timeout (in seconds) for the RESTful HTTP requests. If not set, it uses the ISE_SINGLE_REQUEST_TIMEOUT environment varible; defaults to 60.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ciscoise_sg_mapping_deploy":                            resourceSgMappingDeploy(),
			"ciscoise_guest_user_suspend":                           resourceGuestUserSuspend(),
			"ciscoise_guest_user_reinstate":                         resourceGuestUserReinstate(),
			"ciscoise_guest_user_deny":                              resourceGuestUserDeny(),
			"ciscoise_guest_user_approve":                           resourceGuestUserApprove(),
			"ciscoise_endpoint_certificate":                         resourceEndpointCertificate(),
			"ciscoise_endpoint_register":                            resourceEndpointRegister(),
			"ciscoise_endpoint_deregister":                          resourceEndpointDeregister(),
			"ciscoise_egress_matrix_cell_set_all_status":            resourceEgressMatrixCellSetAllStatus(),
			"ciscoise_egress_matrix_cell_clone":                     resourceEgressMatrixCellClone(),
			"ciscoise_egress_matrix_cell_clear_all":                 resourceEgressMatrixCellClearAll(),
			"ciscoise_active_directory_leave_domain_with_all_nodes": resourceActiveDirectoryLeaveDomainWithAllNodes(),
			"ciscoise_active_directory_leave_domain":                resourceActiveDirectoryLeaveDomain(),
			"ciscoise_active_directory_join_domain":                 resourceActiveDirectoryJoinDomain(),
			"ciscoise_pxgrid_service_reregister":                    resourcePxgridServiceReregister(),
			"ciscoise_pxgrid_service_register":                      resourcePxgridServiceRegister(),
			"ciscoise_pxgrid_account_create":                        resourcePxgridAccountCreate(),
			"ciscoise_threat_vulnerabilities_clear":                 resourceThreatVulnerabilitiesClear(),
			"ciscoise_sg_mapping_group_deploy":                      resourceSgMappingGroupDeploy(),
			"ciscoise_sg_mapping_deploy_all":                        resourceSgMappingDeployAll(),
			"ciscoise_px_grid_settings_auto_approve":                resourcePxGridSettingsAutoApprove(),
			"ciscoise_guest_user_reset_password":                    resourceGuestUserResetPassword(),
			"ciscoise_active_directory_join_domain_with_all_nodes":  resourceActiveDirectoryJoinDomainWithAllNodes(),
			"ciscoise_pxgrid_access_secret":                         resourcePxgridAccessSecret(),
			"ciscoise_pxgrid_service_lookup":                        resourcePxgridServiceLookup(),
			"ciscoise_pxgrid_account_activate":                      resourcePxgridAccountActivate(),
			"ciscoise_node_deployment_sync":                         resourceNodeDeploymentSync(),
			"ciscoise_selfsigned_certificate_generate":              resourceSelfsignedCertificateGenerate(),
			"ciscoise_renew_certificate":                            resourceRenewCertificate(),
			"ciscoise_ise_root_ca_regenerate":                       resourceIseRootCaRegenerate(),
			"ciscoise_bind_signed_certificate":                      resourceBindSignedCertificate(),
			"ciscoise_backup_schedule_config_update":                resourceBackupScheduleConfigUpdate(),
			"ciscoise_backup_schedule_config":                       resourceBackupScheduleConfig(),
			"ciscoise_backup_restore":                               resourceBackupRestore(),
			"ciscoise_system_certificate":                           resourceSystemCertificate(),
			"ciscoise_trusted_certificate":                          resourceTrustedCertificate(),
			"ciscoise_node_group":                                   resourceNodeGroup(),
			"ciscoise_node_deployment":                              resourceNodeDeployment(),
			"ciscoise_device_administration_conditions":             resourceDeviceAdministrationConditions(),
			"ciscoise_device_administration_network_conditions":     resourceDeviceAdministrationNetworkConditions(),
			"ciscoise_device_administration_policy_set":             resourceDeviceAdministrationPolicySet(),
			"ciscoise_device_administration_authentication_rules":   resourceDeviceAdministrationAuthenticationRules(),
			"ciscoise_device_administration_authorization_rules":    resourceDeviceAdministrationAuthorizationRules(),
			"ciscoise_device_administration_local_exception_rules":  resourceDeviceAdministrationLocalExceptionRules(),
			"ciscoise_device_administration_global_exception_rules": resourceDeviceAdministrationGlobalExceptionRules(),
			"ciscoise_device_administration_time_date_conditions":   resourceDeviceAdministrationTimeDateConditions(),
			"ciscoise_aci_settings":                                 resourceAciSettings(),
			"ciscoise_active_directory":                             resourceActiveDirectory(),
			"ciscoise_allowed_protocols":                            resourceAllowedProtocols(),
			"ciscoise_anc_endpoint":                                 resourceAncEndpoint(),
			"ciscoise_anc_policy":                                   resourceAncPolicy(),
			"ciscoise_authorization_profile":                        resourceAuthorizationProfile(),
			"ciscoise_byod_portal":                                  resourceByodPortal(),
			"ciscoise_certificate_profile":                          resourceCertificateProfile(),
			"ciscoise_downloadable_acl":                             resourceDownloadableACL(),
			"ciscoise_egress_matrix_cell":                           resourceEgressMatrixCell(),
			"ciscoise_endpoint":                                     resourceEndpoint(),
			"ciscoise_endpoint_group":                               resourceEndpointGroup(),
			"ciscoise_external_radius_server":                       resourceExternalRadiusServer(),
			"ciscoise_filter_policy":                                resourceFilterPolicy(),
			"ciscoise_guest_smtp_notification_settings":             resourceGuestSmtpNotificationSettings(),
			"ciscoise_guest_ssid":                                   resourceGuestSSID(),
			"ciscoise_guest_type":                                   resourceGuestType(),
			"ciscoise_guest_user":                                   resourceGuestUser(),
			"ciscoise_hotspot_portal":                               resourceHotspotPortal(),
			"ciscoise_identity_group":                               resourceIDentityGroup(),
			"ciscoise_id_store_sequence":                            resourceIDStoreSequence(),
			"ciscoise_internal_user":                                resourceInternalUser(),
			"ciscoise_licensing_registration":                       resourceLicensingRegistration(),
			"ciscoise_licensing_tier_state":                         resourceLicensingTierState(),
			"ciscoise_my_device_portal":                             resourceMyDevicePortal(),
			"ciscoise_network_device":                               resourceNetworkDevice(),
			"ciscoise_network_device_group":                         resourceNetworkDeviceGroup(),
			"ciscoise_native_supplicant_profile":                    resourceNativeSupplicantProfile(),
			"ciscoise_pan_ha":                                       resourcePanHa(),
			"ciscoise_portal_global_setting":                        resourcePortalGlobalSetting(),
			"ciscoise_portal_theme":                                 resourcePortalTheme(),
			"ciscoise_radius_server_sequence":                       resourceRadiusServerSequence(),
			"ciscoise_rest_id_store":                                resourceRestIDStore(),
			"ciscoise_self_registered_portal":                       resourceSelfRegisteredPortal(),
			"ciscoise_sg_acl":                                       resourceSgACL(),
			"ciscoise_sg_mapping":                                   resourceSgMapping(),
			"ciscoise_sg_mapping_group":                             resourceSgMappingGroup(),
			"ciscoise_sgt":                                          resourceSgt(),
			"ciscoise_sg_to_vn_to_vlan":                             resourceSgToVnToVLAN(),
			"ciscoise_sponsored_guest_portal":                       resourceSponsoredGuestPortal(),
			"ciscoise_sponsor_group":                                resourceSponsorGroup(),
			"ciscoise_sponsor_portal":                               resourceSponsorPortal(),
			"ciscoise_sxp_connections":                              resourceSxpConnections(),
			"ciscoise_sxp_local_bindings":                           resourceSxpLocalBindings(),
			"ciscoise_sxp_vpns":                                     resourceSxpVpns(),
			"ciscoise_tacacs_command_sets":                          resourceTacacsCommandSets(),
			"ciscoise_tacacs_external_servers":                      resourceTacacsExternalServers(),
			"ciscoise_tacacs_profile":                               resourceTacacsProfile(),
			"ciscoise_tacacs_server_sequence":                       resourceTacacsServerSequence(),
			"ciscoise_network_access_conditions":                    resourceNetworkAccessConditions(),
			"ciscoise_network_access_dictionary":                    resourceNetworkAccessDictionary(),
			"ciscoise_network_access_dictionary_attribute":          resourceNetworkAccessDictionaryAttribute(),
			"ciscoise_network_access_network_condition":             resourceNetworkAccessNetworkCondition(),
			"ciscoise_network_access_policy_set":                    resourceNetworkAccessPolicySet(),
			"ciscoise_network_access_authentication_rules":          resourceNetworkAccessAuthenticationRules(),
			"ciscoise_network_access_authorization_rules":           resourceNetworkAccessAuthorizationRules(),
			"ciscoise_network_access_local_exception_rules":         resourceNetworkAccessLocalExceptionRules(),
			"ciscoise_network_access_global_exception_rules":        resourceNetworkAccessGlobalExceptionRules(),
			"ciscoise_network_access_time_date_conditions":          resourceNetworkAccessTimeDateConditions(),
			"ciscoise_repository":                                   resourceRepository(),
			"ciscoise_trustsec_nbar_app":                            resourceTrustsecNbarApp(),
			"ciscoise_trustsec_sg_vn_mapping":                       resourceTrustsecSgVnMapping(),
			"ciscoise_trustsec_vn":                                  resourceTrustsecVn(),
			"ciscoise_trustsec_vn_vlan_mapping":                     resourceTrustsecVnVLANMapping(),
			"ciscoise_node_services_sxp_interfaces":                 resourceNodeServicesSxpInterfaces(),
			"ciscoise_node_services_profiler_probe_config":          resourceNodeServicesProfilerProbeConfig(),
			"ciscoise_proxy_connection_settings":                    resourceProxyConnectionSettings(),
			"ciscoise_px_grid_node":                                 resourcePxGridNode(),
			"ciscoise_transport_gateway_settings":                   resourceTransportGatewaySettings(),
			"ciscoise_node_group_node":                              resourceNodeGroupNode(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ciscoise_mnt_account_status":                                         dataSourceMntAccountStatus(),
			"ciscoise_mnt_authentication_status":                                  dataSourceMntAuthenticationStatus(),
			"ciscoise_backup_last_status":                                         dataSourceBackupLastStatus(),
			"ciscoise_csr":                                                        dataSourceCsr(),
			"ciscoise_csr_export":                                                 dataSourceCsrExport(),
			"ciscoise_system_certificate":                                         dataSourceSystemCertificate(),
			"ciscoise_system_certificate_export_info":                             dataSourceSystemCertificateExportInfo(),
			"ciscoise_trusted_certificate":                                        dataSourceTrustedCertificate(),
			"ciscoise_trusted_certificate_export":                                 dataSourceTrustedCertificateExport(),
			"ciscoise_mnt_session_disconnect":                                     dataSourceMntSessionDisconnect(),
			"ciscoise_mnt_session_reauthentication":                               dataSourceMntSessionReauthentication(),
			"ciscoise_node_group":                                                 dataSourceNodeGroup(),
			"ciscoise_node_deployment":                                            dataSourceNodeDeployment(),
			"ciscoise_pan_ha":                                                     dataSourcePanHa(),
			"ciscoise_device_administration_command_set":                          dataSourceDeviceAdministrationCommandSet(),
			"ciscoise_device_administration_conditions":                           dataSourceDeviceAdministrationConditions(),
			"ciscoise_device_administration_conditions_for_authentication_rule":   dataSourceDeviceAdministrationConditionsForAuthenticationRule(),
			"ciscoise_device_administration_conditions_for_authorization_rule":    dataSourceDeviceAdministrationConditionsForAuthorizationRule(),
			"ciscoise_device_administration_conditions_for_policy_set":            dataSourceDeviceAdministrationConditionsForPolicySet(),
			"ciscoise_device_administration_dictionary_attributes_authentication": dataSourceDeviceAdministrationDictionaryAttributesAuthentication(),
			"ciscoise_device_administration_dictionary_attributes_authorization":  dataSourceDeviceAdministrationDictionaryAttributesAuthorization(),
			"ciscoise_device_administration_dictionary_attributes_policy_set":     dataSourceDeviceAdministrationDictionaryAttributesPolicySet(),
			"ciscoise_device_administration_identity_stores":                      dataSourceDeviceAdministrationIDentityStores(),
			"ciscoise_device_administration_network_conditions":                   dataSourceDeviceAdministrationNetworkConditions(),
			"ciscoise_device_administration_policy_set":                           dataSourceDeviceAdministrationPolicySet(),
			"ciscoise_device_administration_authentication_rules":                 dataSourceDeviceAdministrationAuthenticationRules(),
			"ciscoise_device_administration_authorization_rules":                  dataSourceDeviceAdministrationAuthorizationRules(),
			"ciscoise_device_administration_local_exception_rules":                dataSourceDeviceAdministrationLocalExceptionRules(),
			"ciscoise_device_administration_global_exception_rules":               dataSourceDeviceAdministrationGlobalExceptionRules(),
			"ciscoise_device_administration_profiles":                             dataSourceDeviceAdministrationProfiles(),
			"ciscoise_device_administration_service_names":                        dataSourceDeviceAdministrationServiceNames(),
			"ciscoise_device_administration_time_date_conditions":                 dataSourceDeviceAdministrationTimeDateConditions(),
			"ciscoise_aci_bindings":                                               dataSourceAciBindings(),
			"ciscoise_aci_settings":                                               dataSourceAciSettings(),
			"ciscoise_aci_test_connectivity":                                      dataSourceAciTestConnectivity(),
			"ciscoise_active_directory":                                           dataSourceActiveDirectory(),
			"ciscoise_active_directory_get_groups_by_domain_info":                 dataSourceActiveDirectoryGetGroupsByDomainInfo(),
			"ciscoise_active_directory_get_trusted_domains_info":                  dataSourceActiveDirectoryGetTrustedDomainsInfo(),
			"ciscoise_active_directory_get_user_groups_info":                      dataSourceActiveDirectoryGetUserGroupsInfo(),
			"ciscoise_active_directory_is_user_member_of_group":                   dataSourceActiveDirectoryIsUserMemberOfGroup(),
			"ciscoise_admin_user":                                                 dataSourceAdminUser(),
			"ciscoise_allowed_protocols":                                          dataSourceAllowedProtocols(),
			"ciscoise_anc_endpoint":                                               dataSourceAncEndpoint(),
			"ciscoise_anc_endpoint_bulk_monitor_status":                           dataSourceAncEndpointBulkMonitorStatus(),
			"ciscoise_anc_policy":                                                 dataSourceAncPolicy(),
			"ciscoise_anc_policy_bulk_monitor_status":                             dataSourceAncPolicyBulkMonitorStatus(),
			"ciscoise_authorization_profile":                                      dataSourceAuthorizationProfile(),
			"ciscoise_byod_portal":                                                dataSourceByodPortal(),
			"ciscoise_certificate_profile":                                        dataSourceCertificateProfile(),
			"ciscoise_certificate_template":                                       dataSourceCertificateTemplate(),
			"ciscoise_deployment":                                                 dataSourceDeployment(),
			"ciscoise_downloadable_acl":                                           dataSourceDownloadableACL(),
			"ciscoise_egress_matrix_cell":                                         dataSourceEgressMatrixCell(),
			"ciscoise_egress_matrix_cell_bulk_monitor_status":                     dataSourceEgressMatrixCellBulkMonitorStatus(),
			"ciscoise_endpoint":                                                   dataSourceEndpoint(),
			"ciscoise_endpoint_bulk_monitor_status":                               dataSourceEndpointBulkMonitorStatus(),
			"ciscoise_endpoint_get_rejected_endpoints":                            dataSourceEndpointGetRejectedEndpoints(),
			"ciscoise_endpoint_group":                                             dataSourceEndpointGroup(),
			"ciscoise_external_radius_server":                                     dataSourceExternalRadiusServer(),
			"ciscoise_filter_policy":                                              dataSourceFilterPolicy(),
			"ciscoise_guest_location":                                             dataSourceGuestLocation(),
			"ciscoise_guest_smtp_notification_settings":                           dataSourceGuestSmtpNotificationSettings(),
			"ciscoise_guest_ssid":                                                 dataSourceGuestSSID(),
			"ciscoise_guest_type":                                                 dataSourceGuestType(),
			"ciscoise_guest_user":                                                 dataSourceGuestUser(),
			"ciscoise_guest_user_bulk_monitor_status":                             dataSourceGuestUserBulkMonitorStatus(),
			"ciscoise_hotspot_portal":                                             dataSourceHotspotPortal(),
			"ciscoise_identity_group":                                             dataSourceIDentityGroup(),
			"ciscoise_id_store_sequence":                                          dataSourceIDStoreSequence(),
			"ciscoise_internal_user":                                              dataSourceInternalUser(),
			"ciscoise_my_device_portal":                                           dataSourceMyDevicePortal(),
			"ciscoise_network_device":                                             dataSourceNetworkDevice(),
			"ciscoise_network_device_bulk_monitor_status":                         dataSourceNetworkDeviceBulkMonitorStatus(),
			"ciscoise_network_device_group":                                       dataSourceNetworkDeviceGroup(),
			"ciscoise_node":                                                       dataSourceNode(),
			"ciscoise_native_supplicant_profile":                                  dataSourceNativeSupplicantProfile(),
			"ciscoise_system_config_version":                                      dataSourceSystemConfigVersion(),
			"ciscoise_portal":                                                     dataSourcePortal(),
			"ciscoise_portal_global_setting":                                      dataSourcePortalGlobalSetting(),
			"ciscoise_portal_theme":                                               dataSourcePortalTheme(),
			"ciscoise_profiler_profile":                                           dataSourceProfilerProfile(),
			"ciscoise_px_grid_node":                                               dataSourcePxGridNode(),
			"ciscoise_radius_server_sequence":                                     dataSourceRadiusServerSequence(),
			"ciscoise_rest_id_store":                                              dataSourceRestIDStore(),
			"ciscoise_self_registered_portal":                                     dataSourceSelfRegisteredPortal(),
			"ciscoise_session_service_node":                                       dataSourceSessionServiceNode(),
			"ciscoise_sg_acl":                                                     dataSourceSgACL(),
			"ciscoise_sg_acl_bulk_monitor_status":                                 dataSourceSgACLBulkMonitorStatus(),
			"ciscoise_sg_mapping":                                                 dataSourceSgMapping(),
			"ciscoise_sg_mapping_bulk_monitor_status":                             dataSourceSgMappingBulkMonitorStatus(),
			"ciscoise_sg_mapping_deploy_status_info":                              dataSourceSgMappingDeployStatusInfo(),
			"ciscoise_sg_mapping_group":                                           dataSourceSgMappingGroup(),
			"ciscoise_sg_mapping_group_bulk_monitor_status":                       dataSourceSgMappingGroupBulkMonitorStatus(),
			"ciscoise_sg_mapping_group_deploy_status_info":                        dataSourceSgMappingGroupDeployStatusInfo(),
			"ciscoise_sgt":                                                        dataSourceSgt(),
			"ciscoise_sgt_bulk_monitor_status":                                    dataSourceSgtBulkMonitorStatus(),
			"ciscoise_sg_to_vn_to_vlan":                                           dataSourceSgToVnToVLAN(),
			"ciscoise_sg_to_vn_to_vlan_bulk_monitor_status":                       dataSourceSgToVnToVLANBulkMonitorStatus(),
			"ciscoise_sms_provider":                                               dataSourceSmsProvider(),
			"ciscoise_sponsored_guest_portal":                                     dataSourceSponsoredGuestPortal(),
			"ciscoise_sponsor_group":                                              dataSourceSponsorGroup(),
			"ciscoise_sponsor_group_member":                                       dataSourceSponsorGroupMember(),
			"ciscoise_sponsor_portal":                                             dataSourceSponsorPortal(),
			"ciscoise_support_bundle_download":                                    dataSourceSupportBundleDownload(),
			"ciscoise_support_bundle_status":                                      dataSourceSupportBundleStatus(),
			"ciscoise_sxp_connections":                                            dataSourceSxpConnections(),
			"ciscoise_sxp_connections_bulk_monitor_status":                        dataSourceSxpConnectionsBulkMonitorStatus(),
			"ciscoise_sxp_local_bindings":                                         dataSourceSxpLocalBindings(),
			"ciscoise_sxp_local_bindings_bulk_monitor_status":                     dataSourceSxpLocalBindingsBulkMonitorStatus(),
			"ciscoise_sxp_vpns":                                                   dataSourceSxpVpns(),
			"ciscoise_sxp_vpns_bulk_monitor_status":                               dataSourceSxpVpnsBulkMonitorStatus(),
			"ciscoise_tacacs_command_sets":                                        dataSourceTacacsCommandSets(),
			"ciscoise_tacacs_external_servers":                                    dataSourceTacacsExternalServers(),
			"ciscoise_tacacs_profile":                                             dataSourceTacacsProfile(),
			"ciscoise_tacacs_server_sequence":                                     dataSourceTacacsServerSequence(),
			"ciscoise_telemetry_info":                                             dataSourceTelemetryInfo(),
			"ciscoise_mnt_failure_reasons":                                        dataSourceMntFailureReasons(),
			"ciscoise_pxgrid_failures":                                            dataSourcePxgridFailures(),
			"ciscoise_pxgrid_profiles_info":                                       dataSourcePxgridProfilesInfo(),
			"ciscoise_pxgrid_egress_matrices_info":                                dataSourcePxgridEgressMatricesInfo(),
			"ciscoise_pxgrid_egress_policies_info":                                dataSourcePxgridEgressPoliciesInfo(),
			"ciscoise_pxgrid_security_group_acls_info":                            dataSourcePxgridSecurityGroupACLsInfo(),
			"ciscoise_pxgrid_security_groups_info":                                dataSourcePxgridSecurityGroupsInfo(),
			"ciscoise_pxgrid_endpoint_by_mac_info":                                dataSourcePxgridEndpointByMacInfo(),
			"ciscoise_pxgrid_endpoints_info":                                      dataSourcePxgridEndpointsInfo(),
			"ciscoise_pxgrid_endpoints_by_os_type_info":                           dataSourcePxgridEndpointsByOsTypeInfo(),
			"ciscoise_pxgrid_endpoints_by_type_info":                              dataSourcePxgridEndpointsByTypeInfo(),
			"ciscoise_pxgrid_session_by_ip_info":                                  dataSourcePxgridSessionByIPInfo(),
			"ciscoise_pxgrid_session_by_mac_info":                                 dataSourcePxgridSessionByMacInfo(),
			"ciscoise_pxgrid_sessions_info":                                       dataSourcePxgridSessionsInfo(),
			"ciscoise_pxgrid_session_for_recovery_info":                           dataSourcePxgridSessionForRecoveryInfo(),
			"ciscoise_pxgrid_user_group_by_username_info":                         dataSourcePxgridUserGroupByUsernameInfo(),
			"ciscoise_pxgrid_user_groups_info":                                    dataSourcePxgridUserGroupsInfo(),
			"ciscoise_pxgrid_bindings_info":                                       dataSourcePxgridBindingsInfo(),
			"ciscoise_pxgrid_healths_info":                                        dataSourcePxgridHealthsInfo(),
			"ciscoise_pxgrid_performances_info":                                   dataSourcePxgridPerformancesInfo(),
			"ciscoise_network_access_conditions":                                  dataSourceNetworkAccessConditions(),
			"ciscoise_network_access_conditions_for_authentication_rule":          dataSourceNetworkAccessConditionsForAuthenticationRule(),
			"ciscoise_network_access_conditions_for_authorization_rule":           dataSourceNetworkAccessConditionsForAuthorizationRule(),
			"ciscoise_network_access_conditions_for_policy_set":                   dataSourceNetworkAccessConditionsForPolicySet(),
			"ciscoise_network_access_dictionary":                                  dataSourceNetworkAccessDictionary(),
			"ciscoise_network_access_dictionary_attribute":                        dataSourceNetworkAccessDictionaryAttribute(),
			"ciscoise_network_access_dictionary_attributes_authentication":        dataSourceNetworkAccessDictionaryAttributesAuthentication(),
			"ciscoise_network_access_dictionary_attributes_authorization":         dataSourceNetworkAccessDictionaryAttributesAuthorization(),
			"ciscoise_network_access_dictionary_attributes_policy_set":            dataSourceNetworkAccessDictionaryAttributesPolicySet(),
			"ciscoise_network_access_identity_stores":                             dataSourceNetworkAccessIDentityStores(),
			"ciscoise_network_access_network_condition":                           dataSourceNetworkAccessNetworkCondition(),
			"ciscoise_network_access_policy_set":                                  dataSourceNetworkAccessPolicySet(),
			"ciscoise_network_access_authentication_rules":                        dataSourceNetworkAccessAuthenticationRules(),
			"ciscoise_network_access_authorization_rules":                         dataSourceNetworkAccessAuthorizationRules(),
			"ciscoise_network_access_local_exception_rules":                       dataSourceNetworkAccessLocalExceptionRules(),
			"ciscoise_network_access_global_exception_rules":                      dataSourceNetworkAccessGlobalExceptionRules(),
			"ciscoise_network_access_profiles":                                    dataSourceNetworkAccessProfiles(),
			"ciscoise_network_access_security_groups":                             dataSourceNetworkAccessSecurityGroups(),
			"ciscoise_network_access_service_name":                                dataSourceNetworkAccessServiceName(),
			"ciscoise_network_access_time_date_conditions":                        dataSourceNetworkAccessTimeDateConditions(),
			"ciscoise_repository":                                                 dataSourceRepository(),
			"ciscoise_repository_files":                                           dataSourceRepositoryFiles(),
			"ciscoise_mnt_sessions_by_session_id":                                 dataSourceMntSessionsBySessionID(),
			"ciscoise_mnt_session_active_count":                                   dataSourceMntSessionActiveCount(),
			"ciscoise_mnt_session_active_list":                                    dataSourceMntSessionActiveList(),
			"ciscoise_mnt_session_auth_list":                                      dataSourceMntSessionAuthList(),
			"ciscoise_mnt_session_by_ip":                                          dataSourceMntSessionByIP(),
			"ciscoise_mnt_session_by_nas_ip":                                      dataSourceMntSessionByNasIP(),
			"ciscoise_mnt_session_by_mac":                                         dataSourceMntSessionByMac(),
			"ciscoise_mnt_session_posture_count":                                  dataSourceMntSessionPostureCount(),
			"ciscoise_mnt_session_profiler_count":                                 dataSourceMntSessionProfilerCount(),
			"ciscoise_mnt_session_by_username":                                    dataSourceMntSessionByUsername(),
			"ciscoise_tasks":                                                      dataSourceTasks(),
			"ciscoise_mnt_version":                                                dataSourceMntVersion(),
			"ciscoise_resource_version":                                           dataSourceResourceVersion(),
			"ciscoise_trustsec_nbar_app":                                          dataSourceTrustsecNbarApp(),
			"ciscoise_trustsec_sg_vn_mapping":                                     dataSourceTrustsecSgVnMapping(),
			"ciscoise_trustsec_vn":                                                dataSourceTrustsecVn(),
			"ciscoise_trustsec_vn_vlan_mapping":                                   dataSourceTrustsecVnVLANMapping(),
			"ciscoise_node_services_interfaces":                                   dataSourceNodeServicesInterfaces(),
			"ciscoise_node_services_sxp_interfaces":                               dataSourceNodeServicesSxpInterfaces(),
			"ciscoise_node_services_profiler_probe_config":                        dataSourceNodeServicesProfilerProbeConfig(),
			"ciscoise_licensing_connection_type":                                  dataSourceLicensingConnectionType(),
			"ciscoise_licensing_eval_license":                                     dataSourceLicensingEvalLicense(),
			"ciscoise_licensing_feature_to_tier_mapping":                          dataSourceLicensingFeatureToTierMapping(),
			"ciscoise_licensing_registration":                                     dataSourceLicensingRegistration(),
			"ciscoise_licensing_smart_state":                                      dataSourceLicensingSmartState(),
			"ciscoise_licensing_tier_state":                                       dataSourceLicensingTierState(),
			"ciscoise_hotpatch":                                                   dataSourceHotpatch(),
			"ciscoise_patch":                                                      dataSourcePatch(),
			"ciscoise_proxy_connection_settings":                                  dataSourceProxyConnectionSettings(),
			"ciscoise_transport_gateway_settings":                                 dataSourceTransportGatewaySettings(),
			"ciscoise_node_group_node":                                            dataSourceNodeGroupNode(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
