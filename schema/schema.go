package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/maxmanuylov/terraform-provider-ext/ext"
    "github.com/maxmanuylov/utils/intellij-hcl/terraform/provider-schema-generator"
)

func main() {
    provider_schema_generator.Generate(ext.Provider().(*schema.Provider))
}
