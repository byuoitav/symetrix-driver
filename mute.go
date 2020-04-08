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

	c, err := net.Dial("tcp", Address+":48631")
	if err != nil {
		fmt.Println("unable to establish TCP client")
		return fmt.Errorf("unable to establish TCP client: %w", err)
	}
	if muted {
		fmt.Fprintf(c, "CS %v 65535\n", block)
	}
	else {
		fmt.Fprintf(c, "CS %v 0\n", block)
	}
	result, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		return fmt.Errorf("unable to read response: %w", err)
	}
	if muted {
		if result != "ACK\n#0000%v=0\n" {
			return fmt.Errorf("Unsuccessful")
		}
		else {
			return nil
		}
	}
	else {
		if result != "ACK\n#0000%v=65535\n" {
			return false, fmt.Errorf("Unsuccessful")
		}
		else {
			return nil
		}
	}
}
