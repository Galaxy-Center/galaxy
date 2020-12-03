package config

import (
	"sync"
)

const (
	// AppName app name.
	AppName = "Galaxy"
	// AppLogLevel log level.
	AppLogLevel = "GALAXY_LOGLEVEL"
)

var (
	runningTestsMu sync.RWMutex
	muNodeID       sync.Mutex

	// NodeID mark for current node.
	NodeID string
	// default is false.
	testMode bool
)

// IsRunningTests returns true if current system is test mode.
func IsRunningTests() bool {
	runningTestsMu.RLock()
	v := testMode
	runningTestsMu.RUnlock()
	return v
}

// SetTestMode for unitest case.
func SetTestMode(v bool) {
	runningTestsMu.Lock()
	testMode = v
	runningTestsMu.Unlock()
}

// SetNodeID writes NodeID safely.
func SetNodeID(nodeID string) {
	muNodeID.Lock()
	NodeID = nodeID
	muNodeID.Unlock()
}

// GetNodeID reads NodeID safely.
func GetNodeID() string {
	muNodeID.Lock()
	defer muNodeID.Unlock()
	return NodeID
}
