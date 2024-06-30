FROM golang:1.21.6-alpine AS build
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o lab-1

FROM scratch
WORKDIR /app
COPY --from=build /app /usr/bin/server
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app .
ENTRYPOINT ["./lab-1"]