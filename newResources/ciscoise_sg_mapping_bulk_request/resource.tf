
resource "ciscoise_sg_mapping_bulk_request" "example" {
  provider = ciscoise

  operation_type      = "string"
  resource_media_type = "string"
}