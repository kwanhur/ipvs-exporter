package main

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type ipvsCollector struct {
	metrics map[string]*ipvsMetric
}

type ipvsMetric struct {
	Desc    *prometheus.Desc
	ValType prometheus.ValueType
}

//NewIpvsCollector a new ipvsCollector instance
func NewIpvsCollector(namespace string) *ipvsCollector {
	labels := []string{"vip", "vport", "rip", "rport", "protocol"}
	return &ipvsCollector{
		metrics: map[string]*ipvsMetric{
			"ipvs_connections": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "connections"), "ipvs connections counter", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_active_connections": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "active_connections"), "ipvs current active connection", labels, nil),
				ValType: prometheus.GaugeValue,
			},
			"ipvs_inactive_connections": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "inactive_connections"), "ipvs current inactive connection", labels, nil),
				ValType: prometheus.GaugeValue,
			},
			"ipvs_rate_cps": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "rate_cps"), "ipvs new connection counter per second", labels, nil),
				ValType: prometheus.GaugeValue,
			},
			"ipvs_bytes_in": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_in"), "ipvs ingress bytes", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_bytes_out": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_out"), "ipvs egress bytes", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_packets_in": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "packets_in"), "ipvs ingress packets", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_packets_out": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "packets_out"), "ipvs egress packets", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_rate_inbps": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "rate_inbps"), "ipvs ingress rate bits per second", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_rate_outbps": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "rate_outbps"), "ipvs egress rate bits per second", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_rate_inpps": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "rate_inpps"), "ipvs ingress rate packets per second", labels, nil),
				ValType: prometheus.CounterValue,
			},
			"ipvs_rate_outpps": {
				Desc:    prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "rate_outpps"), "ipvs egress rate packets per second", labels, nil),
				ValType: prometheus.CounterValue,
			},
		},
	}
}

//Describe output ipvs metric descriptions
func (c *ipvsCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		ch <- metric.Desc
	}
}

//Collect output ipvs metric values
func (c *ipvsCollector) Collect(ch chan<- prometheus.Metric) {
	ipvs, err := NewIpvsWrapper()
	if err != nil {
		logrus.Errorf("fetch ipvs handler err:%s", err)
		return
	}

	defer ipvs.Close()
	svcs, err := ipvs.GetServices()
	if err != nil {
		logrus.Errorf("fetch ipvs services err:%s", err)
		return
	}
	for _, svc := range svcs {
		labels := []string{svc.Address.String(), strconv.Itoa(int(svc.Port)), "", "", ipvs.Protocol(svc.Protocol)}
		stats := svc.Stats
		metric := c.metrics["ipvs_connections"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.Connections), labels...)
		metric = c.metrics["ipvs_rate_cps"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.CPS), labels...)
		metric = c.metrics["ipvs_bytes_in"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BytesIn), labels...)
		metric = c.metrics["ipvs_bytes_out"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BytesOut), labels...)
		metric = c.metrics["ipvs_packets_in"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PacketsIn), labels...)
		metric = c.metrics["ipvs_packets_out"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PacketsOut), labels...)
		metric = c.metrics["ipvs_rate_inbps"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BPSIn), labels...)
		metric = c.metrics["ipvs_rate_outbps"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BPSOut), labels...)
		metric = c.metrics["ipvs_rate_inpps"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PPSIn), labels...)
		metric = c.metrics["ipvs_rate_outpps"]
		ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PPSOut), labels...)

		dests, err := ipvs.GetDestinations(svc)
		if err != nil {
			logrus.Errorf("fetch destinations err:%s", err)
		} else {
			for _, dest := range dests {
				labels[2] = dest.Address.String()
				labels[3] = strconv.Itoa(int(dest.Port))
				stats := dest.Stats
				metric = c.metrics["ipvs_active_connections"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(dest.ActiveConnections), labels...)
				metric = c.metrics["ipvs_inactive_connections"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(dest.InactiveConnections), labels...)
				metric = c.metrics["ipvs_rate_cps"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.CPS), labels...)
				metric = c.metrics["ipvs_bytes_in"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BytesIn), labels...)
				metric = c.metrics["ipvs_bytes_out"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BytesOut), labels...)
				metric = c.metrics["ipvs_packets_in"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PacketsIn), labels...)
				metric = c.metrics["ipvs_packets_out"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PacketsOut), labels...)
				metric = c.metrics["ipvs_rate_inbps"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BPSIn), labels...)
				metric = c.metrics["ipvs_rate_outbps"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.BPSOut), labels...)
				metric = c.metrics["ipvs_rate_inpps"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PPSIn), labels...)
				metric = c.metrics["ipvs_rate_outpps"]
				ch <- prometheus.MustNewConstMetric(metric.Desc, metric.ValType, float64(stats.PPSOut), labels...)
			}
		}
	}
}
