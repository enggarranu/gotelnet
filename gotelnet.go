// Created By : enggar.ranu@gmail.com
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var ifailed int
	var isuccess int
	var failedIPs []string

	filename := flag.String("f", "list_ip.txt", "file/to/path which txt format that contanis list of ip")
	port := flag.String("p", "22", "telnet port")
	flag.Parse()

	ips := FileHandler(*filename)

	ifailed = 0
	isuccess = 0

	for _, ip := range ips {
		result, err := TelnetHandler(ip, *port)
		if result {
			log.Println("telnet : " + ip + " - port : " + *port + " - status : OK")
			isuccess++
		} else {
			log.Println("telnet : " + ip + " - port : " + *port + " - status : Failed - reaseon: " + err.Error())
			ifailed++
			failedIPs = append(failedIPs, ip)
		}
	}

	log.Printf("\n------------- RESULT telnet port %v -------------\nTelnet Success: %d Host | Telnet Failed: %d Host\n---------------------------------------------------", *port, isuccess, ifailed)
	var ipnya string
	for _, failedIP := range failedIPs {
		ipnya = ipnya + fmt.Sprintf("\n%s", failedIP)
	}
	log.Printf("List failed IP(s): %s", ipnya)
}

// TelnetHandler handle telnet by port and host ip
func TelnetHandler(host string, port string) (bool, error) {
	timeout := time.Second * 5
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		// fmt.Println("Connecting error:", err)
		return false, err
	}
	if conn != nil {
		defer conn.Close()
	}
	return true, nil
}

// FileHandler digunakan untuk mengimport file txt
func FileHandler(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	return txtlines
}
