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

Move binary to `/usr/bin/` and create config file at `/etc/3cx_exporter/config.json`.
Create following systemd service unit at /etc/systemd/system/3cx_exporter.service:

```
[Unit]
Description=3CX Prometheus Exporter
After=network.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/bin/3cx_exporter -config /etc/3cx_exporter/config.json
```

## License

MIT License, Copyright (c) 2018
[Digineo GmbH](https://www.digineo.de/)  
