package symetrix

import (
	"bufio"
	"context"
	"fmt"
	"net"
)

// GetMutedByBlock returns true if the given block is muted.
func (d *SymetrixDSP) GetMutedByBlock(ctx context.Context, block string) (bool, error) {

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
	if result == "0\n" {
		return false, fmt.Errorf("TODO")
	}
	return true, nil
}

// SetMutedByBlock sets the mute state on the given block
func (d *SymetrixDSP) SetMutedByBlock(ctx context.Context, block string, muted bool) error {
	if muted {
		return nil
	}

	c, err := net.Dial("tcp", Address+":48631")
	if err != nil {
		fmt.Println("unable to establish TCP client")
		return
	}

}
