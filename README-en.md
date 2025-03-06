# Port Scanner

A simple concurrent port scanner that performs TCP port scanning from port 1 to 1024 on a specified IP address.

[Japan](README-ja.md)

## Features

- TCP port scanning for specified IP addresses
- High-speed scanning with concurrent processing
- Automatic identification of well-known services (FTP, SSH, HTTP, etc.)
- Scan results output to log file

## Installation

```bash
git clone https://github.com/kaedeek/port_scan-go.git
cd port_scan-go
```

## Usage

```bash
go run main.go -ip <target_ip>
```

### Example

```bash
go run main.go -ip 000.000.0.0
```

Scan results are saved in the `scan.log` file.

## Scan Results

The scan results are output to the log file in the following format:

```
2024/03/XX HH:MM:SS Starting port scan on 000.000.0.0 (ports 1-1024)
2024/03/XX HH:MM:SS Port 22 is open (SSH)
2024/03/XX HH:MM:SS Port 80 is open (HTTP)
2024/03/XX HH:MM:SS Port 443 is open (HTTPS)
2024/03/XX HH:MM:SS Scan completed
```

## Important Notes

- This tool is created for educational purposes
- Please obtain permission from the system owner before performing port scans
- Improper use may result in legal issues

## License

MIT License with additional restrictions.
Copyright (c) 2025 kaedeek

See the LICENSE file for details.
