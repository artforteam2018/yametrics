package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
)

const ServerAddress = "http://localhost:8080/"

type MemStats struct {
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
		Pollcount     int
	}
	metrics     runtime.MemStats
	RandomValue int
}

func (m *MemStats) incr() {
	m.gauge.Pollcount++
	m.RandomValue = rand.Int()
}

func (m *MemStats) Init() {
	m.gauge.Pollcount = 0
	m.RandomValue = rand.Int()
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

func sendHttp(fieldName string, value string) {
	response, err := http.Post(ServerAddress+"update/gauge/"+fieldName+"/"+value, "text/plain", nil)
	if err != nil {
		fmt.Println(response)
		fmt.Printf("memstats sending error: %v\n", err)
		return
	}
}

func (m MemStats) Send() {

	sendHttp("Alloc", strconv.FormatUint(m.gauge.Alloc, 10))
	sendHttp("BuckHashSys", strconv.FormatUint(m.gauge.BuckHashSys, 10))
	sendHttp("Frees", strconv.FormatUint(m.gauge.Frees, 10))
	sendHttp("GCCPUFraction", strconv.FormatFloat(m.gauge.GCCPUFraction, 'f', 12, 64))
	sendHttp("GCSys", strconv.FormatUint(m.gauge.GCSys, 10))
	sendHttp("HeapAlloc", strconv.FormatUint(m.gauge.HeapAlloc, 10))
	sendHttp("HeapIdle", strconv.FormatUint(m.gauge.HeapIdle, 10))
	sendHttp("HeapInuse", strconv.FormatUint(m.gauge.HeapInuse, 10))
	sendHttp("HeapObjects", strconv.FormatUint(m.gauge.HeapObjects, 10))
	sendHttp("HeapReleased", strconv.FormatUint(m.gauge.HeapReleased, 10))
	sendHttp("HeapSys", strconv.FormatUint(m.gauge.HeapSys, 10))
	sendHttp("LastGC", strconv.FormatUint(m.gauge.LastGC, 10))
	sendHttp("Lookups", strconv.FormatUint(m.gauge.Lookups, 10))
	sendHttp("MCacheInuse", strconv.FormatUint(m.gauge.MCacheInuse, 10))
	sendHttp("MCacheSys", strconv.FormatUint(m.gauge.MCacheSys, 10))
	sendHttp("MSpanInuse", strconv.FormatUint(m.gauge.MSpanInuse, 10))
	sendHttp("MSpanSys", strconv.FormatUint(m.gauge.MSpanSys, 10))
	sendHttp("Mallocs", strconv.FormatUint(m.gauge.Mallocs, 10))
	sendHttp("NextGC", strconv.FormatUint(m.gauge.NextGC, 10))
	sendHttp("NumForcedGC", strconv.FormatUint(uint64(m.gauge.NumForcedGC), 10))
	sendHttp("NumGC", strconv.FormatUint(uint64(m.gauge.NumGC), 10))
	sendHttp("OtherSys", strconv.FormatUint(m.gauge.OtherSys, 10))
	sendHttp("PauseTotalNs", strconv.FormatUint(m.gauge.PauseTotalNs, 10))
	sendHttp("StackInuse", strconv.FormatUint(m.gauge.StackInuse, 10))
	sendHttp("StackSys", strconv.FormatUint(m.gauge.StackSys, 10))
	sendHttp("Sys", strconv.FormatUint(m.gauge.Sys, 10))
	sendHttp("TotalAlloc", strconv.FormatUint(m.gauge.TotalAlloc, 10))
	sendHttp("Pollcount", strconv.FormatUint(uint64(m.gauge.Pollcount), 10))

}
