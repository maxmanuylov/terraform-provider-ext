# Terraform Content Provider

This is a plugin for HashiCorp [Terraform](https://terraform.io), which helps fetching some content from a remote URL and save it (or another content) locally to a file.

## Usage

- Download the plugin from [Releases](https://github.com/maxmanuylov/terraform-provider-content/releases) page.
- [Install](https://terraform.io/docs/plugins/basics.html) it, or put into a directory with configuration files.
- Create a sample configuration file `terraform.tf`:
```
resource "content_by_url" "readme" {
  url = "https://raw.githubusercontent.com/maxmanuylov/terraform-provider-content/master/README.md"
}

resource "content_file" "readme_file" {
  path = "./README.md"
  content = "${content_by_url.readme.content}"
}
```
- Run:
```
$ terraform apply
```

Different resource types can be used independently, so you can, for example, save some static content to the local file using just the "content_file" resource.

## The "content_by_url" resource type

### Mandatory Parameters
- `url` - content URL to fetch

## Computed Parameters
- `content` - fetched content

## The "content_file" resource type

### Mandatory Parameters
- `path` - local file path
- `content` - content to save

## Computed Parameters
- `file` - equal to `path` but is set _after_ file is written (so you can depend on this resource and be sure the file already exists by the time you use it)
