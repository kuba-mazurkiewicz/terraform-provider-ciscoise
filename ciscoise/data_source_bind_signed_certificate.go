package ciscoise

import (
	"context"

	"reflect"

	"github.com/CiscoISE/ciscoise-go-sdk/sdk"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceAction
func dataSourceBindSignedCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBindSignedCertificateRead,
		Schema: map[string]*schema.Schema{
			"admin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_extended_validity": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_out_of_date_cert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_replacement_of_certificates": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_replacement_of_portal_group_tag": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ims": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"message": &schema.Schema{
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"portal": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"portal_group_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pxgrid": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"radius": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"saml": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"validate_certificate_extensions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func dataSourceBindSignedCertificateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*isegosdk.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: BindCsr")
		request1 := expandRequestBindSignedCertificateBindCsr(ctx, "", d)

		response1, _, err := client.Certificates.BindCsr(request1)

		if err != nil || response1 == nil {
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing BindCsr", err,
				"Failure at BindCsr, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", *response1)

		vItem1 := flattenCertificatesBindCsrItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting BindCsr response",
				err))
			return diags
		}
		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func expandRequestBindSignedCertificateBindCsr(ctx context.Context, key string, d *schema.ResourceData) *isegosdk.RequestCertificatesBindCsr {
	request := isegosdk.RequestCertificatesBindCsr{}
	if v, ok := d.GetOkExists(key + ".admin"); !isEmptyValue(reflect.ValueOf(d.Get(key+".admin"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".admin"))) {
		request.Admin = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_extended_validity"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_extended_validity"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_extended_validity"))) {
		request.AllowExtendedValidity = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_out_of_date_cert"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_out_of_date_cert"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_out_of_date_cert"))) {
		request.AllowOutOfDateCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_replacement_of_certificates"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_replacement_of_certificates"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_replacement_of_certificates"))) {
		request.AllowReplacementOfCertificates = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".allow_replacement_of_portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".allow_replacement_of_portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".allow_replacement_of_portal_group_tag"))) {
		request.AllowReplacementOfPortalGroupTag = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".data"); !isEmptyValue(reflect.ValueOf(d.Get(key+".data"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".data"))) {
		request.Data = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".eap"); !isEmptyValue(reflect.ValueOf(d.Get(key+".eap"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".eap"))) {
		request.Eap = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".host_name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".host_name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".host_name"))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".id"); !isEmptyValue(reflect.ValueOf(d.Get(key+".id"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".id"))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".ims"); !isEmptyValue(reflect.ValueOf(d.Get(key+".ims"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".ims"))) {
		request.Ims = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".name"); !isEmptyValue(reflect.ValueOf(d.Get(key+".name"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".name"))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".portal"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal"))) {
		request.Portal = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".portal_group_tag"); !isEmptyValue(reflect.ValueOf(d.Get(key+".portal_group_tag"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".portal_group_tag"))) {
		request.PortalGroupTag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(key + ".pxgrid"); !isEmptyValue(reflect.ValueOf(d.Get(key+".pxgrid"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".pxgrid"))) {
		request.Pxgrid = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".radius"); !isEmptyValue(reflect.ValueOf(d.Get(key+".radius"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".radius"))) {
		request.Radius = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".saml"); !isEmptyValue(reflect.ValueOf(d.Get(key+".saml"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".saml"))) {
		request.Saml = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(key + ".validate_certificate_extensions"); !isEmptyValue(reflect.ValueOf(d.Get(key+".validate_certificate_extensions"))) && (ok || !reflect.DeepEqual(v, d.Get(key+".validate_certificate_extensions"))) {
		request.ValidateCertificateExtensions = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenCertificatesBindCsrItem(item *isegosdk.ResponseCertificatesBindCsrResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["message"] = item.Message
	respItem["status"] = item.Status
	return []map[string]interface{}{
		respItem,
	}
}