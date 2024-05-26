package utils

import (
	"bufio"
	"fmt"
	"github.com/jeyren95/network-journey/models"
	"net"
	"os/exec"
	"strings"
)

const DEFAULT_NUM_QUERIES = 1
const HOP_NUMBER_COL = 0
const HOSTNAME_COL = 1
const IP_COL = 2
const RETURN_TIME_COL = 3
const UNIT_OF_TIME_COL = 4

func TraceRoute(reqBody *models.IpHopsReqBody) ([]models.IpHop, error) {
	numQueries := fmt.Sprint("-q ", DEFAULT_NUM_QUERIES)
	maxHops := fmt.Sprint("-m ", reqBody.MaxHops)
	waitTime := fmt.Sprint("-w ", reqBody.WaitTime)

	cmd := exec.Command("traceroute", numQueries, maxHops, waitTime, reqBody.Hostname)
	// returns a pipe that implements the reader interface
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// NewScanner takes in a reader interface
	// the default split function used by scanner is scanLines, that returns each line of text stripped of any trailing EOL marker
	scanner := bufio.NewScanner(stdout)
	currRow := 0

	res := []models.IpHop{}

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		// skip if first row of the traceroute results or skip if query timed out
		if currRow == 0 || hasTimedOut(columns[HOSTNAME_COL]) {
			currRow += 1
			continue
		}

		ipHop := models.IpHop{}
		ipHop.Ip = columns[IP_COL][1 : len(columns[IP_COL])-1]
		ipHop.Hostname = columns[HOSTNAME_COL]
		ipHop.ReturnTime = columns[RETURN_TIME_COL] + columns[UNIT_OF_TIME_COL]
		ipHop.IsIpPrivate = isIpPrivate(ipHop.Ip)
		res = append(res, ipHop)

		currRow += 1
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

func hasTimedOut(col string) bool {
	return col == "*"
}

func isIpPrivate(ip string) bool {
	parsedIp := net.ParseIP(ip)
	return parsedIp.IsPrivate()
}
