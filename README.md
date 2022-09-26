3CX Exporter
============

Prometheus exporter for the 3CX PBX.

## Build

Move to the cloned git directory and do `go build`

## Configuration

Configuration is done with a config file in json format.
Example:
```
{
  "Hostname": "YOUR-PBX-FQDN:PORT",
  "Username": "ADMIN-USERNAME",
  "Password": "ADMIN-PASSWORD"
}
```

## Usage Example

1. Download lastest release curl -o /usr/bin/ https://github.com/tarkmote-ou/3cx_exporter/releases/download/v1.0/3cx-explorer-amd64-linux-1.0 

2. Create config file at `/etc/3cx_exporter/config.json`.

3. Create following systemd service unit at /etc/systemd/system/3cx_exporter.service:

```
[Unit]
Description=3CX Prometheus Exporter
After=network.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/bin/3cx_exporter -config /etc/3cx_exporter/config.json -log_level=INFO -listen=8080
```

## License

MIT License, Copyright (c) 2018
[Digineo GmbH](https://www.digineo.de/)  
