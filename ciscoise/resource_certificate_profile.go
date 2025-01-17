package ciscoise

import (
	"context"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCertificateProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on CertificateProfile.

- This resource allows the client to update a certificate profile.

- This resource allows the client to create a certificate profile.
`,

		CreateContext: resourceCertificateProfileCreate,
		ReadContext:   resourceCertificateProfileRead,
		UpdateContext: resourceCertificateProfileUpdate,
		DeleteContext: resourceCertificateProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: `Unix timestamp records the last time that the resource was updated.`,
				Type:        schema.TypeString,
				Computed:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allowed_as_user_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate_attribute_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_identity_store_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"match_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"username_from": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allowed_as_user_name": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
						"certificate_attribute_name": &schema.Schema{
							Description: `Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in usernameFrom.
		Allowed values:
		- SUBJECT_COMMON_NAME
		- SUBJECT_ALTERNATIVE_NAME
		- SUBJECT_SERIAL_NUMBER
		- SUBJECT
		- SUBJECT_ALTERNATIVE_NAME_OTHER_NAME
		- SUBJECT_ALTERNATIVE_NAME_EMAIL
		- SUBJECT_ALTERNATIVE_NAME_DNS.
		- Additional internal value ALL_SUBJECT_AND_ALTERNATIVE_NAMES is used automatically when usernameFrom=UPN`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"description": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"external_identity_store_name": &schema.Schema{
							Description:      `Referred IDStore name for the Certificate Profile or [not applicable] in case no identity store is chosen`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"link": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"href": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rel": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"match_mode": &schema.Schema{
							Description: `Match mode of the Certificate Profile.
		Allowed values:
		- NEVER
		- RESOLVE_IDENTITY_AMBIGUITY
		- BINARY_COMPARISON`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"name": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"username_from": &schema.Schema{
							Description: `The attribute in the certificate where the user name should be taken from.
		Allowed values:
		- CERTIFICATE (for a specific attribute as defined in certificateAttributeName)
		- UPN (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceCertificateProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CertificateProfile create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	isEnableAutoImport := clientConfig.EnableAutoImport

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestCertificateProfileCreateCertificateProfile(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse1, _, err := client.CertificateProfile.GetCertificateProfileByID(vvID)
			if err == nil && getResponse1 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceCertificateProfileRead(ctx, d, m)
			}
		}
		if okName && vvName != "" {
			getResponse2, _, err := client.CertificateProfile.GetCertificateProfileByName(vvName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = getResponse2.CertificateProfile.ID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceCertificateProfileRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.CertificateProfile.CreateCertificateProfile(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateCertificateProfile", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateCertificateProfile", err))
		return diags
	}
	headers := restyResp1.Header()
	if locationHeader, ok := headers["Location"]; ok && len(locationHeader) > 0 {
		vvID = getLocationID(locationHeader[0])
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceCertificateProfileRead(ctx, d, m)
}

func resourceCertificateProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CertificateProfile read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetCertificateProfileByName")
		vvName := vName

		response1, restyResp1, err := client.CertificateProfile.GetCertificateProfileByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenCertificateProfileGetCertificateProfileByNameItemName(response1.CertificateProfile)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCertificateProfileByName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCertificateProfileByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCertificateProfileByID")
		vvID := vID

		response2, restyResp2, err := client.CertificateProfile.GetCertificateProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenCertificateProfileGetCertificateProfileByIDItemID(response2.CertificateProfile)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCertificateProfileByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCertificateProfileByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceCertificateProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CertificateProfile update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	var vvName string
	if selectedMethod == 1 {
		vvID = vID
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.CertificateProfile.GetCertificateProfileByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetCertificateProfileByName", err,
				"Failure at GetCertificateProfileByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.CertificateProfile != nil {
			vvID = getResp.CertificateProfile.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestCertificateProfileUpdateCertificateProfileByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.CertificateProfile.UpdateCertificateProfileByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateCertificateProfileByID", err, restyResp1.String(),
					"Failure at UpdateCertificateProfileByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateCertificateProfileByID", err,
				"Failure at UpdateCertificateProfileByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceCertificateProfileRead(ctx, d, m)
}

func resourceCertificateProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning CertificateProfile delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing CertificateProfile delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}
func expandRequestCertificateProfileCreateCertificateProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificateProfileCreateCertificateProfile {
	request := isegosdk.RequestCertificateProfileCreateCertificateProfile{}
	request.CertificateProfile = expandRequestCertificateProfileCreateCertificateProfileCertificateProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCertificateProfileCreateCertificateProfileCertificateProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificateProfileCreateCertificateProfileCertificateProfile {
	request := isegosdk.RequestCertificateProfileCreateCertificateProfileCertificateProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_identity_store_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_identity_store_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_identity_store_name")))) {
		request.ExternalIDentityStoreName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_attribute_name")))) {
		request.CertificateAttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_as_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_as_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_as_user_name")))) {
		request.AllowedAsUserName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".match_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".match_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".match_mode")))) {
		request.MatchMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username_from")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username_from")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username_from")))) {
		request.UsernameFrom = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCertificateProfileUpdateCertificateProfileByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificateProfileUpdateCertificateProfileByID {
	request := isegosdk.RequestCertificateProfileUpdateCertificateProfileByID{}
	request.CertificateProfile = expandRequestCertificateProfileUpdateCertificateProfileByIDCertificateProfile(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestCertificateProfileUpdateCertificateProfileByIDCertificateProfile(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificateProfileUpdateCertificateProfileByIDCertificateProfile {
	request := isegosdk.RequestCertificateProfileUpdateCertificateProfileByIDCertificateProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".external_identity_store_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".external_identity_store_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".external_identity_store_name")))) {
		request.ExternalIDentityStoreName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certificate_attribute_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certificate_attribute_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certificate_attribute_name")))) {
		request.CertificateAttributeName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".allowed_as_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".allowed_as_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".allowed_as_user_name")))) {
		request.AllowedAsUserName = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".match_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".match_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".match_mode")))) {
		request.MatchMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username_from")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username_from")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username_from")))) {
		request.UsernameFrom = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
