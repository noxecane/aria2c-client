package aria2c

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

// RPCResult Result of an RPC request holding a possible error
type RPCResult struct {
	Result string
	Error  error
}

type ariaConfig struct {
	Continue    bool   `json:"continue"`
	Dir         string `json:"dir"`
	Parallelism int    `json:"max-concurrent-downloads"`
	Connections int    `json:"max-connection-per-server"`
}

var defaultConfig = ariaConfig{true, "./", 1, 1}

// AriaRPCURL default url to access aria2c
const AriaRPCURL = "http://localhost:6800/jsonrpc"

// AddDownload add a download to aria2c using
// the default settings
func AddDownload(done chan<- RPCResult, urls ...string) {
	result := ""

	// connect to aria2c
	client, err := rpc.Dial(AriaRPCURL)
	if err != nil {
		done <- RPCResult{result, err}
	}

	fmt.Println("Adding", urls, "to aria2c")
	// make your call
	err = client.Call(&result, "aria2.addUri", urls, defaultConfig)
	if err != nil {
		done <- RPCResult{result, err}
	}

	// finish up
	done <- RPCResult{result, nil}
}
