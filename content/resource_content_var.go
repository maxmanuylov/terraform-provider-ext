package content

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/consul/vendor/github.com/hashicorp/go-uuid"
)

func createContentVar(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)

    return nil
}

func readContentVar(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func updateContentVar(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func deleteContentVar(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")

    return nil
}

func contentVarExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    return resourceData.Id() != "", nil
}
