package cpu

import (
	"github.com/yuanqijing/chaos/pkg/perf/cpu/apis"
	"github.com/yuanqijing/chaos/pkg/util"
	"k8s.io/klog/v2"
	"strconv"
)

// command implements for cpu utilization related metrics

// CmdCPUUtilizationOnCore returns the CPU utilization on a given core.
// This function uses the command mpstat -o JSON -P [core] 1 1 to get the CPU utilization.
// mpstat reference: https://linux
func CmdCPUUtilizationOnCore(core int) apis.CpuUtilization {
	cmd := "mpstat -o JSON -P " + strconv.Itoa(core) + " 1 1"
	// exec command
	raw, err := util.ExecCmd(cmd)
	if err != nil {
		klog.Errorf("get cpu utilization on core %d failed: %v", core, err)
		return apis.CpuUtilization{}
	}

	// parse json
	ret, err := util.ParseJson2Map(raw)
	if err != nil {
		klog.Errorf("parse json failed: %v", err)
		return apis.CpuUtilization{}
	}

	host := ret["sysstat"].(map[string]interface{})["hosts"].([]interface{})[0].(map[string]interface{})

	statistic := host["statistics"].([]interface{})[0].(map[string]interface{})

	ts := statistic["timestamp"].(string)

	cpu_load := statistic["cpu-load"].([]interface{})[0].(map[string]interface{})

	return apis.CpuUtilization{
		Core:       core,
		Timestamp:  ts,
		Usr:        cpu_load["usr"].(float64),
		Nice:       cpu_load["nice"].(float64),
		Sys:        cpu_load["sys"].(float64),
		Iowait:     cpu_load["iowait"].(float64),
		Irq:        cpu_load["irq"].(float64),
		Softirq:    cpu_load["softirq"].(float64),
		Steal:      cpu_load["steal"].(float64),
		Guest:      cpu_load["guest"].(float64),
		Guest_nice: cpu_load["guest_nice"].(float64),
		Idle:       cpu_load["idle"].(float64),
	}
}
