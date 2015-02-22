package main

import (
	"errors"
	"flag"
	"fmt"

	helper "github.com/dselans/goroq/helper"
)

const (
	DefaultQuietMode  bool   = false
	DefaultMonitorDir string = "./"
	DefaultOutputFile string = "goroq.log"
	DefaultConfigFile string = "goroq.json"
)

type CliOptions struct {
	QuietMode   bool
	VersionFlag bool
	MonitorDir  string
	OutputFile  string
	ConfigFile  string
}

func displayUsage() {
	helper.CustomExit(fmt.Sprintf("Usage: ./goreq [-h|-v|] [-d directory] [-o output_file] [-c config_file]\n\n"+
		"-h               : Display this help message\n"+
		"-v               : Display goreq version\n"+
		"-q               : Be quiet (and drive) (default: %v)\n"+
		"-d directory     : Set monitor/test dir (default: %v)\n"+
		"-o output_file   : Set output log file (default: %v)\n"+
		"-c config_file   : Set alternate config (default: %v)",
		DefaultQuietMode, DefaultMonitorDir, DefaultOutputFile,
		DefaultConfigFile), 0)
}

// Fetch and validate cli args
func handleCliArgs() *CliOptions {
	opts := &CliOptions{}

	flag.Usage = displayUsage
	flag.BoolVar(&opts.QuietMode, "q", DefaultQuietMode, "Turn off printing to stdout")
	flag.BoolVar(&opts.VersionFlag, "v", false, "Display goreq version")
	flag.StringVar(&opts.MonitorDir, "d", DefaultMonitorDir, "Set monitor and test dir")
	flag.StringVar(&opts.OutputFile, "o", DefaultOutputFile, "Set output log file")
	flag.StringVar(&opts.ConfigFile, "c", DefaultConfigFile, "Set custom config file")

	flag.Parse()

	if opts.VersionFlag {
		helper.CustomExit(fmt.Sprintf("goreq v%v", VERSION), 0)
	}

	if err := validateCliArgs(opts); err != nil {
		helper.CustomExit(err.Error(), 1)
	}

	return opts
}

func validateCliArgs(opts *CliOptions) error {
	// Check whether MonitorDir is a dir
	if !helper.IsDir(opts.MonitorDir) {
		return errors.New(fmt.Sprintf("Monitor directory '%v' does not appear to be a valid directory", opts.MonitorDir))
	}

	// Check that output file is writable
	if !helper.IsWritable(opts.OutputFile) {
		return errors.New(fmt.Sprintf("Output file '%v' is not writable", opts.OutputFile))
	}

	// Some additional config checks should go in here

	return nil
}
