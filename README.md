# Cloudfront emulation for minio

## Env Variables:

| Var Name        | Default Value | Description                                         |
|-----------------|---------------|-----------------------------------------------------|
| MINIO_BUCKET    | (required)    | Name of the MinIO bucket to serve objects from.     |
| MINIO_ENDPOINT  | (required)    | URL endpoint of the MinIO server (e.g., http://localhost:9444). Must include protocol. |
| MINIO_REGION    | us-east-1     | AWS region used for signing requests to MinIO.      |
| MINIO_ACCESS_KEY| (required)    | Access key ID to authenticate with MinIO.           |
| MINIO_SECRET_KEY| (required)    | Secret access key to authenticate with MinIO.       |
| PORT           | 8080          | Port number where the HTTP server listens.           |


## Build Docker

```
 docker build -t ddesyllas924/mockcdn .
 docker push desyllas924/mockcdn
```

