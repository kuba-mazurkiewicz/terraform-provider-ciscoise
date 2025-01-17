package ciscoise

import (
	"context"

	"log"

	isegosdk "github.com/kuba-mazurkiewicz/ciscoise-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRepository() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Repository.

- This will get the full list of repository definitions on the system.

- Get a specific repository identified by the name passed in the URL.
`,

		ReadContext: dataSourceRepositoryRead,
		Schema: map[string]*schema.Schema{
			"repository_name": &schema.Schema{
				Description: `repositoryName path parameter. Unique name for a repository`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"enable_pki": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password can contain alphanumeric and/or special characters.`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"path": &schema.Schema{
							Description: `Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": &schema.Schema{
							Description: `Username may contain alphanumeric and _-./@\\$ characters.`,
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

						"enable_pki": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"password": &schema.Schema{
							Description: `Password can contain alphanumeric and/or special characters.`,
							Type:        schema.TypeString,
							Sensitive:   true,
							Computed:    true,
						},
						"path": &schema.Schema{
							Description: `Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"server_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"user_name": &schema.Schema{
							Description: `Username may contain alphanumeric and _-./@\\$ characters.`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	clientConfig := m.(ClientConfig)
	client := clientConfig.Client

	var diags diag.Diagnostics
	vRepositoryName, okRepositoryName := d.GetOk("repository_name")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okRepositoryName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRepositories")

		response1, restyResp1, err := client.Repository.GetRepositories()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRepositories", err,
				"Failure at GetRepositories, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenRepositoryGetRepositoriesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRepositories response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetRepository")
		vvRepositoryName := vRepositoryName.(string)

		response2, restyResp2, err := client.Repository.GetRepository(vvRepositoryName)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetRepository", err,
				"Failure at GetRepository, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenRepositoryGetRepositoryItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRepository response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenRepositoryGetRepositoriesItems(items *[]isegosdk.ResponseRepositoryGetRepositoriesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["protocol"] = item.Protocol
		respItem["path"] = item.Path
		respItem["password"] = item.Password
		respItem["server_name"] = item.ServerName
		respItem["user_name"] = item.UserName
		respItem["enable_pki"] = boolPtrToString(item.EnablePki)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenRepositoryGetRepositoryItem(item *isegosdk.ResponseRepositoryGetRepositoryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["protocol"] = item.Protocol
	respItem["path"] = item.Path
	respItem["password"] = item.Password
	respItem["server_name"] = item.ServerName
	respItem["user_name"] = item.UserName
	respItem["enable_pki"] = boolPtrToString(item.EnablePki)
	return []map[string]interface{}{
		respItem,
	}
}
