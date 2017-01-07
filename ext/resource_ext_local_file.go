package ext

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
    "os"
)

func createLocalFile(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    path := resourceData.Get("path").(string)

    if err := os.RemoveAll(path); err != nil {
        return err
    }

    if err := _saveLocalFileContent(resourceData); err != nil {
        return err
    }

    if err := _updateLocalDirOrFilePermissions(resourceData); err != nil {
        return err
    }

    resourceData.Set("file", path)
    resourceData.SetId(id)

    return nil
}

func readLocalFile(resourceData *schema.ResourceData, _ interface{}) error {
    return _readLocalDirOrFile("file", resourceData)
}

func updateLocalFile(resourceData *schema.ResourceData, _ interface{}) error {
    if resourceData.HasChange("content") {
        if err := _saveLocalFileContent(resourceData); err != nil {
            return err
        }
    }

    if resourceData.HasChange("permissions") {
        if err := _updateLocalDirOrFilePermissions(resourceData); err != nil {
            return err
        }
    }

    return nil
}

func deleteLocalFile(resourceData *schema.ResourceData, _ interface{}) error {
    return _deleteLocalDirOrFile("file", resourceData)
}

func _saveLocalFileContent(resourceData *schema.ResourceData) error {
    file, err := os.Create(resourceData.Get("path").(string))
    if err != nil {
        return err
    }

    defer file.Close()
    defer file.Sync()

    _, err = file.WriteString(resourceData.Get("content").(string))

    return err
}
