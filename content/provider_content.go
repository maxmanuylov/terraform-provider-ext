package content

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{},
        ResourcesMap: map[string]*schema.Resource{

            "content_by_url": {
                Schema: map[string]*schema.Schema{
                    "url": {
                        Type: schema.TypeString,
                        Required: true,
                    },
                    "content": {
                        Type: schema.TypeString,
                        Computed: true,
                    },
                },
                Create: createContentByUrl,
                Read:   readContentByUrl,
                Update: updateContentByUrl,
                Delete: deleteContentByUrl,
                Exists: contentByUrlExists,
            },

            "content_file": {
                Schema: map[string]*schema.Schema{
                    "path": {
                        Type: schema.TypeString,
                        Required: true,
                        ForceNew: true,
                    },
                    "content": {
                        Type: schema.TypeString,
                        Required: true,
                    },
                    "file": {
                        Type: schema.TypeString,
                        Computed: true,
                    },
                },
                Create: createContentFile,
                Read:   readContentFile,
                Update: updateContentFile,
                Delete: deleteContentFile,
                Exists: contentFileExists,
            },

        },
    }
}
