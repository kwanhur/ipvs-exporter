# Multistage build
# #1 stage to build the Go binary
FROM golang AS builder
WORKDIR /usr/src/app
COPY . .
RUN make build

# #2 stage copies binary from #1 + entrypoint
FROM        quay.io/prometheus/busybox:latest
MAINTAINER  kwanhur <huang_hua2012@163.com>

COPY --from=builder /usr/src/app/ipvs-exporter /usr/bin/ipvs-exporter
COPY docker-entrypoint.sh /bin/docker-entrypoint.sh
RUN chmod 755 /bin/docker-entrypoint.sh

ENV METRICS_ENDPOINT "/metrics"
ENV METRICS_ADDR ":9911"
ENV DEFAULT_METRICS_NS "ipvs"

ENTRYPOINT [ "docker-entrypoint.sh" ]
