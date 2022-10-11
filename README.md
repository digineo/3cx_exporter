3CX Exporter
============

Exporter for the 3CX PBX.


## Usage Example

1. Download lastest release 
   sudo curl -o /usr/bin/3cx-exporter https://github.com/tarkmote-ou/3cx_exporter/releases/download/v2.0/3cx-explorer-amd64-linux-2.0


3. Create following systemd service unit at /etc/systemd/system/3cx_exporter.service:

```
[Unit]
Description=3CX Prometheus Exporter
After=network.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/bin/3cx_exporter  -log_level=INFO -listen=8080 -db_connection=root:1234@/sys?parseTime=true"
```
Where db connection is db mysql db connection string. Format:
username:password@protocol(address)/dbname?param=value
You must use parseTime= true attribute 

## License

MIT License, Copyright (c) 2018
[Digineo GmbH](https://www.digineo.de/)  
