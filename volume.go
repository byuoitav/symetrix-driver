package symetrix

import (
	"context"
	"fmt"
	"net"
)

const (
	maxVolumeLevel = 56000
)

// GetVolumeByBlock returns the volume [0, 100] of the given block.
func (d *DSP) GetVolumeByBlock(ctx context.Context, block string) (int, error) {

	s, err := net.ResolveUDPAddr("udp4", address+":48631")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		return -1, fmt.Errorf("unable to establish UDP client: %w", err)
	}

    defer c.Close()

	text := fmt.Sprintf("GS %v\r\n", block)
	data := []byte(text)
	_, err = c.Write(data)

	if err != nil {
		fmt.Println(err)
		return -1, fmt.Errorf("unable to write to client: %w", err)
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return -1, fmt.Errorf("unable to read response: %w", err)
	}

	val := string(buffer[0:n])
	result, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64)
	
	if result > maxVolumeLevel {
		result = 56000
	}

	volume := int(float64(result*100.0) / maxVolumeLevel)

	return volume, nil
}

// SetVolumeByBlock sets the volume on the given block. Volume must be in the range [0, 100].
func (d *DSP) SetVolumeByBlock(ctx context.Context, block string, volume int) error {
	if volume < 0 || volume > 100 {
		return fmt.Errorf("volume must be in range [0, 100]")
	}
	
	s, err := net.ResolveUDPAddr("udp4", address+":48631")
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		return fmt.Errorf("unable to establish UDP client: %w", err)
	}

    defer c.Close()
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

    if (val == "ACK\r") {
        fmt.Println("GOOD\n")
        return nil
    }

	return fmt.Errorf("Unsuccessful Volume Change")
}
