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

func resourceTrustsecSgVnMappingBulkUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on sgVnMapping.
- Update SG-VN Mappings in bulk
`,

		CreateContext: resourceTrustsecSgVnMappingBulkUpdateCreate,
		ReadContext:   resourceTrustsecSgVnMappingBulkUpdateRead,
		DeleteContext: resourceTrustsecSgVnMappingBulkUpdateDelete,

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

						"id": &schema.Schema{
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
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestSgVnMappingBulkUpdateSgVnMappings`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Identifier of the SG-VN mapping`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"last_update": &schema.Schema{
										Description: `Timestamp for the last update of the SG-VN mapping`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"sg_name": &schema.Schema{
										Description: `Name of the associated Security Group to be used for identity if id is not provided`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"sgt_id": &schema.Schema{
										Description: `Identifier of the associated Security Group which is required unless its name is provided`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"vn_id": &schema.Schema{
										Description: `Identifier for the associated Virtual Network which is required unless its name is provided`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
									"vn_name": &schema.Schema{
										Description: `Name of the associated Virtual Network to be used for identity if id is not provided`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceTrustsecSgVnMappingBulkUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning BulkUpdateSgVnMappings create")
	log.Printf("[DEBUG] Missing BulkUpdateSgVnMappings create on Cisco ISE. It will only be create it on Terraform")
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	request1 := expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappings(ctx, "parameters.0", d)

	response1, restyResp1, err := client.SgVnMapping.BulkUpdateSgVnMappings(request1)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing BulkUpdateSgVnMappings", err,
			"Failure at BulkUpdateSgVnMappings, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	vItem1 := flattenSgVnMappingBulkUpdateSgVnMappingsItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting BulkUpdateSgVnMappings response",
			err))
		return diags
	}
	_ = d.Set("last_updated", getUnixTimeString())
	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return resourceTrustsecSgVnMappingBulkUpdateRead(ctx, d, m)
}

func resourceTrustsecSgVnMappingBulkUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceTrustsecSgVnMappingBulkUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning TrustsecSgVnMappingBulkUpdate delete for id=[%s]", d.Id())
	var diags diag.Diagnostics
	log.Printf("[DEBUG] Missing TrustsecSgVnMappingBulkUpdate delete on Cisco ISE. It will only be delete it on Terraform id=[%s]", d.Id())
	return diags
}

func expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappings(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestSgVnMappingBulkUpdateSgVnMappings {
	request := isegosdk.RequestSgVnMappingBulkUpdateSgVnMappings{}
	if v := expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings {
	request := []isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings{}
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
		i := expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTrustsecSgVnMappingBulkUpdateBulkUpdateSgVnMappingsItem(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings {
	request := isegosdk.RequestItemSgVnMappingBulkUpdateSgVnMappings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update")))) {
		request.LastUpdate = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sg_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sg_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sg_name")))) {
		request.SgName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sgt_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sgt_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sgt_id")))) {
		request.SgtID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_id")))) {
		request.VnID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vn_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vn_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vn_name")))) {
		request.VnName = interfaceToString(v)
	}
	return &request
}

func flattenSgVnMappingBulkUpdateSgVnMappingsItem(item *isegosdk.ResponseSgVnMappingBulkUpdateSgVnMappings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}
