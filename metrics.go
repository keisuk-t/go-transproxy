package transproxy

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	defaultListenPort = ":3134"
)

type Exporter struct {
	Nodename          prometheus.Labels
	ProxyHttpTotal    *prometheus.CounterVec
	ProxyHttpsTotal   *prometheus.CounterVec
	NoproxyHttpTotal  *prometheus.CounterVec
	NoproxyHttpsTotal *prometheus.CounterVec
}

func NewExporter() (*Exporter, error) {
	transproxyLabels := []string{"nodename"}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return &Exporter{
		Nodename: prometheus.Labels{
			"nodename": hostname,
		},
		ProxyHttpTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "transproxy_proxy_http_total",
			Help: "transproxy counter of forwarded http packets",
		}, transproxyLabels),
		ProxyHttpsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "transproxy_proxy_https_total",
			Help: "transproxy counter of forwarded https packets",
		}, transproxyLabels),
		NoproxyHttpTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "transproxy_noproxy_http_total",
			Help: "transproxy counter of not forwarded http packets",
		}, transproxyLabels),
		NoproxyHttpsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "transproxy_noproxy_https_total",
			Help: "transproxy counter of not forwarded https packets",
		}, transproxyLabels),
	}, nil
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.ProxyHttpTotal.Describe(ch)
	e.ProxyHttpsTotal.Describe(ch)
	e.NoproxyHttpTotal.Describe(ch)
	e.NoproxyHttpsTotal.Describe(ch)
}
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.ProxyHttpTotal.Collect(ch)
	e.ProxyHttpsTotal.Collect(ch)
	e.NoproxyHttpTotal.Collect(ch)
	e.NoproxyHttpsTotal.Collect(ch)
}
