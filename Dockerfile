FROM golang:1.11 AS build-env
COPY . /app
RUN cd /app && GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o main cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build-env /app/main /app/main
ENTRYPOINT /app/main
EXPOSE 8080