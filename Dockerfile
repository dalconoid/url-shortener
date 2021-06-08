FROM golang:1.13-alpine AS build

WORKDIR /go/src/url-shortener

COPY . .
COPY ./config.yaml /go/bin

RUN go install ./...

FROM alpine:3.12
WORKDIR /usr/bin
COPY --from=build /go/bin .

#docker build . -t url_srv
#docker run --link pg_url --rm -p 8085:8085 -d --name url_shortener url_srv url-shortener
#docker kill url_shortener