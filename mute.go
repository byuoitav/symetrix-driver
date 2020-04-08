package symetrix

include ("net"
         "dsp"
         "fmt"
         "os"
         "strings"
         "bufio"
)

// GetMutedByBlock returns true if the given block is muted.
func (d *DSP) GetMutedByBlock(ctx context.Context, block string) (bool, error) {

	c, err := net.Dial("tcp", Address+":48631")
	if err != nil {
    		fmt.Println("unable to establish TCP client")
    	return
	}
	Fprintf(c,"GS "+block+'\n')
	if (bufio.NewReader(c).ReadString('\n') == "0\n") {
  		return false
	}
	return true
}

// SetMutedByBlock sets the mute state on the given block
func (d *DSP) SetMutedByBlock(ctx context.Context, block string, muted bool) (error) {
	if (muted) {
		return nil
	}
	
	c, err := net.Dial("tcp", Address+":48631")
	if err != nil {
    		fmt.Println("unable to establish TCP client")
    	return
	}
	Fprintf(c,"CS "+block+"65535\n")
}
