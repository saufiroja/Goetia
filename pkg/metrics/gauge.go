package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
)

type MetricsGauge struct {
	TotalCPU    *prometheus.GaugeVec
	TotalMemory *prometheus.GaugeVec
}

func NewMetricsGauge(reg prometheus.Registerer, service string) *MetricsGauge {
	m := &MetricsGauge{
		TotalCPU: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: service,
			Name:      "total_cpu",
			Help:      "The total number of cpu",
		}, []string{"version"}),
		TotalMemory: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: service,
			Name:      "total_memory",
			Help:      "The total number of memory",
		}, []string{"version"}),
	}
	reg.MustRegister(m.TotalCPU, m.TotalMemory)
	return m
}

func (m *MetricsGauge) SetTotalCPU() {
	numCpu := runtime.NumCPU()
	m.TotalCPU.WithLabelValues("v1").Set(float64(numCpu))
}

func (m *MetricsGauge) SetTotalMemory() {
	var mStat runtime.MemStats
	runtime.ReadMemStats(&mStat)
	m.TotalMemory.WithLabelValues("v1").Set(float64(mStat.TotalAlloc))
}
