package content

import (
    "crypto/rand"
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
    "math/big"
)

func createContentPassword(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)

    password := make([]byte, resourceData.Get("length").(int))

    big94 := big.NewInt(94)
    for i := range password {
        n, err := rand.Int(rand.Reader, big94)
        if err != nil {
            return err
        }
        password[i] = 33 + byte(n.Int64())
    }

    resourceData.Set("value", string(password))

    return nil
}

func readContentPassword(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func updateContentPassword(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func deleteContentPassword(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")

    return nil
}

func contentPasswordExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    return resourceData.Id() != "", nil
}
