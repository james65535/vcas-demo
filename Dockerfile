FROM golang:1.10 AS build-env
COPY /cmd /go/src/app/cmd
COPY /pkg /go/src/app/pkg
WORKDIR /go/src/app
#RUN go install -i
RUN set -x && \
    #go get github.com/2tvenom/go-test-teamcity && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main /go/src/app/cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build-env /go/src/app/main /app/main
COPY /public /app/public
ENTRYPOINT /app/main
EXPOSE 8080