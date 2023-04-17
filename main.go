package main

import (
	"encoding/json"
	"log"
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
	apiKey := r.Header.Get("x-api-key")
	if !authori(apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	v, _ := mem.VirtualMemory()
	percent, _ := cpu.Percent(time.Second, false)

	var result Metric
	result.Mem = v.UsedPercent
	result.Cpu = percent[0]
	a, _ := json.Marshal(result)
	w.Write(a)

}
func authori(apiKey string) bool {
	return apiKey == "q_EaTiX+xbBXLyO05.+zDXjI+Qi_X0v"

}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/metric", sendData)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}
