FROM golang:1.22.5-alpine as builder
RUN apk add --no-cache bash git curl
RUN mkdir /app
WORKDIR /app
COPY . /app
ARG configFile
RUN go build -mod vendor -o main -ldflags="-X main.configFile=${configFile}" cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata curl
RUN mkdir /app
COPY --from=builder /app/main /app
COPY ./configs/config-stage.yaml  /configs/config-stage.yaml
EXPOSE 8080
CMD ["/app/main", "-c", "/configs/config-stage.yaml", "-migrate", "-seed"]