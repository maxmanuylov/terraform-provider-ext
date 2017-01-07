#!/bin/bash

INTELLIJ_PLUGINS_DIR="<Insert your value>"
TERRAFORM_DIR="<Insert your value>"
LOCAL_OS="<Insert your value>"

cp -f bin/ext.json "$INTELLIJ_PLUGINS_DIR/intellij-hcl/classes/terraform/model/providers/ext.json"
cp -f "bin/$LOCAL_OS/terraform-provider-ext" "$TERRAFORM_DIR/terraform-provider-ext"
