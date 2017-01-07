# Terraform Extensions Provider

This is a plugin for HashiCorp [Terraform](https://terraform.io), which provides one with abilities to:

* Fetch remote resources by URLs
* Use variables with interpolated values
* Create local files/directories

## Usage

- Download the plugin from [Releases](https://github.com/maxmanuylov/terraform-provider-ext/releases) page.
- [Install](https://terraform.io/docs/plugins/basics.html) it, or put into a directory with configuration files.
- Create a sample configuration file `example.tf`:
```
data "ext_remote" "readme" {
  url = "https://raw.githubusercontent.com/maxmanuylov/terraform-provider-ext/master/README.md"
}

data "ext_var" "readme_storage_path" {
  value = "${path.root}/readme_storage"
}

resource "ext_local_dir" "readme_storage" {
  path = "${data.ext_var.readme_storage_path.value}"
  permissions = "777"
}

resource "ext_local_file" "readme" {
  path = "${ext_local_dir.readme_storage.dir}/README.md"
  content = "${data.ext_remote.readme.content}"
  permissions = "644"
}
```
- Run:
```
$ terraform apply
```

## The "ext_remote" resource / data source type

Fetches resource content from the specified URL and provides it as a computed attribute. Data source performs fetch on every refresh, while resource caches the result and refetches it only in case of URL change. 

### Mandatory Parameters
- `url` - remote resource URL to fetch

### Computed Parameters
- `content` - fetched content

## The "ext_local_dir" resource type

Manages the local directory by the specified path. All parent directories are created as well if needed.

### Mandatory Parameters
- `path` - local directory path

### Optional Parameters
- `permissions` - string with octal directory permissions (e.g. "644")

### Computed Parameters
- `dir` - equal to `path` but is set _after_ the directory is created (so you can depend on this resource and be sure the directory already exists by the time you use it)

## The "ext_local_file" resource type

Saves the specified content to the local file.

### Mandatory Parameters
- `path` - local file path
- `content` - content to save

### Optional Parameters
- `permissions` - string with octal file permissions (e.g. "644")

### Computed Parameters
- `file` - equal to `path` but is set _after_ the file is written (so you can depend on this resource and be sure the file already exists by the time you use it)

## The "ext_var" data source type

Provides simple named variable. Unlike the built-in Terraform variables these variables can have interpolations in its values.

### Mandatory Parameters
- `value` - variable value

## The "ext_svar" data source type

Same as above, but its value is marked as sensitive to be not displayed in logs.
