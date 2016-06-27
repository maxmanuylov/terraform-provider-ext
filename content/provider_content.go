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

            "content_dir": {
                Schema: map[string]*schema.Schema{
                    "path": {
                        Type: schema.TypeString,
                        Required: true,
                        ForceNew: true,
                    },
                    "permissions": {
                        Type: schema.TypeString,
                        Optional: true,
                    },
                    "dir": {
                        Type: schema.TypeString,
                        Computed: true,
                    },
                },
                Create: createContentDir,
                Read:   readContentDir,
                Update: updateContentDir,
                Delete: deleteContentDir,
                Exists: contentDirExists,
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
                    "permissions": {
                        Type: schema.TypeString,
                        Optional: true,
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

            "content_var": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type: schema.TypeString,
                        Required: true,
                    },
                },
                Create: createContentVar,
                Read:   readContentVar,
                Update: updateContentVar,
                Delete: deleteContentVar,
                Exists: contentVarExists,
            },

        },
    }
}
