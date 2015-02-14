package main

import (
	"errors"
	"flag"
	"fmt"
)

const (
	DefaultMonitorDir string = "./"
	DefaultOutputFile string = "./goreq.log"
	DefaultConfigFile string = "~/.goreqrc"
)

type CliOptions struct {
	VersionFlag bool
	MonitorDir  string
	OutputFile  string
	ConfigFile  string
}

func displayUsage() {
	HelperCustomExit(fmt.Sprintf("Usage: ./goreq [-h|-v|] [-d directory] [-o output_file] [-c config_file]\n\n"+
		"-h               : Display this help message\n"+
		"-v               : Display goreq version\n"+
		"-d directory     : Set monitor/test dir (default: %v)\n"+
		"-o output_file   : Set output log file (default: %v)\n"+
		"-c config_file   : Set alternate config (default: %v)",
		DefaultMonitorDir, DefaultOutputFile, DefaultConfigFile), 0)
}

// Fetch and validate cli args
func handleCliArgs() *CliOptions {
	opts := &CliOptions{}

	flag.Usage = displayUsage
	flag.BoolVar(&opts.VersionFlag, "v", false, "Display goreq version")
	flag.StringVar(&opts.MonitorDir, "d", DefaultMonitorDir, "Set monitor and test dir")
	flag.StringVar(&opts.OutputFile, "o", DefaultOutputFile, "Set output log file")
	flag.StringVar(&opts.ConfigFile, "c", DefaultConfigFile, "Set custom config file")

	flag.Parse()

	if opts.VersionFlag {
		HelperCustomExit(fmt.Sprintf("goreq v%v", VERSION), 0)
	}

	if err := validateCliArgs(opts); err != nil {
		HelperCustomExit(err.Error(), 1)
	}

	return opts
}

func validateCliArgs(opts *CliOptions) error {
	// Check whether MonitorDir is a dir
	if !HelperIsDir(opts.MonitorDir) {
		return errors.New("Monitor directory does not appear to be a valid directory")
	}

	// Check that output file is writable
	if !HelperIsWritable(opts.OutputFile) {
		return errors.New("Output file is not writable")
	}

	// Some additional config checks should go in here

	return nil
}
