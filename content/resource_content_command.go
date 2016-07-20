package content

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
)

func createContentCommand(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)

    return nil
}

func readContentCommand(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func deleteContentCommand(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")

    return nil
}

func contentCommandExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    return resourceData.Id() != "", nil
}
