package symetrix

import (
	"bufio"
	"context"
	"fmt"
	"math"
	"net"
	"strconv"
)

const (
	maxVolumeLevel = 56000
)

// GetVolumeByBlock returns the volume [0, 100] of the given block.
func (d *SymetrixDSP) GetVolumeByBlock(ctx context.Context, block string) (int, error) {

	c, err := net.Dial("tcp", d.Address+":48631")
	if err != nil {
		fmt.Println("unable to establish TCP client")
		return -1, fmt.Errorf("unable to establish TCP client: %w", err)
	}
	fmt.Fprintf(c, "GS %v\n", block)
	strResult, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return -1, fmt.Errorf("unable to read response: %w", err)
	}
	result, err := strconv.Atoi(strResult)
	if err != nil {
		return -1, fmt.Errorf("unable to convert volume string to int: %w", err)
	}
	if result > maxVolumeLevel {
		fmt.Println("fader is above max volume level")
		err = d.SetVolumeByBlock(ctx, block, 100)
		if err != nil {
			return -1, fmt.Errorf("unable to set volume: %w", err)
		}
		result = 56000
	}
	result = int(math.Round((float64(result) / maxVolumeLevel) * 100))

	return result, nil
}

// SetVolumeByBlock sets the volume on the given block. Volume must be in the range [0, 100].
func (d *SymetrixDSP) SetVolumeByBlock(ctx context.Context, block string, volume int) error {
	if volume < 0 || volume > 100 {
		return fmt.Errorf("volume must be in range [0, 100]")
	}

	volume *= maxVolumeLevel
	c, err := net.Dial("tcp", d.Address+":48631")
	if err != nil {
		fmt.Println("unable to establish TCP client")
		return fmt.Errorf("unable to establish TCP client: %w", err)
	}

	fmt.Fprintf(c, "CS %v %v\n", block, volume)
	result, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}
	if result != fmt.Sprintf("ACK\n#0000%v=%v\n", block, volume) {
		return fmt.Errorf("Unsuccessful")
	} else {
		return nil
	}
}
