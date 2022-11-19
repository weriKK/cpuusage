package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cpuLoad = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "myapp_cpu_usage_percent_sec",
		Help: "Measured CPU usage percentage per second",
	})
)

func measureCpuStats() {

	measurementPeriod := 5 * time.Second

	for {
		cgs0 := GetCgroupCpuStats()
		time.Sleep(measurementPeriod)
		cgs1 := GetCgroupCpuStats()

		diff := float64(cgs1.Usage - cgs0.Usage)
		normalizedQuota := float64(cgs0.BwQuota) * (float64(measurementPeriod.Microseconds()) / float64(cgs0.BwPeriod))

		usagePct := (diff / normalizedQuota) * 100

		fmt.Printf("%+v  \t%f -> %f\n", cgs1, diff, usagePct)

		cpuLoad.Set(usagePct)
	}
}

/*

/sys/fs/cgroup # cat cpu.stat
usage_usec 118070
user_usec 36329
system_usec 81741
nr_periods 141
nr_throttled 0
throttled_usec 0
nr_bursts 0
burst_usec 0

*/

// CpuStats contains cgroup cpu controller settings and statistics. All values are in microserconds (us).
type CpuStats struct {
	BwQuota  int
	BwPeriod int
	Usage    int
	User     int
	System   int
}

func GetCgroupCpuStats() *CpuStats {
	quota, period := readCpuBandwidth()
	usage, user, system := readCpuStats()
	return &CpuStats{
		BwQuota:  quota,
		BwPeriod: period,
		Usage:    usage,
		User:     user,
		System:   system,
	}
}

func readCpuBandwidth() (int, int) {

	f, err := os.Open("/sys/fs/cgroup/cpu.max")
	if err != nil {
		panic(err)
	}

	bufReader := bufio.NewReaderSize(f, 64)
	line, err := bufReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	parts := strings.Fields(line)
	if len(parts) != 2 {
		panic("invalid cpu.max format")
	}

	quota, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	period, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return quota, period
}

func readCpuStats() (int, int, int) {

	f, err := os.Open("/sys/fs/cgroup/cpu.stat")
	if err != nil {
		panic(err)
	}

	stats := map[string]int{}

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		parts := strings.Fields(fs.Text())
		if len(parts) != 2 {
			panic("invalid cpu.stats format")
		}

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		stats[parts[0]] = val
	}

	return stats["usage_usec"], stats["user_usec"], stats["system_usec"]
}
