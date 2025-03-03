FROM golang:alpine as builder

LABEL org.opencontainers.image.source="https://github.com/CrowdStrike/vulnapp"

RUN apk add --no-cache git

ADD . $GOPATH/src/github.com/crowdstrike/shell2http
WORKDIR $GOPATH/src/github.com/crowdstrike/shell2http

ENV CGO_ENABLED=0
ENV GOARCH=$TARGETARCH
ENV GOOS=linux

RUN go build -v -trimpath -ldflags="-w -s -X 'main.version=$(git describe --abbrev=0 --tags | sed s/v//)'" -o /go/bin/shell2http .

# final image
FROM quay.io/crowdstrike/detection-container

COPY entrypoint.sh /
COPY images /images
COPY --from=builder /go/bin/shell2http /shell2http

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
