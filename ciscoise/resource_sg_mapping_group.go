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

func resourceSgMappingGroup() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on IPToSGTMappingGroup.

- This resource allows the client to update an IP to SGT mapping group by ID.

- This resource deletes an IP to SGT mapping group.

- This resource creates an IP to SGT mapping group.
`,

		CreateContext: resourceSgMappingGroupCreate,
		ReadContext:   resourceSgMappingGroupRead,
		UpdateContext: resourceSgMappingGroupUpdate,
		DeleteContext: resourceSgMappingGroupDelete,
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

						"deploy_to": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set or unless deployType=ALL`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"deploy_type": &schema.Schema{
							Description: `Allowed values:
- ALL,
- ND,
- NDG`,
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
						"sgt": &schema.Schema{
							Description: `Mandatory unless mappingGroup is set`,
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

						"deploy_to": &schema.Schema{
							Description:      `Mandatory unless mappingGroup is set or unless deployType=ALL`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"deploy_type": &schema.Schema{
							Description: `Allowed values:
		- ALL,
		- ND,
		- NDG`,
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: diffSupressOptional(),
							Computed:         true,
						},
						"id": &schema.Schema{
							Description: `id path parameter.`,
							Type:        schema.TypeString,
							Required:    true,
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
						"sgt": &schema.Schema{
							Description:      `Mandatory unless mappingGroup is set`,
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

func resourceSgMappingGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroup create")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSgMappingGroupCreateIPToSgtMappingGroup(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vID, okID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName, _ := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okID && vvID != "" {
		getResponse2, _, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroupByID(vvID)
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			resourceMap["name"] = vvName
			d.SetId(joinResourceID(resourceMap))
			return resourceSgMappingGroupRead(ctx, d, m)
		}
	} else {
		queryParams2 := isegosdk.GetIPToSgtMappingGroupQueryParams{}

		response2, _, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroup(&queryParams2)
		if response2 != nil && err == nil {
			items2 := getAllItemsIPToSgtMappingGroupGetIPToSgtMappingGroup(m, response2, &queryParams2)
			item2, err := searchIPToSgtMappingGroupGetIPToSgtMappingGroup(m, items2, vvName, vvID)
			if err == nil && item2 != nil {
				resourceMap := make(map[string]string)
				resourceMap["id"] = vvID
				resourceMap["name"] = vvName
				d.SetId(joinResourceID(resourceMap))
				return resourceSgMappingGroupRead(ctx, d, m)
			}
		}
	}
	restyResp1, err := client.IPToSgtMappingGroup.CreateIPToSgtMappingGroup(request1)
	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateIPToSgtMappingGroup", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateIPToSgtMappingGroup", err))
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
	return resourceSgMappingGroupRead(ctx, d, m)
}

func resourceSgMappingGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroup read for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 2 {
		vvName := vName
		vvID := vID
		log.Printf("[DEBUG] Selected method: GetIPToSgtMappingGroup")
		queryParams1 := isegosdk.GetIPToSgtMappingGroupQueryParams{}

		response1, restyResp1, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroup(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		items1 := getAllItemsIPToSgtMappingGroupGetIPToSgtMappingGroup(m, response1, &queryParams1)
		item1, err := searchIPToSgtMappingGroupGetIPToSgtMappingGroup(m, items1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		vItem1 := flattenIPToSgtMappingGroupGetIPToSgtMappingGroupByIDItem(item1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMappingGroup search response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMappingGroup search response",
				err))
			return diags
		}

	}
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetIPToSgtMappingGroupByID")
		vvID := vID

		response2, restyResp2, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroupByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenIPToSgtMappingGroupGetIPToSgtMappingGroupByIDItem(response2.SgMappingGroup)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMappingGroupByID response",
				err))
			return diags
		}
		if err := d.Set("parameters", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetIPToSgtMappingGroupByID response",
				err))
			return diags
		}
		return diags

	}
	return diags
}

func resourceSgMappingGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroup update for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// NOTE: Added getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetIPToSgtMappingGroupQueryParams{}
		getResp1, _, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroup(&queryParams1)
		if err == nil && getResp1 != nil {
			items1 := getAllItemsIPToSgtMappingGroupGetIPToSgtMappingGroup(m, getResp1, &queryParams1)
			item1, err := searchIPToSgtMappingGroupGetIPToSgtMappingGroup(m, items1, vName, vID)
			if err == nil && item1 != nil {
				vvID = vID
			}
		}
	}
	if selectedMethod == 1 {
		vvID = vID
	}
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestSgMappingGroupUpdateIPToSgtMappingGroupByID(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		response1, restyResp1, err := client.IPToSgtMappingGroup.UpdateIPToSgtMappingGroupByID(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateIPToSgtMappingGroupByID", err, restyResp1.String(),
					"Failure at UpdateIPToSgtMappingGroupByID, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateIPToSgtMappingGroupByID", err,
				"Failure at UpdateIPToSgtMappingGroupByID, unexpected response", ""))
			return diags
		}
		_ = d.Set("last_updated", getUnixTimeString())
	}

	return resourceSgMappingGroupRead(ctx, d, m)
}

func resourceSgMappingGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning SgMappingGroup delete for id=[%s]", d.Id())
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID, okID := resourceMap["id"]
	vName, okName := resourceMap["name"]

	method1 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params
	if selectedMethod == 2 {
		queryParams1 := isegosdk.GetIPToSgtMappingGroupQueryParams{}

		getResp1, _, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroup(&queryParams1)
		if err != nil || getResp1 == nil {
			// Assume that element it is already gone
			return diags
		}
		items1 := getAllItemsIPToSgtMappingGroupGetIPToSgtMappingGroup(m, getResp1, &queryParams1)
		item1, err := searchIPToSgtMappingGroupGetIPToSgtMappingGroup(m, items1, vName, vID)
		if err != nil || item1 == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	if selectedMethod == 1 {
		vvID = vID
		getResp, _, err := client.IPToSgtMappingGroup.GetIPToSgtMappingGroupByID(vvID)
		if err != nil || getResp == nil {
			// Assume that element it is already gone
			return diags
		}
	}
	restyResp1, err := client.IPToSgtMappingGroup.DeleteIPToSgtMappingGroupByID(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteIPToSgtMappingGroupByID", err, restyResp1.String(),
				"Failure at DeleteIPToSgtMappingGroupByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteIPToSgtMappingGroupByID", err,
			"Failure at DeleteIPToSgtMappingGroupByID, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSgMappingGroupCreateIPToSgtMappingGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingGroupCreateIPToSgtMappingGroup {
	request := isegosdk.RequestIPToSgtMappingGroupCreateIPToSgtMappingGroup{}
	request.SgMappingGroup = expandRequestSgMappingGroupCreateIPToSgtMappingGroupSgMappingGroup(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgMappingGroupCreateIPToSgtMappingGroupSgMappingGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingGroupCreateIPToSgtMappingGroupSgMappingGroup {
	request := isegosdk.RequestIPToSgtMappingGroupCreateIPToSgtMappingGroupSgMappingGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_to")))) {
		request.DeployTo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_type")))) {
		request.DeployType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgMappingGroupUpdateIPToSgtMappingGroupByID(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID {
	request := isegosdk.RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByID{}
	request.SgMappingGroup = expandRequestSgMappingGroupUpdateIPToSgtMappingGroupByIDSgMappingGroup(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSgMappingGroupUpdateIPToSgtMappingGroupByIDSgMappingGroup(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDSgMappingGroup {
	request := isegosdk.RequestIPToSgtMappingGroupUpdateIPToSgtMappingGroupByIDSgMappingGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt")))) {
		first, _ := replaceRegExStrings(interfaceToString(v), "", `\s*\(.*\)$`, "")
		request.Sgt = first
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_to")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_to")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_to")))) {
		request.DeployTo = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".deploy_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".deploy_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".deploy_type")))) {
		request.DeployType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func getAllItemsIPToSgtMappingGroupGetIPToSgtMappingGroup(m interface{}, response *isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroup, queryParams *isegosdk.GetIPToSgtMappingGroupQueryParams) []isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResources {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var respItems []isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResources
	for response.SearchResult != nil && response.SearchResult.Resources != nil && len(*response.SearchResult.Resources) > 0 {
		respItems = append(respItems, *response.SearchResult.Resources...)
		if response.SearchResult.NextPage != nil && response.SearchResult.NextPage.Rel == "next" {
			href := response.SearchResult.NextPage.Href
			page, size, err := getNextPageAndSizeParams(href)
			if err != nil {
				break
			}
			if queryParams != nil {
				queryParams.Page = page
				queryParams.Size = size
			}
			response, _, err = client.IPToSgtMappingGroup.GetIPToSgtMappingGroup(queryParams)
			if err != nil {
				break
			}
			// All is good, continue to the next page
			continue
		}
		// Does not have next page finish iteration
		break
	}
	return respItems
}

func searchIPToSgtMappingGroupGetIPToSgtMappingGroup(m interface{}, items []isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupSearchResultResources, name string, id string) (*isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByIDSgMappingGroup, error) {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client
	var err error
	var foundItem *isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByIDSgMappingGroup
	for _, item := range items {
		if id != "" && item.ID == id {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByID
			getItem, _, err = client.IPToSgtMappingGroup.GetIPToSgtMappingGroupByID(id)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetIPToSgtMappingGroupByID")
			}
			foundItem = getItem.SgMappingGroup
			return foundItem, err
		} else if name != "" && item.Name == name {
			// Call get by _ method and set value to foundItem and return
			var getItem *isegosdk.ResponseIPToSgtMappingGroupGetIPToSgtMappingGroupByID
			getItem, _, err = client.IPToSgtMappingGroup.GetIPToSgtMappingGroupByID(item.ID)
			if err != nil {
				return foundItem, err
			}
			if getItem == nil {
				return foundItem, fmt.Errorf("Empty response from %s", "GetIPToSgtMappingGroupByID")
			}
			foundItem = getItem.SgMappingGroup
			return foundItem, err
		}
	}
	return foundItem, err
}
