package apis

import "strconv"

type CpuUtilization struct {
	Core      int
	Timestamp string
	Usr       float64
	Nice      float64
	Sys       float64
	Iowait    float64
	Irq       float64
	Softirq   float64
	Steal     float64
	Guest     float64
	GuestNice float64
	Idle      float64
}

func NewCpuUtilization() *CpuUtilization {
	return &CpuUtilization{}
}

func (c *CpuUtilization) String() string {
	return " Timestamp: " + c.Timestamp +
		"Core: " + strconv.Itoa(c.Core) +
		" Usr: " + strconv.FormatFloat(c.Usr, 'f', 2, 64) +
		" Nice: " + strconv.FormatFloat(c.Nice, 'f', 2, 64) +
		" Sys: " + strconv.FormatFloat(c.Sys, 'f', 2, 64) +
		" Iowait: " + strconv.FormatFloat(c.Iowait, 'f', 2, 64) +
		" Irq: " + strconv.FormatFloat(c.Irq, 'f', 2, 64) +
		" Softirq: " + strconv.FormatFloat(c.Softirq, 'f', 2, 64) +
		" Steal: " + strconv.FormatFloat(c.Steal, 'f', 2, 64) +
		" Guest: " + strconv.FormatFloat(c.Guest, 'f', 2, 64) +
		" GuestNice: " + strconv.FormatFloat(c.GuestNice, 'f', 2, 64) +
		" Idle: " + strconv.FormatFloat(c.Idle, 'f', 2, 64)
}
