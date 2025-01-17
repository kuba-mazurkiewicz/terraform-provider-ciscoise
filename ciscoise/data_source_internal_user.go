package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInternalUser() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on InternalUser.

- This data source allows the client to get an internal user by name.

- This data source allows the client to get an internal user by ID.

- This data source allows the client to get all the internal users.

Filter:

[firstName, lastName, identityGroup, name, description, email, enabled]

Sorting:

[name, description]
`,

		ReadContext: dataSourceInternalUserRead,
		Schema: map[string]*schema.Schema{
			"filter": &schema.Schema{
				Description: `filter query parameter. 

**Simple filtering** should be available through the filter query string parameter. The structure of a filter is
a triplet of field operator and value separated with dots. More than one filter can be sent. The logical operator
common to ALL filter criteria will be by default AND, and can be changed by using the "filterType=or" query
string parameter. Each resource Data model description should specify if an attribute is a filtered field.



              Operator    | Description 

              ------------|----------------

              EQ          | Equals 

              NEQ         | Not Equals 

              GT          | Greater Than 

              LT          | Less Then 

              STARTSW     | Starts With 

              NSTARTSW    | Not Starts With 

              ENDSW       | Ends With 

              NENDSW      | Not Ends With 

              CONTAINS	  | Contains 

              NCONTAINS	  | Not Contains 

