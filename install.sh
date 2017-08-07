#!/bin/bash

TERRAFORM_DIR="<Insert your value>"
LOCAL_OS="<Insert your value>"

cp -f bin/ext.json "$HOME/.terraform.d/schemas/ext.json"
cp -f "bin/$LOCAL_OS/terraform-provider-ext" "$TERRAFORM_DIR/terraform-provider-ext"
