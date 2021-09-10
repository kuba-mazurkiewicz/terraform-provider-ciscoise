package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceEgressMatrixCellClone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEgressMatrixCellCloneRead,
		Schema: map[string]*schema.Schema{
			"dst_sgt_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"src_sgt_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"result_value": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceEgressMatrixCellCloneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vSrcSgtID := d.Get("src_sgt_id")
	vDstSgtID := d.Get("dst_sgt_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: CloneMatrixCell")
		vvID := vID.(string)
		vvSrcSgtID := vSrcSgtID.(string)
		vvDstSgtID := vDstSgtID.(string)

		response1, _, err := client.EgressMatrixCell.CloneMatrixCell(vvID, vvSrcSgtID, vvDstSgtID)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CloneMatrixCell", err,
				"Failure at CloneMatrixCell, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenEgressMatrixCellCloneMatrixCellItem(response1.OperationResult)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CloneMatrixCell response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEgressMatrixCellCloneMatrixCellItem(item *isegosdk.ResponseEgressMatrixCellCloneMatrixCellOperationResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["result_value"] = flattenEgressMatrixCellCloneMatrixCellItemResultValue(item.ResultValue)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEgressMatrixCellCloneMatrixCellItemResultValue(items *[]isegosdk.ResponseEgressMatrixCellCloneMatrixCellOperationResultResultValue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["value"] = item.Value
		respItem["name"] = item.Name
	}
	return respItems

}