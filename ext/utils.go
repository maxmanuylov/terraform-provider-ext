package ext

import (
    "strconv"
    "os"
)

func _toNumberPermissions(permissionsStr string) (os.FileMode, error) {
    permissions, err := strconv.ParseInt(permissionsStr, 8, 32)
    if err != nil {
        return 0, err
    }

    return os.FileMode(permissions), nil
}
