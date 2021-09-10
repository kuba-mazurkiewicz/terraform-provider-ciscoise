package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePanHa() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePanHaRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failed_attempts": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"polling_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"primary_health_check_node": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"secondary_health_check_node": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePanHaRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetPanHaStatus")

		response1, _, err := client.PanHa.GetPanHaStatus()

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetPanHaStatus", err,
				"Failure at GetPanHaStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItems1 := flattenPanHaGetPanHaStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPanHaStatus response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPanHaGetPanHaStatusItems(items *[]isegosdk.ResponsePanHaGetPanHaStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["is_enabled"] = item.IsEnabled
		respItem["primary_health_check_node"] = item.PrimaryHealthCheckNode
		respItem["secondary_health_check_node"] = item.SecondaryHealthCheckNode
		respItem["polling_interval"] = item.PollingInterval
		respItem["failed_attempts"] = item.FailedAttempts
		respItems = append(respItems, respItem)
	}
	return respItems
}