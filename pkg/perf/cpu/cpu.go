package cpu

import (
	"github.com/yuanqijing/chaos/pkg/util"
	"strconv"
	"strings"
)

type CPU struct {
	cores []int
}

func NewCPU(config *Config) *CPU {
	//  common
	cores, err := GetCores()
	if err != nil {
		panic(err)
	}

	// settings considering config

	return &CPU{
		cores: cores,
	}
}

// GetUtilOnCore returns the CPU utilization on a given core.
func (c *CPU) GetUtilOnCore(core int) (float64, error) {
	// read /sys/devices/system/cpu/cpu0/cpufreq/stats/time_in_state
	// TODO: implement
	_, err := util.ReadFile("/sys/devices/system/cpu/cpu" + strconv.Itoa(core) + "/cpufreq/stats/time_in_state")
	if err != nil {
		return 0, err
	}
	return 0, nil
}

// GetCores returns the number of cores on the system.
func GetCores() ([]int, error) {
	// read /sys/devices/system/cpu/online
	raw, err := util.ReadFile("/sys/devices/system/cpu/online")
	if err != nil {
		return nil, err
	}

	// parse the output
	// the output is as follows: 0-3
	b, e := func(raw []byte) (int, int) {
		// get rid of blank spaces
		str := strings.Split(strings.TrimSpace(string(raw)), "-")
		// parse to int, we trust this is safe
		a, _ := strconv.Atoi(str[0])
		b, _ := strconv.Atoi(str[1])
		return a, b
	}(raw)

	var cores []int
	for i := b; i <= e; i++ {
		cores = append(cores, i)
	}
	return cores, nil
}
