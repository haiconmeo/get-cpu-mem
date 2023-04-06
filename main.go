package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metric struct {
	Mem float64
	Cpu float64
}

func sendData(w http.ResponseWriter, r *http.Request) {
	v, _ := mem.VirtualMemory()
	percent, _ := cpu.Percent(time.Second, false)

	var result Metric
	result.Mem = v.UsedPercent
	result.Cpu = percent[0]
	a, _ := json.Marshal(result)
	w.Write(a)

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/metric", sendData)
	http.ListenAndServe(":4545", mux)

}
