package configs

type Config struct {
	FileName      string  // File name for input data
	ChannelSize   int     // Channel size for calculating distances in concurrently
	CostPerDriver float64 // Minute cost for a driver; default 500minuts/driver
}
