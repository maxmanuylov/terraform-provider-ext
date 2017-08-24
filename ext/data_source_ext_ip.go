package ext

import (
    "fmt"
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
    "net"
)

func readIP(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    _, ipNet, err := net.ParseCIDR(resourceData.Get("subnet_cidr").(string))
    if err != nil {
        return fmt.Errorf("Invalid subnet CIDR: %v", err)
    }

    ips := resourceData.Get("values").([]interface{})

    for _, ipValue := range ips {
        if ipStr, ok := ipValue.(string); ok {
            if ip := net.ParseIP(ipStr); ip != nil && ipNet.Contains(ip) {
                resourceData.Set("value", ipStr)
                resourceData.SetId(id)

                return nil
            }
        }
    }

    return fmt.Errorf("No suitable IP address found: %v", ips)
}
