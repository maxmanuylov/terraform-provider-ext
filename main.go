package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/maxmanuylov/terraform-provider-content/content"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: content.Provider,
    })
}