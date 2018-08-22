FROM golang:1.10-stretch AS BUILD_SYSTEM

COPY main.go /microservice/main.go

WORKDIR /microservice

RUN go build -o microservice

FROM scratch

COPY --from=BUILD_SYSTEM /microservice/microservice /microservice

EXPOSE 8080

ENTRYPOINT ['/microservice']


