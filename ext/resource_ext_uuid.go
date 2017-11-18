package ext

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
)

func createUuid(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }
    
    resourceData.Set("value", id)
    resourceData.SetId(id)

    return nil
}

func readUuid(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func deleteUuid(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")
    resourceData.Set("value", "")

    return nil
}
