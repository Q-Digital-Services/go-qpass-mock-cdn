# Cloudfront emulation for minio

## Env Variables:

| Var Name        | Default Value | Description                                         |
|-----------------|---------------|-----------------------------------------------------|
| BUCKET    | (required)    | Name of the MinIO,localstack or s3 bucket to serve objects from.     |
| ENDPOINT  | (required)    | URL endpoint of the MinIO server (e.g., http://localhost:9444). Must include protocol. |
| REGION    | us-east-1     | AWS region used for signing requests to MinIO.      |
| ACCESS_KEY| (required)    | Access key ID to authenticate with MinIO.           |
| SECRET_KEY| (required)    | Secret access key to authenticate with MinIO.       |
| PORT           | 8080          | Port number where the HTTP server listens.           |


## Build Docker

> NOTE if pushing upon `master` the whople process is automated via github actions.

### Login

```
# Replace mypassword with your own password of github access token
export GH_TOKEN="mypassword" 
# Replace with your username
export GH_USERNAME="myuser"
echo $GH_TOKEN | docker login ghcr.io -u ${GH_USERNAME} --password-stdin
```

### Publish

```
docker build -t ghcr.io/q-digital-services/go-qpass-mock-cdn .
docker push ghcr.io/q-digital-services/go-qpass-mock-cdn
```

## Testing locally

In order to setup dependencies and run it lorally run:

```
bash testrun.sh
```

This exposes some basic env variables, install dependencies and runs the cdn for local development.
Testrun uses port `9080` instead of default 8080.

