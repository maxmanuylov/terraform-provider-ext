package content

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/consul/vendor/github.com/hashicorp/go-uuid"
    "os"
)

func createContentDir(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)

    if err = _deleteContentDir(resourceData); err != nil {
        return err
    }

    if err = _createContentDir(resourceData); err != nil {
        return err
    }

    return nil
}

func readContentDir(resourceData *schema.ResourceData, _ interface{}) error {
    if err := _createContentDir(resourceData); err != nil {
        return err
    }

    return nil
}

func deleteContentDir(resourceData *schema.ResourceData, _ interface{}) error {
    if err := _deleteContentDir(resourceData); err != nil {
        return err
    }

    resourceData.SetId("")

    return nil
}

func contentDirExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    info, err := os.Stat(resourceData.Get("path").(string))
    if err != nil && os.IsNotExist(err) {
        return false, nil
    }

    return err == nil && info.IsDir(), err
}

func _deleteContentDir(resourceData *schema.ResourceData) error {
    path := resourceData.Get("path").(string)

    if err := os.RemoveAll(path); err != nil {
        return err
    }

    resourceData.Set("dir", "")

    return nil
}

func _createContentDir(resourceData *schema.ResourceData) error {
    path := resourceData.Get("path").(string)

    if err := os.MkdirAll(path, 0777); err != nil {
        return err
    }

    resourceData.Set("dir", path)

    return nil
}