`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"filter_type": &schema.Schema{
				Description: `filterType query parameter. The logical operator common to ALL filter criteria will be by default AND, and can be changed by using the parameter`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name path parameter.`,
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
			"sortasc": &schema.Schema{
				Description: `sortasc query parameter. sort asc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sortdsc": &schema.Schema{
				Description: `sortdsc query parameter. sort desc`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item_id": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"change_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_attributes": &schema.Schema{
							Description: `Key value map`,
							// CHECK: The type of this param
							// Replaced List to Map
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": &schema.Schema{
							Description: `Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'.
The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiry_date": &schema.Schema{
							Description: `To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"expiry_date_enabled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_groups": &schema.Schema{
							Description: `CSV of identity group IDs`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_name": &schema.Schema{
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
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"password_idstore": &schema.Schema{
							Description: `The id store where the internal user's password is kept`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
					},
				},
			},
			"item_name": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"change_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_attributes": &schema.Schema{
							Description: `Key value map`,
							// CHECK: The type of this param
							// Replaced List to Map
							Type:     schema.TypeMap,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"email": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable_password": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": &schema.Schema{
							Description: `Whether the user is enabled/disabled. To use it as filter, the values should be 'Enabled' or 'Disabled'.
The values are case sensitive. For example, '[ERSObjectURL]?filter=enabled.EQ.Enabled'`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"expiry_date": &schema.Schema{
							Description: `To store the internal user's expiry date information. It's format is = 'YYYY-MM-DD'`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"expiry_date_enabled": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"first_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_groups": &schema.Schema{
							Description: `CSV of identity group IDs`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"last_name": &schema.Schema{
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
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"password_idstore": &schema.Schema{
							Description: `The id store where the internal user's password is kept`,
							Type:        schema.TypeString,
							Sensitive:   true,
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

func dataSourceInternalUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vPage, okPage := d.GetOk("page")
	vSize, okSize := d.GetOk("size")
	vSortasc, okSortasc := d.GetOk("sortasc")
	vSortdsc, okSortdsc := d.GetOk("sortdsc")
	vFilter, okFilter := d.GetOk("filter")
	vFilterType, okFilterType := d.GetOk("filter_type")
	vName, okName := d.GetOk("name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okPage, okSize, okSortasc, okSortdsc, okFilter, okFilterType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)
	method3 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 3 %v", method3)

	selectedMethod := pickMethod([][]bool{method1, method2, method3})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetInternalUser")
		queryParams1 := isegosdk.GetInternalUserQueryParams{}

		if okPage {
			queryParams1.Page = vPage.(int)
		}
		if okSize {
			queryParams1.Size = vSize.(int)
		}
		if okSortasc {
			queryParams1.Sortasc = vSortasc.(string)
		}
		if okSortdsc {
			queryParams1.Sortdsc = vSortdsc.(string)
		}
		if okFilter {
			queryParams1.Filter = interfaceToSliceString(vFilter)
		}
		if okFilterType {
			queryParams1.FilterType = vFilterType.(string)
		}

		response1, restyResp1, err := client.InternalUser.GetInternalUser(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUser", err,
				"Failure at GetInternalUser, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		var items1 []isegosdk.ResponseInternalUserGetInternalUserSearchResultResources
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
				response1, _, err = client.InternalUser.GetInternalUser(&queryParams1)
				if err != nil {
					break
				}
				// All is good, continue to the next page
				continue
			}
			// Does not have next page finish iteration
			break
		}
		vItems1 := flattenInternalUserGetInternalUserItems(&items1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUser response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetInternalUserByName")
		vvName := vName.(string)

		response2, restyResp2, err := client.InternalUser.GetInternalUserByName(vvName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByName", err,
				"Failure at GetInternalUserByName, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItemName2 := flattenInternalUserGetInternalUserByNameItemName(response2.InternalUser)
		if err := d.Set("item_name", vItemName2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUserByName response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 3 {
		log.Printf("[DEBUG] Selected method: GetInternalUserByID")
		vvID := vID.(string)

		response3, restyResp3, err := client.InternalUser.GetInternalUserByID(vvID)

		if err != nil || response3 == nil {
			if restyResp3 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp3.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetInternalUserByID", err,
				"Failure at GetInternalUserByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response3))

		vItemID3 := flattenInternalUserGetInternalUserByIDItemID(response3.InternalUser)
		if err := d.Set("item_id", vItemID3); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetInternalUserByID response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenInternalUserGetInternalUserItems(items *[]isegosdk.ResponseInternalUserGetInternalUserSearchResultResources) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["link"] = flattenInternalUserGetInternalUserItemsLink(item.Link)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenInternalUserGetInternalUserItemsLink(item *isegosdk.ResponseInternalUserGetInternalUserSearchResultResourcesLink) []map[string]interface{} {
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

func flattenInternalUserGetInternalUserByNameItemName(item *isegosdk.ResponseInternalUserGetInternalUserByNameInternalUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["enabled"] = boolPtrToString(item.Enabled)
	respItem["email"] = item.Email
	respItem["password"] = item.Password
	respItem["first_name"] = item.FirstName
	respItem["last_name"] = item.LastName
	respItem["change_password"] = boolPtrToString(item.ChangePassword)
	respItem["identity_groups"] = item.IDentityGroups
	respItem["expiry_date_enabled"] = boolPtrToString(item.ExpiryDateEnabled)
	respItem["expiry_date"] = item.ExpiryDate
	respItem["enable_password"] = item.EnablePassword
	respItem["custom_attributes"] = flattenInternalUserGetInternalUserByNameItemNameCustomAttributes(item.CustomAttributes)
	respItem["password_idstore"] = item.PasswordIDStore
	respItem["link"] = flattenInternalUserGetInternalUserByNameItemNameLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenInternalUserGetInternalUserByNameItemNameCustomAttributes(item *isegosdk.ResponseInternalUserGetInternalUserByNameInternalUserCustomAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return respItem

}

func flattenInternalUserGetInternalUserByNameItemNameLink(item *isegosdk.ResponseInternalUserGetInternalUserByNameInternalUserLink) []map[string]interface{} {
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

func flattenInternalUserGetInternalUserByIDItemID(item *isegosdk.ResponseInternalUserGetInternalUserByIDInternalUser) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["enabled"] = boolPtrToString(item.Enabled)
	respItem["email"] = item.Email
	respItem["password"] = item.Password
	respItem["first_name"] = item.FirstName
	respItem["last_name"] = item.LastName
	respItem["change_password"] = boolPtrToString(item.ChangePassword)
	respItem["identity_groups"] = item.IDentityGroups
	respItem["expiry_date_enabled"] = boolPtrToString(item.ExpiryDateEnabled)
	respItem["expiry_date"] = item.ExpiryDate
	respItem["enable_password"] = item.EnablePassword
	respItem["custom_attributes"] = flattenInternalUserGetInternalUserByIDItemIDCustomAttributes(item.CustomAttributes)
	respItem["password_idstore"] = item.PasswordIDStore
	respItem["link"] = flattenInternalUserGetInternalUserByIDItemIDLink(item.Link)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenInternalUserGetInternalUserByIDItemIDCustomAttributes(item *isegosdk.ResponseInternalUserGetInternalUserByIDInternalUserCustomAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return respItem

}

func flattenInternalUserGetInternalUserByIDItemIDLink(item *isegosdk.ResponseInternalUserGetInternalUserByIDInternalUserLink) []map[string]interface{} {
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
