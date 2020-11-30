# Renson Endura Delta Monitor
Small utility to read the configuration from a [Renson Endura Delta](https://www.renson.eu/en-gb/producten-zoeken/ventilatie/mechanische-ventilatie/units/endura-delta) ventilation unit.

When the filter needs cleaning a mail is sent via [mailgun](https://www.mailgun.com/).

# Installation
## Raspberry Pi OS
### Install Go
1. `wget https://golang.org/dl/go1.15.5.linux-armv6l.tar.gz` (latest version can be found at the [golang site](https://golang.org/dl/))
2. `sudo tar -xvf go1.15.5.linux-armv6l.tar.gz`
3. `sudo mv go /usr/local`

### Clone repository
1. Clone the repository to your Raspberry Pi installation

### Install Go packages
1. Run `go get gopkg.in/mailgun/mailgun-go.v1`
2. Run `go get github.com/robfig/cron`

### Build
1. Run `go build monitor.go`

### Install as a service
1. Run `vi /lib/systemd/system/rensonenduradelta.service`
2. Add the following section to the file:

```
[Unit]
Description=Renson Endura Delta monitor

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/renson-endura-delta-monitor/monitor

[Install]
WantedBy=multi-user.target
```

3. Run `service rensonenduradelta start`