package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kaedeek/port_scan-go/logger"
)

const (
	startPort      = 0
	endPort        = 65535
	timeout        = time.Second
	maxConcurrency = 1000
)

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

func ScanPort(ctx context.Context, ip string, port int, wg *sync.WaitGroup, sem chan struct{}, mu *sync.Mutex) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	case sem <- struct{}{}:
	}

	defer func() { <-sem }()

	address := net.JoinHostPort(ip, fmt.Sprintf("%d", port))
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err == nil {
		conn.Close()
		service := getServiceName(port)
		mu.Lock()
		logger.LogV2(fmt.Sprintf("%d | %s Open", port, service))
		mu.Unlock()
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	logger.LogV1("Please enter an IP address to scan: ")
	ip, _ := reader.ReadString('\n')
	ip = strings.TrimSpace(ip)

	logger.LogV1(fmt.Sprintf("Port scan started | %s (Press 'e' to stop)", ip) + "\n\n")

	var wg sync.WaitGroup
	var mu sync.Mutex
	sem := make(chan struct{}, maxConcurrency)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		keyReader := bufio.NewReader(os.Stdin)
		for {
			char, _ := keyReader.ReadByte()
			if char == 'e' || char == 'E' {
				cancel()
				logger.LogV1("\nScan stopped :(")
				return
			}
		}
	}()

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go ScanPort(ctx, ip, port, &wg, sem, &mu)
	}

	wg.Wait()
	logger.LogV1("Scan completed :)\n")

	logger.LogV1("Press Enter to exit...")
	reader.ReadBytes('\n')
}
