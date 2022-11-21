package cpu

import (
	"github.com/yuanqijing/chaos/pkg/perf/cpu/apis"
	"github.com/yuanqijing/chaos/pkg/util"
	"k8s.io/klog/v2"
	"strconv"
	"strings"
)

// command implements for cpu utilization related metrics

// UtilizationPerCore returns the CPU utilization on a given core.
// This function uses the command mpstat -o JSON -P [core] 1 1 to get the CPU utilization.
// mpstat reference: https://man7.org/linux/man-pages/man1/mpstat.1.html
func (c *cmd) UtilizationPerCore(core int) apis.CpuUtilization {
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
		Core:      core,
		Timestamp: ts,
		Usr:       cpu_load["usr"].(float64),
		Nice:      cpu_load["nice"].(float64),
		Sys:       cpu_load["sys"].(float64),
		Iowait:    cpu_load["iowait"].(float64),
		Irq:       cpu_load["irq"].(float64),
		Softirq:   cpu_load["softirq"].(float64),
		Steal:     cpu_load["steal"].(float64),
		Guest:     cpu_load["guest"].(float64),
		GuestNice: cpu_load["guest_nice"].(float64),
		Idle:      cpu_load["idle"].(float64),
	}
}

// UtilizationSystemWide returns total system wide utilization.
// This function uses command vmstat -n --unit M --wide --timestamp 1 2
// --procs-- -----------------------memory---------------------- ---swap-- -----io---- -system-- --------cpu-------- -----timestamp-----
//   r    b         swpd         free         buff        cache   si   so    bi    bo   in   cs  us  sy  id  wa  st                 UTC
//   1    0            0          981           55          767    0    0     5     1   17   27   0   0 100   0   0 2022-11-21 07:27:27
//   0    0            0          981           55          767    0    0     0     0   48   63   0   0 100   0   0 2022-11-21 07:27:28
func (c *cmd) UtilizationSystemWide() apis.Vmstat {
	cmd := "vmstat -n --unit M --wide --timestamp 1 2" + "| tail -n 1" + "| awk '{print $1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19}'"
	// exec command
	raw, err := util.ExecCmd(cmd)
	if err != nil {
		klog.Errorf("get system wide cpu utilization failed: %v", err)
		return apis.Vmstat{}
	}

	// parse raw
	ret := strings.Split(string(raw), " ")
	if len(ret) != 19 {
		klog.Errorf("parse vmstat failed: %v", err)
		return apis.Vmstat{}
	}

	f := func(i int, err error) int {
		if err != nil {
			klog.Errorf("parse vmstat failed: %v", err)
			return 0
		}
		return i
	}

	return apis.Vmstat{
		Timestamp: ret[17] + " " + ret[18],
		R:         f(strconv.Atoi(ret[0])),
		B:         f(strconv.Atoi(ret[1])),
		Swpd:      f(strconv.Atoi(ret[2])),
		Free:      f(strconv.Atoi(ret[3])),
		Buff:      f(strconv.Atoi(ret[4])),
		Cache:     f(strconv.Atoi(ret[5])),
		Si:        f(strconv.Atoi(ret[6])),
		So:        f(strconv.Atoi(ret[7])),
		Bi:        f(strconv.Atoi(ret[8])),
		Bo:        f(strconv.Atoi(ret[9])),
		In:        f(strconv.Atoi(ret[10])),
		Cs:        f(strconv.Atoi(ret[11])),
		Us:        f(strconv.Atoi(ret[12])),
		Sy:        f(strconv.Atoi(ret[13])),
		Id:        f(strconv.Atoi(ret[14])),
		Wa:        f(strconv.Atoi(ret[15])),
		St:        f(strconv.Atoi(ret[16])),
	}
}
