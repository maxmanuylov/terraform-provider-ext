package content

import (
    "github.com/hashicorp/terraform/helper/schema"
    "net/http"
    "io/ioutil"
    "github.com/hashicorp/consul/vendor/github.com/hashicorp/go-uuid"
)

func createContentByUrl(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    resourceData.SetId(id)

    if err = _fetchContentByUrl(resourceData); err != nil {
        return err
    }

    return nil
}

func readContentByUrl(resourceData *schema.ResourceData, _ interface{}) error {
    if err := _fetchContentByUrl(resourceData); err != nil {
        return err
    }

    return nil
}

func updateContentByUrl(resourceData *schema.ResourceData, _ interface{}) error {
    if !resourceData.HasChange("url") {
        return nil
    }

    if err := _fetchContentByUrl(resourceData); err != nil {
        return err
    }

    return nil
}

func deleteContentByUrl(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")
    resourceData.Set("content", "")

    return nil
}

func contentByUrlExists(resourceData *schema.ResourceData, _ interface{}) (bool, error) {
    return resourceData.Get("content").(string) != "", nil
}

func _fetchContentByUrl(resourceData *schema.ResourceData) error {
    url := resourceData.Get("url").(string)

    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return err
    }

    request.Header.Add("User-Agent", "curl/7.43.0")
    request.Header.Add("Accept", "*/*")

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        return err
    }

    defer response.Body.Close()

    content, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    resourceData.Set("content", string(content))

    return nil
}
