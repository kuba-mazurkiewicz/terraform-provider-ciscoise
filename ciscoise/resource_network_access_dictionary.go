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

func resourceNetworkAccessDictionary() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Access - Dictionary.

- Network Access Create a new Dictionary.

- Network Access Update a Dictionary.

- Network Access Delete a Dictionary.
`,

		CreateContext: resourceNetworkAccessDictionaryCreate,
		ReadContext:   resourceNetworkAccessDictionaryRead,
		UpdateContext: resourceNetworkAccessDictionaryUpdate,
		DeleteContext: resourceNetworkAccessDictionaryDelete,
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

						"description": &schema.Schema{
							Description: `The description of the Dictionary`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"dictionary_attr_type": &schema.Schema{
							Description: `The dictionary attribute type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Identifier for the dictionary`,
							Type:        schema.TypeString,
							Computed:    true,
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
							Description: `The dictionary name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `The dictionary version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description:      `The description of the Dictionary`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"dictionary_attr_type": &schema.Schema{
							Description:      `The dictionary attribute type`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Description:      `Identifier for the dictionary`,
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
							Description:      `The dictionary name`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"version": &schema.Schema{
							Description:      `The dictionary version`,
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

func resourceNetworkAccessDictionaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessDictionary create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionaries(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vName, okName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vID, _ := resourceItem["id"]
	vvID := interfaceToString(vID)
	if okName && vvName != "" {
		getResponse2, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["name"] = vvName
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceNetworkAccessDictionaryRead(ctx, d, m)
		}
	} else {
		response2, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()
		if response2 != nil && err == nil {
			items2 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, response2)
			item2, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["name"] = vvName
				resourceMap["id"] = vvID
				d.SetId(joinResourceID(resourceMap))
				return resourceNetworkAccessDictionaryRead(ctx, d, m)
			}
		}
	}
	resp1, restyResp1, err := client.NetworkAccessDictionary.CreateNetworkAccessDictionaries(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateNetworkAccessDictionaries", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateNetworkAccessDictionaries", err))
		return diags
	}
	if vvID != resp1.Response.ID {
		vvID = resp1.Response.ID
	}
	if vvName != resp1.Response.Name {
		vvName = resp1.Response.Name
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceNetworkAccessDictionaryRead(ctx, d, m)
}

func resourceNetworkAccessDictionaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessDictionary read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaries")

		response1, restyResp1, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, response1)
		item1, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenNetworkAccessDictionaryGetNetworkAccessDictionaryByNameItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaries search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaries search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryByName")
		vvName := vName

		response2, restyResp2, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNetworkAccessDictionaryGetNetworkAccessDictionaryByNameItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryByName response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaryByName response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceNetworkAccessDictionaryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessDictionary update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvName string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		getResp1, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()
		if err == nil && getResp1 != nil {
			items1 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, getResp1)
			item1, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items1, vName, vID)
			if err == nil && item1 != nil {
				if vName != item1.Name {
					vvName = item1.Name
				} else {
					vvName = vName
				}
			}
		}
	}
	if selectedMethod == 1 {
		vvName = vName
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.NetworkAccessDictionary.UpdateNetworkAccessDictionaryByName(vvName, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateNetworkAccessDictionaryByName", err, restyResp1.String(),
					"Failure at UpdateNetworkAccessDictionaryByName, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateNetworkAccessDictionaryByName", err,
				"Failure at UpdateNetworkAccessDictionaryByName, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceNetworkAccessDictionaryRead(ctx, d, m)
}

func resourceNetworkAccessDictionaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning NetworkAccessDictionary delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, okName := resourceMap["name"]
	vID, okID := resourceMap["id"]

	method1 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvName string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {

		getResp1, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m, getResp1)
		item1, err := searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
		if vName != item1.Name {
			vvName = item1.Name
		} else {
			vvName = vName
		}
	}
	if selectedMethod == 1 {
		vvName = vName
		getResp, _, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	response1, restyResp1, err := client.NetworkAccessDictionary.DeleteNetworkAccessDictionaryByName(vvName)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteNetworkAccessDictionaryByName", err, restyResp1.String(),
				"Failure at DeleteNetworkAccessDictionaryByName, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteNetworkAccessDictionaryByName", err,
			"Failure at DeleteNetworkAccessDictionaryByName, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionaries(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionaries {
	request := isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionaries{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_attr_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_attr_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_attr_type")))) {
		request.DictionaryAttrType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink {
	request := isegosdk.RequestNetworkAccessDictionaryCreateNetworkAccessDictionariesLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName {
	request := isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByName{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dictionary_attr_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dictionary_attr_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dictionary_attr_type")))) {
		request.DictionaryAttrType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink {
	request := isegosdk.RequestNetworkAccessDictionaryUpdateNetworkAccessDictionaryByNameLink{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".href")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".href")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".href")))) {
		request.Href = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rel")))) {
		request.Rel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsNetworkAccessDictionaryGetNetworkAccessDictionaries(m interface{}, response *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaries) []isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse {
	var respItems []isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse
	if response.Response != nil && len(*response.Response) > 0 {
		respItems = append(respItems, *response.Response...)
	}
	return respItems
}

func searchNetworkAccessDictionaryGetNetworkAccessDictionaries(m interface{}, items []isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse, name string, id string) (*isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName
			getItem, _, err = client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(item.Name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessDictionaryByName")
			}
			foundItem = getItem.Response
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByName
			getItem, _, err = client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(name)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetNetworkAccessDictionaryByName")
			}
			foundItem = getItem.Response
			return foundItem, err
		}
	}
	return foundItem, err
}
