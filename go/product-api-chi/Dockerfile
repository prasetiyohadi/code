FROM golang:1-alpine as builder
WORKDIR /app
COPY . /app
RUN apk --no-cache --no-progress add\
 ca-certificates=20191127-r5 tzdata=2021a-r0\
 && update-ca-certificates\
 && go mod download\
 && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/app main.go\
 && rm -rf /var/cache/apk/*

FROM scratch
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
CMD ["./app"]
