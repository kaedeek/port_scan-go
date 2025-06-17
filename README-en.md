# Port Scanner

A simple concurrent port scanner that performs TCP port scanning from port 0 to 65535 on a specified IP address.

[Japanese](README.md)

## Features

- **TCP port scanning for specified IP addresses**
- **Fast scanning execution with concurrent processing**
- **Automatic identification of well-known services (FTP, SSH, HTTP, etc.)**

> [!CAUTION]
> Prohibited Actions:
>
> - This tool is created for educational purposes
> - Obtain permission from the system owner before performing port scanning
> - Improper use may lead to legal issues

## Scan Results

Scan results are output in the following format

```
[+] Please enter an IP address to scan:
000.000.0.0
[+] Port scan started | 000.000.0.0 (Press 'e' to stop)


[-] 80 | HTTP Open
[-] 53 | DNS Open
[-] 27998 | Unknown Open
[-] 37444 | Unknown Open
[-] 37443 | Unknown Open
[-] 49652 | Unknown Open
[-] 49653 | Unknown Open
[+] Scan completed :)


[+] Press Enter to exit...
```

## EXE FILE LIST

[EXE FILE](exe/portscanner.exe)

## Message List

| Message        | Description     |
| -------------- | --------------- |
| `Unknown Open` | No service name |
