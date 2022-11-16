package monitor

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"gitlab.spatio-inc.com/clotho/social-base/env"
)

var (
	localReg    = prometheus.NewRegistry()
	coFqNameMap sync.Map
	gaFqNameMap sync.Map
)

func EmitCounter(name string, help string, value float64, tags map[string]string) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	tagName := make([]string, 0)
	for k := range tags {
		tagName = append(tagName, k)
	}

	// same name collector should register only once
	tmp, ok := coFqNameMap.Load(name)
	if !ok {
		v := prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: name,
			Help: help,
			ConstLabels: map[string]string{
				"cluster_name": env.GetClusterName(),
				"pod_name":     env.GetPodName(),
				"service_name": env.GetServiceName(),
				"env":          env.GetEnv(),
			},
		}, tagName)
		err := localReg.Register(v)
		if err != nil {
			return
		}
		coFqNameMap.Store(name, v)
		tmp = v
	}
	vec, ok := tmp.(*prometheus.CounterVec)
	if !ok || vec == nil {
		return
	}
	vec.With(tags).Add(value)
}

func EmitGauge(name string, help string, value float64, tags map[string]string) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	tagName := make([]string, 0)
	for k := range tags {
		tagName = append(tagName, k)
	}

	tmp, ok := gaFqNameMap.Load(name)
	if !ok {
		v := prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: name,
			Help: help,
			ConstLabels: map[string]string{
				"cluster_name": env.GetClusterName(),
				"pod_name":     env.GetPodName(),
				"service_name": env.GetServiceName(),
				"env":          env.GetEnv(),
			},
		}, tagName)
		err := localReg.Register(v)
		if err != nil {
			return
		}
		gaFqNameMap.Store(name, v)
		tmp = v
	}
	vec, ok := tmp.(*prometheus.GaugeVec)
	if !ok || vec == nil {
		return
	}
	vec.With(tags).Set(value)
}

func EmitHistogram(name string, help string, value float64, tags map[string]string) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	tagName := make([]string, 0)
	for k := range tags {
		tagName = append(tagName, k)
	}

	tmp, ok := gaFqNameMap.Load(name)
	if !ok {
		v := prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name: name,
			Help: help,
			ConstLabels: map[string]string{
				"cluster_name": env.GetClusterName(),
				"pod_name":     env.GetPodName(),
				"service_name": env.GetServiceName(),
				"env":          env.GetEnv(),
			},
		}, tagName)
		err := localReg.Register(v)
		if err != nil {
			return
		}
		gaFqNameMap.Store(name, v)
		tmp = v
	}
	vec, ok := tmp.(*prometheus.HistogramVec)
	if !ok || vec == nil {
		return
	}
	vec.With(tags).Observe(value)
}

func EmitSummary(name string, help string, value float64, tags map[string]string) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	tagName := make([]string, 0)
	for k := range tags {
		tagName = append(tagName, k)
	}

	tmp, ok := gaFqNameMap.Load(name)
	if !ok {
		v := prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Name: name,
			Help: help,
			ConstLabels: map[string]string{
				"cluster_name": env.GetClusterName(),
				"pod_name":     env.GetPodName(),
				"service_name": env.GetServiceName(),
				"env":          env.GetEnv(),
			},
		}, tagName)
		err := localReg.Register(v)
		if err != nil {
			return
		}
		gaFqNameMap.Store(name, v)
		tmp = v
	}
	vec, ok := tmp.(*prometheus.SummaryVec)
	if !ok || vec == nil {
		return
	}
	vec.With(tags).Observe(value)
}
