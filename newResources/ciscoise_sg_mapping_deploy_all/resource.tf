resource "ciscoise_sg_mapping_deploy_all" "example" {
  provider = ciscoise
  lifecycle {
    create_before_destroy = true
  }
  parameters{
    
  }
}