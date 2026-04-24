FROM golang:alpine as builder

RUN apk add --no-cache git

ADD . $GOPATH/src/github.com/crowdstrike/shell2http
WORKDIR $GOPATH/src/github.com/crowdstrike/shell2http

ENV CGO_ENABLED=0
ENV GOARCH=$TARGETARCH
ENV GOOS=linux

RUN go build -v -trimpath -ldflags="-w -s -X 'main.version=$(git describe --abbrev=0 --tags | sed s/v//)'" -o /go/bin/shell2http .
RUN go build -v -trimpath -o /go/bin/gendetections ./scripts/gendetections

# Generate detections.json from the base image's scripts
FROM quay.io/crowdstrike/detection-container AS detector
COPY --from=builder /go/bin/gendetections /gendetections
RUN /gendetections /home/eval/bin > /detections.json

# final image
FROM quay.io/crowdstrike/detection-container

LABEL org.opencontainers.image.source="https://github.com/CrowdStrike/vulnapp"

COPY --from=detector /detections.json /detections.json
COPY --from=builder /go/bin/shell2http /shell2http
COPY entrypoint.sh /
COPY images /images

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
