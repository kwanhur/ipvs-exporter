package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
)

var (
	listenAddress   = flag.String("telemetry.address", ":9911", "Address to listen on for telemetry.")
	metricPath      = flag.String("telemetry.endpoint", "/metrics", "Path under which to expose metrics.")
	metricNamespace = flag.String("metrics.namespace", "ipvs", "Prometheus metrics namespace.")
	showVersion     = flag.Bool("version", false, "Show version information and exit.")
	goMetrics       = flag.Bool("go.metrics", false, "Export process and go metrics.")
)

func init() {
	err := prometheus.Register(version.NewCollector("ipvs_exporter"))
	if err != nil {
		logrus.Fatalf("register ipvs_exporter failed:%s", err)
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(version.Print("ipvs_exporter"))
		return
	}
	logrus.Infof("Starting ipvs_exporter %s", version.Info())
	logrus.Infof("Build context %s", version.BuildContext())

	err := prometheus.Register(NewIpvsCollector(*metricNamespace))
	if err != nil {
		logrus.Fatalf("register collector %s failed:%s", *metricNamespace, err)
	}

	if !*goMetrics {
		prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{
			PidFn: func() (int, error) {
				return os.Getpid(), nil
			},
		}))
		prometheus.Unregister(prometheus.NewGoCollector())
	}

	logrus.Infof("Providing metrics at %s%s with namespace %s", *listenAddress, *metricPath, *metricNamespace)
	logrus.Infof("Scraping metrics from %s", *metricPath)

	http.Handle(*metricPath, promhttp.Handler())
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(`<html>
			<head><title>Ipvs Exporter</title></head>
			<body>
			<h1>Ipvs Exporter</h1>
			<p><a href="` + *metricPath + `">Metrics</a></p>
			</body>
			</html>`))
		if err != nil {
			logrus.Fatalf("handle request failed:%s", err)
		}
	})

	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		logrus.Fatalf("Fatal shutdown exporter err:%s", err)
	}
}
