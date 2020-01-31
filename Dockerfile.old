FROM golang:1.12 AS build-env
COPY . /go/src/app
WORKDIR /go/src/app
#RUN go install -i
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main /go/src/app/cmd/main.go

FROM scratch
# WORKDIR /app
COPY --from=build-env /go/src/app/main main
COPY public public
CMD ["/main"]
EXPOSE 8080
