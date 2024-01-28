package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	devices prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myapp",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices.",
		}),
	}
	reg.MustRegister(m.devices)
	return m
}
