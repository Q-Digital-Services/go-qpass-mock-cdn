#!/bin/bash

export BUCKET=q-pass-public-test
export ENDPOINT=http://localhost:4566
export ACCESS_KEY=admin
export SECRET_KEY=Secure123$
export PORT=9080
go mod vendor
go run gocdn.go