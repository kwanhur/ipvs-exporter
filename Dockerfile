FROM        quay.io/prometheus/busybox:latest
MAINTAINER  kwanhur <huang_hua2012@163.com>

COPY ipvs-exporter  /usr/bin/ipvs-exporter
COPY docker-entrypoint.sh /bin/docker-entrypoint.sh

ENV METRICS_ENDPOINT "/metrics"
ENV METRICS_ADDR ":9911"
ENV DEFAULT_METRICS_NS "ipvs"

ENTRYPOINT [ "docker-entrypoint.sh" ]