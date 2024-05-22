package utils

import (
	"bufio"
	"fmt"
	"github.com/jeyren95/network-journey/models"
	"os/exec"
	"strconv"
	"strings"
)

func TraceRoute(reqBody *models.IpHopsReqBody) ([]models.IpHop, error) {
	numQueries := fmt.Sprint("-q ", reqBody.NumQueries)
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
		// skip if first row of the traceroute results
		if currRow == 0 {
			currRow += 1
			continue
		}

		line := scanner.Text()
		columns := strings.Fields(line)

		ipHop := models.IpHop{}

		for index, col := range columns {
			if isHopIndex(index) || hasTimedOut(col) {
				continue
			}

			if isIpColumn(col) {
				ipHop.Ip = col[1 : len(col)-1]
			} else if ipHop.Hostname == "" {
				ipHop.Hostname = col
			} else {
				if isUnitOfTimeColumn(col) {
					ipHop.ReturnTimes[len(ipHop.ReturnTimes)-1] += col
				} else {
					ipHop.ReturnTimes = append(ipHop.ReturnTimes, col)
				}
			}
		}

		res = append(res, ipHop)
		currRow += 1
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return res, nil
}

func isHopIndex(columnNumber int) bool {
	return columnNumber == 0
}

func hasTimedOut(col string) bool {
	return col == "*"
}

func isIpColumn(col string) bool {
	return strings.ContainsRune(col, '(') && strings.ContainsRune(col, ')')
}

func isUnitOfTimeColumn(col string) bool {
	_, err := strconv.ParseFloat(col, 32)
	return err != nil
}
