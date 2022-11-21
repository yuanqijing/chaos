package apis

import "strconv"

// Vmstat is a metrics struct for vmstat output.
type Vmstat struct {
	Timestamp string

	// Procs
	// R: The number of runnable processes (running or waiting for run time).
	R int
	// B: The number of processes in uninterruptible sleep.
	B int

	// Memory
	// Swpd: The amount of swap memory used. Unit: MB
	Swpd int
	// Free: The amount of idle memory. Unit: MB
	Free int
	// Buff: The amount of memory used as buffers. Unit: MB
	Buff int
	// Cache: The amount of memory used as cache. Unit: MB
	Cache int
	// Inact: The amount of inactive memory. Unit: MB
	Inact int
	// Active: The amount of active memory. Unit: MB
	Active int

	// Swap
	// Si: Amount of memory swapped in from disk (/s).
	Si int
	// So: Amount of memory swapped to disk (/s).
	So int

	// IO
	// Bi: Blocks received from a block device (blocks/s).
	Bi int
	// Bo: Blocks sent to a block device (blocks/s).
	Bo int

	// System
	// In: The number of interrupts per second, including the clock.
	In int
	// Cs: The number of context switches per second.
	Cs int

	// CPU
	// Us: Time spent running non-kernel code.
	Us int
	// Sy: Time spent running kernel code.
	Sy int
	// Id: Time spent idle.
	Id int
	// Wa: Time spent waiting for I/O to complete.
	Wa int
	// St: Time stolen from a virtual machine.
	St int
}

// NewVmstat returns a new Vmstat struct.
func NewVmstat() *Vmstat {
	return &Vmstat{}
}

// String returns a string representation of the Vmstat struct.
func (v *Vmstat) String() string {
	return "Timestamp: " + v.Timestamp +
		"R: " + strconv.Itoa(v.R) +
		" B: " + strconv.Itoa(v.B) +
		" Swpd: " + strconv.Itoa(v.Swpd) + "MB" +
		" Free: " + strconv.Itoa(v.Free) + "MB" +
		" Buff: " + strconv.Itoa(v.Buff) + "MB" +
		" Cache: " + strconv.Itoa(v.Cache) + "MB" +
		" Inact: " + strconv.Itoa(v.Inact) + "MB" +
		" Active: " + strconv.Itoa(v.Active) + "MB" +
		" Si: " + strconv.Itoa(v.Si) +
		" So: " + strconv.Itoa(v.So) +
		" Bi: " + strconv.Itoa(v.Bi) +
		" Bo: " + strconv.Itoa(v.Bo) +
		" In: " + strconv.Itoa(v.In) +
		" Cs: " + strconv.Itoa(v.Cs) +
		" Us: " + strconv.Itoa(v.Us) +
		" Sy: " + strconv.Itoa(v.Sy) +
		" Id: " + strconv.Itoa(v.Id) +
		" Wa: " + strconv.Itoa(v.Wa) +
		" St: " + strconv.Itoa(v.St)
}
