package content

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
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

func updateContentDir(resourceData *schema.ResourceData, _ interface{}) error {
    if !resourceData.HasChange("permissions") {
        return nil
    }

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
    permissionsStr := resourceData.Get("permissions").(string)

    var permissions os.FileMode
    if permissionsStr == "" {
        permissions = 0777
    } else {
        var err error
        permissions, err = _toNumberPermissions(permissionsStr)
        if err != nil {
            return err
        }
    }

    if err := os.MkdirAll(path, permissions); err != nil {
        return err
    }

    if permissionsStr != "" {
        if err := os.Chmod(path, permissions); err != nil {
            return err
        }
    }

    resourceData.Set("dir", path)

    return nil
}
