package ciscoise

import (
	"context"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSgACLBulkMonitorStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSgACLBulkMonitorStatusRead,
		Schema: map[string]*schema.Schema{
			"bulkid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bulk_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fail_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"media_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"operation_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resources_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"resources_status": &schema.Schema{
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
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_execution_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"success_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSgACLBulkMonitorStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics
	vBulkid := d.Get("bulkid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: MonitorBulkStatusSecurityGroupsACL")
		vvBulkid := vBulkid.(string)

		response1, _, err := client.SecurityGroupsACLs.MonitorBulkStatusSecurityGroupsACL(vvBulkid)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing MonitorBulkStatusSecurityGroupsACL", err,
				"Failure at MonitorBulkStatusSecurityGroupsACL, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLItem(response1.BulkStatus)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting MonitorBulkStatusSecurityGroupsACL response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLItem(item *isegosdk.ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLBulkStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["bulk_id"] = item.BulkID
	respItem["media_type"] = item.MediaType
	respItem["execution_status"] = item.ExecutionStatus
	respItem["operation_type"] = item.OperationType
	respItem["start_time"] = item.StartTime
	respItem["resources_count"] = item.ResourcesCount
	respItem["success_count"] = item.SuccessCount
	respItem["fail_count"] = item.FailCount
	respItem["resources_status"] = flattenSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLItemResourcesStatus(item.ResourcesStatus)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLItemResourcesStatus(items *[]isegosdk.ResponseSecurityGroupsACLsMonitorBulkStatusSecurityGroupsACLBulkStatusResourcesStatus) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["resource_execution_status"] = item.ResourceExecutionStatus
		respItem["status"] = item.Status
	}
	return respItems

}