package config

import (
	"io/ioutil"
	"os"
	"sync"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const (
	// AppName app name.
	AppName = "Galaxy"
	// VERSION mark app version.
	VERSION = "0.0.1"
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

// InitialiseSystem done all things about app runtime env configurationss.
func InitialiseSystem() error {
	SetNodeID("solo-" + uuid.NewV4().String())

	if IsRunningTests() && GetLogLevel() == "" {
		// `go test` without GALAXY_LOGLEVEL set defaults to no log
		// output
		log.Level = logrus.ErrorLevel
		log.Out = ioutil.Discard
	}

	if !IsRunningTests() {
		globalConf := Config{}
		if err := Load(confPaths, &globalConf); err != nil {
			return err
		}
		if globalConf.PIDFileLocation == "" {
			globalConf.PIDFileLocation = "/var/run/galaxy/galaxy_service.pid"
		}
		// It's necessary to set global conf before and after calling afterConfSetup as global conf
		// is being used by dependencies of the even handler init and then conf is modified again.
		SetGlobal(globalConf)
		afterConfSetup(&globalConf)
		SetGlobal(globalConf)
	}
	log.WithFields(logrus.Fields{
		"App":    AppName,
		"NodeID": NodeID,
	}).Infof("Initialised env setup")
	return nil
}

// afterConfSetup takes care of non-sensical config values (such as zero
// timeouts) and sets up a few globals that depend on the config.
func afterConfSetup(conf *Config) {
	// TODO somethings after config setup
}

// GetLogLevel returns log level of current app.
func GetLogLevel() string {
	return os.Getenv(AppLogLevel)
}

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
