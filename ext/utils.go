package ext

import (
    "github.com/hashicorp/terraform/helper/schema"
    "io/ioutil"
    "os"
    "strconv"
)

func _toNumberPermissions(permissionsStr string) (os.FileMode, error) {
    permissions, err := strconv.ParseInt(permissionsStr, 8, 32)
    if err != nil {
        return 0, err
    }

    return os.FileMode(permissions), nil
}

func _toStringPermissions(permissions os.FileMode) string {
    return strconv.FormatInt(int64(permissions), 8)
}

func _readLocalDirOrFile(kind string, resourceData *schema.ResourceData) error {
    path := resourceData.Get("path").(string)
    isFile := kind == "file"

    info, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return _clearLocalDirOrFileData(kind, resourceData)
        }
        return err
    }

    if info.IsDir() == isFile {
        return _deleteLocalDirOrFile(kind, resourceData)
    }

    if isFile {
        content, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }

        resourceData.Set("content", string(content))
    }

    if resourceData.Get("permissions").(string) != "" {
        resourceData.Set("permissions", _toStringPermissions(info.Mode().Perm()))
    }

    return nil
}

func _updateLocalDirOrFilePermissions(resourceData *schema.ResourceData) error {
    permissionsStr := resourceData.Get("permissions").(string)
    if permissionsStr == "" {
        return nil
    }

    permissions, err := _toNumberPermissions(permissionsStr)
    if err != nil {
        return err
    }

    path := resourceData.Get("path").(string)

    return os.Chmod(path, permissions)
}

func _deleteLocalDirOrFile(kind string, resourceData *schema.ResourceData) error {
    path := resourceData.Get("path").(string)

    if err := os.RemoveAll(path); err != nil {
        return err
    }

    return _clearLocalDirOrFileData(kind, resourceData)
}

func _clearLocalDirOrFileData(kind string, resourceData *schema.ResourceData) error {
    resourceData.SetId("")
    resourceData.Set(kind, "")

    return nil
}

