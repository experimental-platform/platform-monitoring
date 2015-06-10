package main

import (
	"bitbucket.org/bertimus9/systemstat"
	"flag"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/fsouza/go-dockerclient"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

type MonitorSample struct {
	Name     string  `json:"name"`
	CpuPct   float64 `json:"cpu_percent,omitempty"`
	MemUsed  uint64  `json:"mem_used"`
	MemTotal uint64  `json:"mem_total"`
}

func getSystemMonitorSample() MonitorSample {
	cpuSample := systemstat.GetCPUSample()
	idle := float64(cpuSample.Idle)
	total := float64(cpuSample.Total)

	memSample := systemstat.GetMemSample()

	var stat MonitorSample
	stat.Name = "system"
	stat.CpuPct = (total - idle) * 100.0 / total
	stat.MemTotal = memSample.MemTotal
	stat.MemUsed = memSample.MemUsed
	return stat
}

func getContainerMonitorSample(name string) (MonitorSample, error) {
	var stat MonitorSample
	dockerStats := make(chan *docker.Stats)
	errC := make(chan error, 1)
	go func() {
		errC <- client.Stats(docker.StatsOptions{name, dockerStats})
		close(errC)
	}()

	statsResult, ok := <-dockerStats
	if !ok {
		return stat, <-errC
	}
	stat.Name = name
	stat.MemTotal = statsResult.MemoryStats.Limit
	stat.MemUsed = statsResult.MemoryStats.Usage
	return stat, nil
}

var client *docker.Client

func main() {
	var port int
	flag.IntVar(&port, "port", 3001, "server port")
	flag.Parse()
	fmt.Println("Port: ", port)

	endpoint := "unix:///var/run/docker.sock"
	var err error
	client, err = docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.JSON(http.StatusOK, getSystemMonitorSample())
	})

	m.Get("/:name", func(args martini.Params, r render.Render) {
		name := args["name"]
		sample, err := getContainerMonitorSample(name)
		if err == nil {
			r.JSON(http.StatusOK, sample)
		} else {
			r.JSON(http.StatusBadRequest, err)
		}
	})

	http.Handle("/", m)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
