FROM scratch
WORKDIR /go/src/app
ADD cmd/main /go/src/app
ADD public /go/src/app
ENTRYPOINT ["/go/src/app/main"]
EXPOSE 8080