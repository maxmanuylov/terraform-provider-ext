package ext

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
    "os"
)

func createLocalDir(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    permissionsStr := resourceData.Get("permissions").(string)
    if permissionsStr == "" {
        permissionsStr = "777"
    }

    permissions, err := _toNumberPermissions(permissionsStr)
    if err != nil {
        return err
    }

    path := resourceData.Get("path").(string)

    if err := os.RemoveAll(path); err != nil {
        return err
    }

    if err := os.MkdirAll(path, permissions); err != nil {
        return err
    }

    resourceData.Set("dir", path)
    resourceData.SetId(id)

    return nil
}

func readLocalDir(resourceData *schema.ResourceData, _ interface{}) error {
    return _readLocalDirOrFile("dir", resourceData)
}

func updateLocalDir(resourceData *schema.ResourceData, _ interface{}) error {
    if resourceData.HasChange("permissions") {
        return _updateLocalDirOrFilePermissions(resourceData)
    }

    return nil
}

func deleteLocalDir(resourceData *schema.ResourceData, _ interface{}) error {
    return _deleteLocalDirOrFile("dir", resourceData)
}
