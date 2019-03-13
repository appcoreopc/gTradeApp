FROM golang:1.12-alpine AS builder

ADD https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

RUN apk add --no-cache git
RUN apk update && apk add bash

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/appcoreopc/gTradeApp
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM golang:1.12-alpine
COPY --from=builder /app ./
CMD ["./app"]

# sudo docker build -t kepung/gtradeapp . 