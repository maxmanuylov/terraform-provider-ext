package ext

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
)

func createCachedVar(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    cacheValue(resourceData)
    
    resourceData.SetId(id)

    return nil
}

func readCachedVar(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func updateCachedVar(resourceData *schema.ResourceData, _ interface{}) error {
    if resourceData.HasChange("trigger") {
        cacheValue(resourceData)
    }

    return nil
}

func deleteCachedVar(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")
    resourceData.Set("cached", "")

    return nil
}

func cacheValue(resourceData *schema.ResourceData) {
    resourceData.Set("cached", resourceData.Get("value"))
}
