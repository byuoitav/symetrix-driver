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
	maxVolumeLevel = 56000
)

// GetVolumeByBlock returns the volume [0, 100] of the given block.
func (d *DSP) GetVolumes(ctx context.Context, blocks []string) (map[string]int, error) {
	toReturn := make(map[string]int)

	for _, block := range blocks {
		s, err := net.ResolveUDPAddr("udp4", d.Address+":48631")
		c, err := net.DialUDP("udp4", nil, s)
		if err != nil {
			return map[string]int{block: -1}, fmt.Errorf("unable to establish UDP client: %w", err)
		}

		defer c.Close()
		c.SetReadDeadline(time.Now().Add(5 * time.Second))

		text := fmt.Sprintf("GS %v\r\n", block)
		data := []byte(text)
		_, err = c.Write(data)

		if err != nil {
			fmt.Println(err)
			return map[string]int{block: -1}, fmt.Errorf("unable to write to client: %w", err)
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return map[string]int{block: -1}, fmt.Errorf("unable to read response: %w", err)
		}

		val := string(buffer[0:n])
		result, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64)

		if result > maxVolumeLevel {
			result = 56000
		}

		volume := int(float64(result*100.0) / maxVolumeLevel)
		toReturn[block] = volume
	}

	return toReturn, nil
}

// SetVolumeByBlock sets the volume on the given block. Volume must be in the range [0, 100].
func (d *DSP) SetVolume(ctx context.Context, block string, volume int) error {
	if volume < 0 || volume > 100 {
		return fmt.Errorf("volume must be in range [0, 100]")
	}

	s, err := net.ResolveUDPAddr("udp4", d.Address+":48631")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		return fmt.Errorf("unable to establish UDP client: %w", err)
	}

	defer c.Close()
	c.SetReadDeadline(time.Now().Add(5 * time.Second))
	volume = volume * (maxVolumeLevel / 100)
	text := fmt.Sprintf("CS %v %v\r\n", block, volume)
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

	if val == "ACK\r" {
		return nil
	}

	return fmt.Errorf("Unsuccessful Volume Change")
}
