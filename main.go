package main

import (
	"fmt"
	"github.com/buildkite/agent/buildkite/logger"
	flag "github.com/ogier/pflag"
	"io"
	"net/http"
	"os"
	"strconv"
	"syscall"
)

// The name of this tool
var NAME = "elb-healthcheck-pid"

// The current version of this tool
var VERSION = "0.1"

// The PID that we'll be monitoring
var pid int

// Handles requests to /
func rootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("Hello! If you'd like to know the status of %d, navigate to /status\n", pid))
}

// Handles requests to /status, will return a 200 or 503 depending on whether
// or not the process is running.
func statusHandler(w http.ResponseWriter, req *http.Request) {
	if checkPid(pid) {
		io.WriteString(w, fmt.Sprintf("PID %d: ✔\n", pid))
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(w, fmt.Sprintf("PID %d: ✘\n", pid))
	}
}

// Starts a web server
func startWebServer(port int) {
	// Register request handlers
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/status", http.HandlerFunc(statusHandler))

	// Start listing on a given port
	logger.Info("Server started, CTRL+C to stop")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	if err != nil {
		logger.Fatal("Failed to start web server: %s", err.Error())
	}
}

// Will return true if the process with PID exists.
func checkPid(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		logger.Debug("Failed to find process with PID %d (%s)", pid, err.Error())
		return false
	}

	err = process.Signal(os.Signal(syscall.Signal(0)))
	if err != nil {
		logger.Debug("Failed to signal process with PID %d (%s)", pid, err.Error())
		return false
	}

	return true
}

func main() {
	// Customize the usage text
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <pid-to-monitor>\n", NAME)
		flag.PrintDefaults()
	}

	// Setup the command line flags
	portPointer := flag.Int("port", 4567, "Run the server on the specified port")
	debugPointer := flag.Bool("debug", false, "Enable debug mode")
	versionPointer := flag.Bool("version", false, "Prints the current version")
	flag.Parse()

	if *versionPointer {
		fmt.Printf("%s %s\n", NAME, VERSION)
		os.Exit(0)
	}

	// Turn on debugging if --debug was passed
	if *debugPointer {
		logger.SetLevel(logger.DEBUG)
	}

	// Make sure they've passed the PID to monitor
	if len(flag.Args()) != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Convert the PID to an integer
	i, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		logger.Fatal("Failed to convert %s to an integer (%s)", flag.Arg(0), err)
	}

	// Set the global PID to monitor
	pid = i

	logger.Info("Monitoring PID: %d", i)
	logger.Info("Visit http://0.0.0.0:%d/status to check the status of the process.", *portPointer)

	startWebServer(*portPointer)
}
