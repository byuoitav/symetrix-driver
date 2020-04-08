package symetrix

import (
	"context"
	"fmt"
	"net"
	"bufio"
	"math"
)

const (
	maxVolumeLevel = 56000
)

// GetVolumeByBlock returns the volume [0, 100] of the given block.
func (d *DSP) GetVolumeByBlock(ctx context.Context, block string) (int, error) {

	c, err := net.Dial("tcp", d.Address+":48631")
	if err != nil {
		fmt.Println("unable to establish TCP client")
		return false, fmt.Errorf("unable to establish TCP client: %w", err)
	}
	fmt.Fprintf(c, "GS %v\n", block)
	result, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("unable to read response: %w", err)
	}
	if result > maxVolumeLevel {
		fmt.Println("fader is above max volume level")
		SetVolumeByBlock(ctx,block,100) error
		result = 56000
	}
	result = math.Round((result / maxVolumeLevel) * 100)

	return result, nil
}

// SetVolumeByBlock sets the volume on the given block. Volume must be in the range [0, 100].
func (d *DSP) SetVolumeByBlock(ctx context.Context, block string, volume int) error {
	if volume < 0 || volume > 100 {
		return fmt.Errorf("volume must be in range [0, 100]")
	}
	
	volume *= maxVolumeLevel
	if ()
	c, err := net.Dial("tcp", Address+":48631")
	if err != nil {
		fmt.Println("unable to establish TCP client")
		return fmt.Errorf("unable to establish TCP client: %w", err)
	}

	fmt.Fprintf(c, "CS %v %v\n", block, volume)
	result, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}
	if result != "ACK\n#0000"+block+"="+volume+"\n" {
		return fmt.Errorf("Unsuccessful")
	}
	else {
		return nil
	}
}
