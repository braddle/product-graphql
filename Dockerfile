# BUILD
FROM golang:latest as build

WORKDIR /service
ADD . /service

RUN cd /service && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /gql-service .

CMD /gql-service

# TEST
FROM build as test

# PRODUCTION
FROM alpine:latest as production

RUN apk --no-cache add ca-certificates
COPY --from=build /service ./
RUN chmod +x ./gql-service

ENTRYPOINT ["./gql-service"]

EXPOSE 8080