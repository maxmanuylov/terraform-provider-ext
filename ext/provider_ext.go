package ext

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{},
        ResourcesMap: map[string]*schema.Resource{

            "ext_ceph_key": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type: schema.TypeString,
                        Computed: true,
                        Sensitive: true,
                    },
                },
                Create: createCephKey,
                Read:   readCephKey,
                Delete: deleteCephKey,
            },

            "ext_local_dir": {
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
                Create: createLocalDir,
                Read:   readLocalDir,
                Update: updateLocalDir,
                Delete: deleteLocalDir,
            },

            "ext_local_file": {
                Schema: map[string]*schema.Schema{
                    "path": {
                        Type: schema.TypeString,
                        Required: true,
                        ForceNew: true,
                    },
                    "content": {
                        Type: schema.TypeString,
                        Required: true,
                        Sensitive: true,
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
                Create: createLocalFile,
                Read:   readLocalFile,
                Update: updateLocalFile,
                Delete: deleteLocalFile,
            },

            "ext_remote": {
                Schema: getRemoteSchema(),
                Create: createRemote,
                Read:   readRemote,
                Update: updateRemote,
                Delete: deleteRemote,
            },

            "ext_cached_var": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type: schema.TypeString,
                        Required: true,
                    },
                    "cached": {
                        Type: schema.TypeString,
                        Computed: true,
                    },
                    "trigger": {
                        Type: schema.TypeString,
                        Optional: true,
                    },
                },
                Create: createCachedVar,
                Read:   readCachedVar,
                Update: updateCachedVar,
                Delete: deleteCachedVar,
            },

            "ext_cached_svar": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type: schema.TypeString,
                        Required: true,
                        Sensitive: true,
                    },
                    "cached": {
                        Type: schema.TypeString,
                        Computed: true,
                        Sensitive: true,
                    },
                    "trigger": {
                        Type: schema.TypeString,
                        Optional: true,
                    },
                },
                Create: createCachedVar,
                Read:   readCachedVar,
                Update: updateCachedVar,
                Delete: deleteCachedVar,
            },

            "ext_uuid": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type:     schema.TypeString,
                        Computed: true,
                    },
                },
                Create: createUuid,
                Read:   readUuid,
                Delete: deleteUuid,
            },

        },

        DataSourcesMap: map[string]*schema.Resource{

            "ext_remote": {
                Schema: getRemoteSchema(),
                Read:   createRemote,
            },

            "ext_ip": {
                Schema: map[string]*schema.Schema{
                    "values": {
                        Type: schema.TypeList,
                        Required: true,
                        Elem: &schema.Schema{
                            Type: schema.TypeString,
                        },
                    },
                    "subnet_cidr": {
                        Type: schema.TypeString,
                        Required: true,
                    },
                    "value": {
                        Type: schema.TypeString,
                        Computed: true,
                    },
                },
                Read:   readIP,
            },

            "ext_var": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type: schema.TypeString,
                        Required: true,
                    },
                },
                Read:   readExtVar,
            },

            "ext_svar": {
                Schema: map[string]*schema.Schema{
                    "value": {
                        Type: schema.TypeString,
                        Required: true,
                        Sensitive: true,
                    },
                },
                Read:   readExtVar,
            },

        },
    }
}

func getRemoteSchema() map[string]*schema.Schema {
    return map[string]*schema.Schema{
        "url": {
            Type: schema.TypeString,
            Required: true,
        },
        "content": {
            Type: schema.TypeString,
            Computed: true,
            Sensitive: true,
        },
    }
}
