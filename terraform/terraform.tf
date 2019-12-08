terraform {
  required_version = "0.12.6"

  backend "remote" {
    organization = "connpass-map-api"

    workspaces {
      name = "connpass-map-api-workspace"
    }
  }
}
