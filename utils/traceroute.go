package utils

import (
	"fmt"
	"os/exec"
	"bufio"
	"github.com/jeyren95/network-journey/models"
)

const DEFAULT_NUM_QUERIES = 1

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
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
	}


	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return nil, nil
}





