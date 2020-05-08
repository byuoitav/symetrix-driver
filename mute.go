package symetrix

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	MUTE_ENABLE_VAL = 65535
)

// GetMutedByBlock returns true if the given block is muted.
func (d *DSP) GetMutedByBlock(ctx context.Context, block string) (bool, error) {
	s, err := net.ResolveUDPAddr("udp4", d.Address+":48631")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		return false, fmt.Errorf("unable to establish UDP client: %w", err)
	}

	defer c.Close()
	c.SetReadDeadline(time.Now().Add(5 * time.Second))

	text := fmt.Sprintf("GS %v\r\n", block)
	data := []byte(text)
	_, err = c.Write(data)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("unable to write to client: %w", err)
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("unable to read response: %w", err)
	}

	val := string(buffer[0:n])
	result, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64)

	if result == MUTE_ENABLE_VAL {
		return true, nil
	}

	return false, nil
}

// SetMutedByBlock sets the mute state on the given block
func (d *DSP) SetMutedByBlock(ctx context.Context, block string, muted bool) error {

	s, err := net.ResolveUDPAddr("udp4", d.Address+":48631")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		return fmt.Errorf("unable to establish UDP client: %w", err)
	}
	defer c.Close()
	c.SetReadDeadline(time.Now().Add(5 * time.Second))

	muteVal := 0
	if muted {
		muteVal = MUTE_ENABLE_VAL
	}
	text := fmt.Sprintf("CS %v %v\r\n", block, muteVal)

	data := []byte(text)
	_, err = c.Write(data)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("unable to write to client: %w", err)
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("unable to read response: %w", err)
	}

	val := fmt.Sprintf("%s", string(buffer[0:n]))
	if muted {
		if val == "ACK\r" {
			return nil
		}
		return fmt.Errorf("Unsuccessful")
	} else {
		if val == "ACK\r" {
			return nil
		}
		return fmt.Errorf("Unsuccessful")
	}
}
