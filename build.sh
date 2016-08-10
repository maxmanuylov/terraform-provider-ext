#!/bin/bash

VERSION="v1.0"

rm -rf bin

export GO15VENDOREXPERIMENT=1

GOOS=darwin  GOARCH=amd64 go build -o bin/macos/terraform-provider-content
GOOS=linux   GOARCH=amd64 go build -o bin/linux/terraform-provider-content
GOOS=windows GOARCH=amd64 go build -o bin/windows/terraform-provider-content.exe

tar czf bin/terraform-provider-content-$VERSION-macos.tar.gz --directory=bin/macos terraform-provider-content
tar czf bin/terraform-provider-content-$VERSION-linux.tar.gz --directory=bin/linux terraform-provider-content
zip     bin/terraform-provider-content-$VERSION-windows.zip -j bin/windows/terraform-provider-content.exe

go run schema/schema.go content bin
