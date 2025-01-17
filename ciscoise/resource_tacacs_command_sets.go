package ciscoise

import (
	"context"
	"fmt"
	"reflect"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTacacsCommandSets() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on TACACSCommandSets.

- This resource allows the client to update TACACS command sets.

- This resource deletes TACACS command sets.

- This resource creates TACACS command sets.
`,

		CreateContext: resourceTacacsCommandSetsCreate,
		ReadContext:   resourceTacacsCommandSetsRead,
		UpdateContext: resourceTacacsCommandSetsUpdate,
		DeleteContext: resourceTacacsCommandSetsDelete,
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

						"commands": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"command_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"arguments": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"command": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
												"grant": &schema.Schema{
													Description: `Allowed values: PERMIT, DENY, DENY_ALWAYS`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"description": &schema.Schema{
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
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"permit_unmatched": &schema.Schema{
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

						"commands": &schema.Schema{
							Type:             schema.TypeList,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"command_list": &schema.Schema{
										Type:             schema.TypeList,
										Optional:         true,
										DiffSuppressFunc: diffSupressOptional(),
										Computed:         true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"arguments": &schema.Schema{
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"command": &schema.Schema{
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
												"grant": &schema.Schema{
													Description:      `Allowed values: PERMIT, DENY, DENY_ALWAYS`,
													Type:             schema.TypeString,
													Optional:         true,
													DiffSuppressFunc: diffSupressOptional(),
													Computed:         true,
												},
											},
										},
									},
								},
							},
						},
						"description": &schema.Schema{
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
						"name": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"permit_unmatched": &schema.Schema{
							Type:             schema.TypeString,
							ValidateFunc:     validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:         true,
							DiffSuppressFunc: diffSupressBool(),
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceTacacsCommandSetsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TacacsCommandSets create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	isEnableAutoImport := clientConfig.EnableAutoImport
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestTacacsCommandSetsCreateTacacsCommandSets(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if isEnableAutoImport {
		if okID && vvID != "" {
			getResponse1, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByID(vvID)
			if err == nil && getResponse1 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceTacacsCommandSetsRead(ctx, d, m)
			}
		}
		if okName && vvName != "" {
			getResponse2, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByName(vvName)
			if err == nil && getResponse2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = getResponse2.TacacsCommandSets.ID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceTacacsCommandSetsRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.TacacsCommandSets.CreateTacacsCommandSets(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTacacsCommandSets", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTacacsCommandSets", err))
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
	return resourceTacacsCommandSetsRead(ctx, d, m)
}

func resourceTacacsCommandSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TacacsCommandSets read for id=[%s]", d.Id())
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
		log.Printf("[DEBUG] Selected method: GetTacacsCommandSetsByName")
		vvName := vName

		response1, restyResp1, err := client.TacacsCommandSets.GetTacacsCommandSetsByName(vvName)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItemName1 := flattenTacacsCommandSetsGetTacacsCommandSetsByNameItemName(response1.TacacsCommandSets)
		if err := d.Set("item", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemName1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByName response",
				err))
			return diags
		}
		return diags

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTacacsCommandSetsByID")
		vvID := vID

		response2, restyResp2, err := client.TacacsCommandSets.GetTacacsCommandSetsByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemID2 := flattenTacacsCommandSetsGetTacacsCommandSetsByIDItemID(response2.TacacsCommandSets)
		if err := d.Set("item", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItemID2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTacacsCommandSetsByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceTacacsCommandSetsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TacacsCommandSets update for id=[%s]", d.Id())
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
		getResp, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByName(vvName)
		if err != nil || getResp == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTacacsCommandSetsByName", err,
				"Failure at GetTacacsCommandSetsByName, unexpected response", ""))
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsCommandSets != nil {
			vvID = getResp.TacacsCommandSets.ID
		}
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.TacacsCommandSets.UpdateTacacsCommandSetsByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTacacsCommandSetsByID", err, restyResp1.String(),
					"Failure at UpdateTacacsCommandSetsByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTacacsCommandSetsByID", err,
				"Failure at UpdateTacacsCommandSetsByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceTacacsCommandSetsRead(ctx, d, m)
}

func resourceTacacsCommandSetsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TacacsCommandSets delete for id=[%s]", d.Id())
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
		getResp, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 2 {
		vvName = vName
		getResp, _, err := client.TacacsCommandSets.GetTacacsCommandSetsByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
		//Set value vvID = getResp.
		if getResp.TacacsCommandSets != nil {
			vvID = getResp.TacacsCommandSets.ID
		}
	}
	restyResp1, err := client.TacacsCommandSets.DeleteTacacsCommandSetsByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteTacacsCommandSetsByID", err, restyResp1.String(),
				"Failure at DeleteTacacsCommandSetsByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteTacacsCommandSetsByID", err,
			"Failure at DeleteTacacsCommandSetsByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestTacacsCommandSetsCreateTacacsCommandSets(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSets {
	request := isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSets{}
	request.TacacsCommandSets = expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSets(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSets(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSets {
	request := isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSets{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".permit_unmatched")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".permit_unmatched")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".permit_unmatched")))) {
		request.PermitUnmatched = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".commands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".commands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".commands")))) {
		request.Commands = expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommands(ctx, key+".commands.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommands(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommands {
	request := isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommands{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command_list")))) {
		request.CommandList = expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandListArray(ctx, key+".command_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList {
	request := []isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList {
	request := isegosdk.RequestTacacsCommandSetsCreateTacacsCommandSetsTacacsCommandSetsCommandsCommandList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".grant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".grant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".grant")))) {
		request.Grant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".arguments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".arguments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".arguments")))) {
		request.Arguments = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByID {
	request := isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByID{}
	request.TacacsCommandSets = expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSets(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSets(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSets {
	request := isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSets{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".permit_unmatched")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".permit_unmatched")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".permit_unmatched")))) {
		request.PermitUnmatched = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".commands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".commands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".commands")))) {
		request.Commands = expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommands(ctx, key+".commands.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommands(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommands {
	request := isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommands{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command_list")))) {
		request.CommandList = expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandListArray(ctx, key+".command_list", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandListArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList {
	request := []isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList {
	request := isegosdk.RequestTacacsCommandSetsUpdateTacacsCommandSetsByIDTacacsCommandSetsCommandsCommandList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".grant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".grant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".grant")))) {
		request.Grant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".command")))) {
		request.Command = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".arguments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".arguments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".arguments")))) {
		request.Arguments = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
