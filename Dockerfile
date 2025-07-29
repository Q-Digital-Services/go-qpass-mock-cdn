FROM golang AS build

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN go build -o gocdn

FROM scratch

ENV MINIO_BUCKET=""
ENV MINIO_ENDPOINT="http://localhost:9444"
ENV MINIO_REGION="us-east-1"
ENV MINIO_ACCESS_KEY=""
ENV MINIO_SECRET_KEY=""
ENV PORT=80

COPY --from=build --chmod=0755 /usr/src/app/gocdn /usr/bin/

CMD ["/usr/bin/gocdn"]
