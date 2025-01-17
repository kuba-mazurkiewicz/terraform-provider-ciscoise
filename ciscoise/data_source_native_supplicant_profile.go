package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNativeSupplicantProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on NativeSupplicantProfile.

- This data source allows the client to get a native supplicant profile by ID.

- This data source allows the client to get all the native supplicant profiles.
`,

		ReadContext: dataSourceNativeSupplicantProfileRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"page": &schema.Schema{
				Description: `page query parameter. Page number`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. Number of objects returned per page`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
						"wireless_profiles": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"action_type": &schema.Schema{
										Description: `Action type for WifiProfile.
Allowed values:
- ADD,
- UPDATE,
- DELETE
(required for updating existing WirelessProfile)`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"allowed_protocol": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"certificate_template_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"previous_ssid": &schema.Schema{
										Description: `Previous ssid for WifiProfile (required for updating existing WirelessProfile)`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
					},
				},
			},
		},
	}
}

func dataSourceNativeSupplicantProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNativeSupplicantProfile")
		queryParams1 := isegosdk.GetNativeSupplicantProfileQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}

		response1, restyResp1, err := client.NativeSupplicantProfile.GetNativeSupplicantProfile(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNativeSupplicantProfile", err,
				"Failure at GetNativeSupplicantProfile, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources
		for response1.SearchResult != nil && response1.SearchResult.Resources != nil && len(*response1.SearchResult.Resources) > 0 {
			items1 = append(items1, *response1.SearchResult.Resources...)
			if response1.SearchResult.NextPage != nil && response1.SearchResult.NextPage.Rel == "next" {
				href := response1.SearchResult.NextPage.Href
				page, size, err := getNextPageAndSizeParams(href)
				if err != nil {
					break
				}
				queryParams1.Page = page
				queryParams1.Size = size
				response1, _, err = client.NativeSupplicantProfile.GetNativeSupplicantProfile(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenNativeSupplicantProfileGetNativeSupplicantProfileItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNativeSupplicantProfile response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNativeSupplicantProfileByID")
		vvID := vID.(string)

		response2, restyResp2, err := client.NativeSupplicantProfile.GetNativeSupplicantProfileByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetNativeSupplicantProfileByID", err,
				"Failure at GetNativeSupplicantProfileByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItem(response2.ERSNSpProfile)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetNativeSupplicantProfileByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNativeSupplicantProfileGetNativeSupplicantProfileItems(items *[]isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenNativeSupplicantProfileGetNativeSupplicantProfileItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNativeSupplicantProfileGetNativeSupplicantProfileItemsLink(item *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileSearchResultResourcesLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItem(item *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["wireless_profiles"] = flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItemWirelessProfiles(item.WirelessProfiles)
	respItem["link"] = flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItemLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItemWirelessProfiles(items *[]isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfileWirelessProfiles) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid"] = item.SSID
		respItem["allowed_protocol"] = item.AllowedProtocol
		respItem["certificate_template_id"] = item.CertificateTemplateID
		respItem["action_type"] = item.ActionType
		respItem["previous_ssid"] = item.PreviousSSID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNativeSupplicantProfileGetNativeSupplicantProfileByIDItemLink(item *isegosdk.ResponseNativeSupplicantProfileGetNativeSupplicantProfileByIDERSNSpProfileLink) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rel"] = item.Rel
	respItem["href"] = item.Href
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
