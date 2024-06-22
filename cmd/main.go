package main

import (
	"flag"

	"github.com/sdblg/vrp/pkg/configs"
	"github.com/sdblg/vrp/pkg/services"
)

var Version string

func main() {
	// printCurrentVersion()
	cfg := configs.Config{}

	flag.StringVar(&cfg.FileName, "file", "./test-data/TrainingProblems/problem1.txt", "file-to-path/file-name.txt")
	flag.Float64Var(&cfg.CostPerDriver, "cost", 500.00, "Minute cost for a driver; default 500minuts/driver")
	flag.IntVar(&cfg.ChannelSize, "size", 5, "Channel size for calculating distances in concurrently")
	flag.Parse()

	// begin := time.Now()

	// fmt.Println("Current version is", Version)
	// fmt.Println("Reading a file from", cfg.FileName)

	service, _ := services.New(cfg)
	service.Do()

	// fmt.Println("Consumed total time:", time.Since(begin))
}

func printCurrentVersion() {
	if Version == "" {
		Version = "local"
	}
}
