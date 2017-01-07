package ext

import (
    "github.com/hashicorp/go-uuid"
    "github.com/hashicorp/terraform/helper/schema"
    "io/ioutil"
    "net/http"
)

func createUrlContent(resourceData *schema.ResourceData, _ interface{}) error {
    id, err := uuid.GenerateUUID()
    if err != nil {
        return err
    }

    if err = _fetchUrlContent(resourceData); err != nil {
        return err
    }

    resourceData.SetId(id)

    return nil
}

func readUrlContent(_ *schema.ResourceData, _ interface{}) error {
    return nil
}

func updateUrlContent(resourceData *schema.ResourceData, _ interface{}) error {
    if resourceData.HasChange("url") {
        return _fetchUrlContent(resourceData)
    }

    return nil
}

func deleteUrlContent(resourceData *schema.ResourceData, _ interface{}) error {
    resourceData.SetId("")
    resourceData.Set("content", "")

    return nil
}

func _fetchUrlContent(resourceData *schema.ResourceData) error {
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
