package cpu

type Config struct {
	// Features enable
	// PerCPU enables per-cpu metrics to be collected.
	PerCPU bool
	//
	SystemWide bool
}

func NewDefaultConfig() *Config {
	return &Config{}
}
