package content

import (
    "crypto/rand"
    "encoding/base64"
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
    "math/big"
    "time"
)

func createContentCephKey(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)
    created := time.Now()

    key := make([]byte, 44)

    writeValue(key[0:2], 1) // AES
    writeValue(key[2:6], int(created.Unix()))
    writeValue(key[6:10], created.Nanosecond())
    writeValue(key[10:12], 32) // secret length

    secret := key[12:]

    big255 := big.NewInt(255)
    for i := range secret {
        n, err := rand.Int(rand.Reader, big255)
        if err != nil {
            return err
        }
        secret[i] = 1 + byte(n.Int64())
    }

    resourceData.Set("value", base64.StdEncoding.EncodeToString(key))

    return nil
}

func writeValue(buf []byte, value int) {
    for i := range buf {
        buf[i] = byte((value >> (8 * uint(i))) & 255)
    }
}

func readContentCephKey(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func deleteContentCephKey(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")

    return nil
}

func contentCephKeyExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    return resourceData.Id() != "", nil
}
