FROM golang:1.14-alpine AS build

WORKDIR /build

ADD ./fetching-service .
RUN go get -v
RUN CGO=0 go build -o api

FROM alpine

ENV DOCKERIZE_VERSION v0.6.1

RUN apk add --no-cache openssl
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

WORKDIR /usr/bin

COPY --from=build /build/api .
RUN chmod +x api

CMD ["api", "serve"]
