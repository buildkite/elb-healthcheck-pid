# ELB PID Checker

A web server for use with Amazon ELB's Health Checker to monitor the status of a process on the host machine.

### Usage

```bash
$ elb-healthcheck-pid --help
Usage: elb-healthcheck-pid [options] <pid-to-monitor>
      --debug=false: Enable debug mode
      --port=4567: Run the server on the specified port
      --version=false: Prints the current version
```

To monitor a PID, simply run the tool like so:

```bash
$ elb-healthcheck-pid 1234
```

You can then see the status of this PID by visiting: `http://0.0.0.0:4567/status`
