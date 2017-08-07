#!/bin/bash

VERSION="v1.4"

govendor sync

rm -rf bin

export GO15VENDOREXPERIMENT=1
export GOARCH=amd64

GOOS=darwin  go build -o bin/macos/terraform-provider-ext
GOOS=linux   go build -o bin/linux/terraform-provider-ext
GOOS=windows go build -o bin/windows/terraform-provider-ext.exe

tar czf bin/terraform-provider-ext-$VERSION-macos.tar.gz --directory=bin/macos terraform-provider-ext
tar czf bin/terraform-provider-ext-$VERSION-linux.tar.gz --directory=bin/linux terraform-provider-ext
zip     bin/terraform-provider-ext-$VERSION-windows.zip -j bin/windows/terraform-provider-ext.exe

go run schema/schema.go ext bin
