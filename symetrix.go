package symetrix

import (
	"fmt"
	"strings"
)

// DSP represents a DSP
type DSP struct {
	Address string
}

// This is where things like a status struct could go etc.

func parseBlock(block string) (string, string, error) {
	parsedBlock := strings.Split(block, "|")
	if len(parsedBlock) == 1 {
		return "", "", fmt.Errorf("block is not in the correct format. Expecting gain#|mute#, recieved: %s", block)
	}

	gainBlock := parsedBlock[0]
	muteBlock := parsedBlock[1]

	return gainBlock, muteBlock, nil
}
