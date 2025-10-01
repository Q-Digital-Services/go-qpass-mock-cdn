#!/bin/bash

export MINIO_BUCKET=q-pass-public-test
export MINIO_ENDPOINT=http://localhost:4566
export MINIO_ACCESS_KEY=admin
export MINIO_SECRET_KEY=Secure123$
export PORT=9080
go mod vendor
go run gocdn.go