FROM golang AS build

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -o gocdn .

FROM scratch

ENV BUCKET=""
ENV ENDPOINT="http://localhost:9444"
ENV REGION="us-east-1"
ENV ACCESS_KEY=""
ENV SECRET_KEY=""

EXPOSE 8080

COPY --from=build --chmod=0755 /usr/src/app/gocdn /gocdn
CMD ["/gocdn"]
