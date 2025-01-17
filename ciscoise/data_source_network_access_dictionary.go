package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkAccessDictionary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Access - Dictionary.

- Get all Dictionaries.

- GET a dictionary by name
`,

		ReadContext: dataSourceNetworkAccessDictionaryRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name path parameter. the dictionary name`,
				Type:        schema.TypeString,
				Optional:    true,
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
			"items": &schema.Schema{
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
		},
	}
}

func dataSourceNetworkAccessDictionaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaries")

		response1, restyResp1, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaries()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaries", err,
				"Failure at GetNetworkAccessDictionaries, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkAccessDictionaryGetNetworkAccessDictionariesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNetworkAccessDictionaries response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAccessDictionaryByName")
		vvName := vName.(string)

		response2, restyResp2, err := client.NetworkAccessDictionary.GetNetworkAccessDictionaryByName(vvName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNetworkAccessDictionaryByName", err,
				"Failure at GetNetworkAccessDictionaryByName, unexpected response", ""))
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
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkAccessDictionaryGetNetworkAccessDictionariesItems(items *[]isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["description"] = item.Description
		respItem["dictionary_attr_type"] = item.DictionaryAttrType
		respItem["id"] = item.ID
		respItem["link"] = flattenNetworkAccessDictionaryGetNetworkAccessDictionariesItemsLink(item.Link)
		respItem["name"] = item.Name
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkAccessDictionaryGetNetworkAccessDictionariesItemsLink(item *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionariesResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkAccessDictionaryGetNetworkAccessDictionaryByNameItem(item *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["description"] = item.Description
	respItem["dictionary_attr_type"] = item.DictionaryAttrType
	respItem["id"] = item.ID
	respItem["link"] = flattenNetworkAccessDictionaryGetNetworkAccessDictionaryByNameItemLink(item.Link)
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkAccessDictionaryGetNetworkAccessDictionaryByNameItemLink(item *isegosdk.ResponseNetworkAccessDictionaryGetNetworkAccessDictionaryByNameResponseLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["href"] = item.Href
	respItem["rel"] = item.Rel
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
