package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	versalog "github.com/VersaLog/VersaLog.go/VersaLog"
)

const (
	startPort      = 0
	endPort        = 65535
	timeout        = time.Second
	maxConcurrency = 1000
)

type ScanResult struct {
	Port    int
	IsOpen  bool
	Service string
}

func ScanPort(ip string, port int) ScanResult {
	address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	result := ScanResult{Port: port}

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		result.IsOpen = false
		return result
	}
	defer conn.Close()

	result.IsOpen = true
	result.Service = getServiceName(port)
	return result
}

func getServiceName(port int) string {
	services := map[int]string{
		0:     "Reserved",
		1:     "TCPMUX",
		7:     "Echo",
		9:     "Discard",
		13:    "Daytime",
		17:    "Quote of the Day",
		19:    "Character Generator",
		20:    "FTP-DATA",
		21:    "FTP",
		22:    "SSH",
		23:    "Telnet",
		25:    "SMTP",
		37:    "Time",
		42:    "Nameserver",
		43:    "WHOIS",
		49:    "TACACS",
		53:    "DNS",
		67:    "DHCP Server",
		68:    "DHCP Client",
		69:    "TFTP",
		70:    "Gopher",
		79:    "Finger",
		80:    "HTTP",
		88:    "Kerberos",
		102:   "ISO-TSAP",
		110:   "POP3",
		111:   "RPC",
		113:   "Ident",
		119:   "NNTP",
		123:   "NTP",
		135:   "MSRPC",
		137:   "NetBIOS Name Service",
		138:   "NetBIOS Datagram",
		139:   "NetBIOS Session",
		143:   "IMAP",
		161:   "SNMP",
		162:   "SNMP-TRAP",
		179:   "BGP",
		194:   "IRC",
		220:   "IMAP3",
		389:   "LDAP",
		443:   "HTTPS",
		445:   "Microsoft-DS",
		465:   "SMTPS",
		514:   "Syslog",
		515:   "LPD",
		587:   "Submission",
		636:   "LDAPS",
		993:   "IMAPS",
		995:   "POP3S",
		1080:  "SOCKS",
		1433:  "MSSQL",
		1521:  "Oracle",
		1723:  "PPTP",
		3306:  "MySQL",
		3389:  "RDP",
		5432:  "PostgreSQL",
		5900:  "VNC",
		6379:  "Redis",
		8080:  "HTTP-Proxy",
		8443:  "HTTPS-Alt",
		27017: "MongoDB",
		27018: "MongoDB-Shard",
		27019: "MongoDB-Router",
		50070: "Hadoop-NameNode",
		50075: "Hadoop-DataNode",
		50090: "Hadoop-SecondaryNameNode",
	}
	if service, ok := services[port]; ok {
		return service
	}
	return "Unknown"
}

func main() {
	logger := versalog.NewVersaLog("simple2", false, true, "Scan", false, true, true, nil, false, false)

	ip := flag.String("ip", "", "Target IP address to scan")
	flag.Parse()

	if *ip == "" {
		fmt.Println("Error: IP address is required")
		fmt.Println("Usage: go run main.go -ip <target_ip>")
		os.Exit(1)
	}

	logger.Info(fmt.Sprintf("Starting port scan on %s\n", *ip))

	// Create channels for results and synchronization
	results := make(chan ScanResult, endPort-startPort+1)
	var wg sync.WaitGroup

	// Launch port scanning goroutines
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			result := ScanPort(*ip, p)
			if result.IsOpen {
				results <- result
			}
		}(port)
	}

	// Wait for all scans to complete in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process and log results
	for result := range results {
		logger.Info(fmt.Sprintf("Port %d is open (%s)\n", result.Port, result.Service))
	}

	logger.Info("Scan completed")
}
