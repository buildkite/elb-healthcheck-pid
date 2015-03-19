# ELB PID Checker

A web server for use with Amazon ELB's Health Checker to monitor the status of a process on the host machine.

## Usage

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

## Installation

Downloading the latest release for your OS here: https://github.com/buildkite/elb-healthcheck-pid/releases

## Licence

```
The MIT License (MIT)

Copyright (c) 2015 Buildkite Pty. Ltd.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
