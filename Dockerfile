FROM registry.access.redhat.com/ubi9/go-toolset:1.26 AS builder

ARG TARGETARCH

ENV CGO_ENABLED=0
ENV GOARCH=$TARGETARCH
ENV GOOS=linux

# go-toolset runs as UID 1001 (HOME/GOPATH live under /opt/app-root/src).
# Source must be owned by that user, otherwise git rejects the repo as
# "dubious ownership", breaking `git describe` and Go's VCS stamping.
WORKDIR /opt/app-root/src/vulnapp
COPY --chown=1001:0 . .

RUN go build -v -trimpath -ldflags="-w -s -X 'main.version=$(git describe --abbrev=0 --tags | sed s/v//)'" -o /opt/app-root/src/go/bin/shell2http .
RUN go build -v -trimpath -o /opt/app-root/src/go/bin/gendetections ./scripts/gendetections

# Generate detections.json from the base image's scripts
FROM quay.io/crowdstrike/detection-container AS detector
COPY --from=builder /opt/app-root/src/go/bin/gendetections /gendetections
RUN /gendetections /home/eval/bin > /detections.json

# final image
FROM quay.io/crowdstrike/detection-container

LABEL org.opencontainers.image.source="https://github.com/CrowdStrike/vulnapp" \
      license="MIT" \
      vendor="CrowdStrike" \
      url="https://github.com/CrowdStrike/vulnapp" \
      vcs-ref="" \
      release=""

COPY --from=detector /detections.json /detections.json
COPY --from=builder /opt/app-root/src/go/bin/shell2http /shell2http
COPY entrypoint.sh /
COPY images /images

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
