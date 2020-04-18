#!/bin/sh
set -eo pipefail
METRICS_NS=${METRICS_NS:-$DEFAULT_METRICS_NS}

# If there are any arguments then we want to run those instead
echo "[$0] - Metrics Namespace  --> [$METRICS_NS]"
echo "[$0] - Running metrics ipvs-exporter"

exec ipvs-exporter -telemetry.address $METRICS_ADDR -telemetry.endpoint $METRICS_ENDPOINT -metrics.namespace $METRICS_NS
#fi
