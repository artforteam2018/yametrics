package memstats

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
)

type MemStats struct {
	address string
	counter struct {
		Pollcount int
	}
	gauge struct {
		Alloc         uint64
		BuckHashSys   uint64
		Frees         uint64
		GCCPUFraction float64
		GCSys         uint64
		HeapAlloc     uint64
		HeapIdle      uint64
		HeapInuse     uint64
		HeapObjects   uint64
		HeapReleased  uint64
		HeapSys       uint64
		LastGC        uint64
		Lookups       uint64
		MCacheInuse   uint64
		MCacheSys     uint64
		MSpanInuse    uint64
		MSpanSys      uint64
		Mallocs       uint64
		NextGC        uint64
		NumForcedGC   uint32
		NumGC         uint32
		OtherSys      uint64
		PauseTotalNs  uint64
		StackInuse    uint64
		StackSys      uint64
		Sys           uint64
		TotalAlloc    uint64
		RandomValue   int
	}
	metrics runtime.MemStats
}

func (m *MemStats) incr() {
	m.counter.Pollcount++
	m.gauge.RandomValue = rand.Int()
}

func (m *MemStats) Init() {
	m.counter.Pollcount = 0
	m.gauge.RandomValue = rand.Int()
	m.metrics = runtime.MemStats{}
}

func (m *MemStats) Scan() {
	runtime.ReadMemStats(&m.metrics)

	m.gauge.Alloc = m.metrics.Alloc
	m.gauge.BuckHashSys = m.metrics.BuckHashSys
	m.gauge.Frees = m.metrics.Frees
	m.gauge.GCCPUFraction = m.metrics.GCCPUFraction
	m.gauge.GCSys = m.metrics.GCSys
	m.gauge.HeapAlloc = m.metrics.HeapAlloc
	m.gauge.HeapIdle = m.metrics.HeapIdle
	m.gauge.HeapInuse = m.metrics.HeapInuse
	m.gauge.HeapObjects = m.metrics.HeapObjects
	m.gauge.HeapReleased = m.metrics.HeapReleased
	m.gauge.HeapSys = m.metrics.HeapSys
	m.gauge.LastGC = m.metrics.LastGC
	m.gauge.Lookups = m.metrics.Lookups
	m.gauge.MCacheInuse = m.metrics.MCacheInuse
	m.gauge.MCacheSys = m.metrics.MCacheSys
	m.gauge.MSpanInuse = m.metrics.MSpanInuse
	m.gauge.MSpanSys = m.metrics.MSpanSys
	m.gauge.Mallocs = m.metrics.Mallocs
	m.gauge.NextGC = m.metrics.NextGC
	m.gauge.NumForcedGC = m.metrics.NumForcedGC
	m.gauge.NumGC = m.metrics.NumGC
	m.gauge.OtherSys = m.metrics.OtherSys
	m.gauge.PauseTotalNs = m.metrics.PauseTotalNs
	m.gauge.StackInuse = m.metrics.StackInuse
	m.gauge.StackSys = m.metrics.StackSys
	m.gauge.Sys = m.metrics.Sys
	m.gauge.TotalAlloc = m.metrics.TotalAlloc

	m.incr()
}

func (m MemStats) sendHTTPGauge(fieldName string, value string) {
	m.sendHTTP(fieldName, value, m.address+"update/gauge/")
}

func (m MemStats) sendHTTPCounter(fieldName string, value string) {
	m.sendHTTP(fieldName, value, m.address+"update/counter/")
}

func (m MemStats) sendHTTP(fieldName string, value string, addressWithType string) {
	fmt.Println("send to address", addressWithType+fieldName+"/"+value)

	response, err := http.Post(addressWithType+fieldName+"/"+value, "text/plain", nil)
	if err != nil {
		fmt.Printf("memstats sending error: %v\n", err)
		return
	} else {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("response close error: %v\n", err)
			return
		}
	}
}

func (m MemStats) Send() {
	m.sendHTTPGauge("Alloc", strconv.FormatUint(m.gauge.Alloc, 10))
	m.sendHTTPGauge("BuckHashSys", strconv.FormatUint(m.gauge.BuckHashSys, 10))
	m.sendHTTPGauge("Frees", strconv.FormatUint(m.gauge.Frees, 10))
	m.sendHTTPGauge("GCCPUFraction", strconv.FormatFloat(m.gauge.GCCPUFraction, 'f', -1, 64))
	m.sendHTTPGauge("GCSys", strconv.FormatUint(m.gauge.GCSys, 10))
	m.sendHTTPGauge("HeapAlloc", strconv.FormatUint(m.gauge.HeapAlloc, 10))
	m.sendHTTPGauge("HeapIdle", strconv.FormatUint(m.gauge.HeapIdle, 10))
	m.sendHTTPGauge("HeapInuse", strconv.FormatUint(m.gauge.HeapInuse, 10))
	m.sendHTTPGauge("HeapObjects", strconv.FormatUint(m.gauge.HeapObjects, 10))
	m.sendHTTPGauge("HeapReleased", strconv.FormatUint(m.gauge.HeapReleased, 10))
	m.sendHTTPGauge("HeapSys", strconv.FormatUint(m.gauge.HeapSys, 10))
	m.sendHTTPGauge("LastGC", strconv.FormatUint(m.gauge.LastGC, 10))
	m.sendHTTPGauge("Lookups", strconv.FormatUint(m.gauge.Lookups, 10))
	m.sendHTTPGauge("MCacheInuse", strconv.FormatUint(m.gauge.MCacheInuse, 10))
	m.sendHTTPGauge("MCacheSys", strconv.FormatUint(m.gauge.MCacheSys, 10))
	m.sendHTTPGauge("MSpanInuse", strconv.FormatUint(m.gauge.MSpanInuse, 10))
	m.sendHTTPGauge("MSpanSys", strconv.FormatUint(m.gauge.MSpanSys, 10))
	m.sendHTTPGauge("Mallocs", strconv.FormatUint(m.gauge.Mallocs, 10))
	m.sendHTTPGauge("NextGC", strconv.FormatUint(m.gauge.NextGC, 10))
	m.sendHTTPGauge("NumForcedGC", strconv.FormatUint(uint64(m.gauge.NumForcedGC), 10))
	m.sendHTTPGauge("NumGC", strconv.FormatUint(uint64(m.gauge.NumGC), 10))
	m.sendHTTPGauge("OtherSys", strconv.FormatUint(m.gauge.OtherSys, 10))
	m.sendHTTPGauge("PauseTotalNs", strconv.FormatUint(m.gauge.PauseTotalNs, 10))
	m.sendHTTPGauge("StackInuse", strconv.FormatUint(m.gauge.StackInuse, 10))
	m.sendHTTPGauge("StackSys", strconv.FormatUint(m.gauge.StackSys, 10))
	m.sendHTTPGauge("Sys", strconv.FormatUint(m.gauge.Sys, 10))
	m.sendHTTPGauge("TotalAlloc", strconv.FormatUint(m.gauge.TotalAlloc, 10))
	m.sendHTTPCounter("PollCount", strconv.FormatUint(uint64(m.counter.Pollcount), 10))
	m.sendHTTPGauge("RandomValue", strconv.FormatUint(uint64(m.gauge.RandomValue), 10))

}

func InitMemstats(address string) MemStats {
	stats := MemStats{}
	stats.address = "http://" + address + "/"
	return stats
}
