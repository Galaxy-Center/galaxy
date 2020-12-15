package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	logger "github.com/galaxy-center/galaxy/log"
	"github.com/kelseyhightower/envconfig"
)

var (
	log      = logger.Get()
	globalMu sync.Mutex
	global   atomic.Value

	// confPaths is the series of paths to try to use as config files. The
	// first one to exist will be used. If none exists, a default config
	// will be written to the first path in the list.
	//
	// When --conf=foo is used, this will be replaced by []string{"foo"}.
	confPaths = []string{
		"galaxy.conf",
		"~/.config/galaxy/galaxy.conf",
		"/etc/galaxy/galaxy.conf",
	}

	// Default writes to First-Level path while not found any config file.
	Default = Config{
		MySQLConfig: DBConfig{
			User:     "lance",
			Password: "Lancexu@1992",
			Host:     "localhost",
			Port:     3306,
			Database: "galaxy_test",
		},
		App: App{
			AppName: AppName,
			Version: VERSION,
		},
	}
)

const (
	// MySQLDSNFormat string format
	MySQLDSNFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local"
	envPrefix      = "GALAXY"
)

func init() {
	SetGlobal(Config{})
	InitialiseSystem()
}

// SetGlobal store global config to atomic.Value
func SetGlobal(conf Config) {
	globalMu.Lock()
	defer globalMu.Unlock()
	global.Store(conf)
}

// Global returns global config from atomic.Value
func Global() Config {
	return global.Load().(Config)
}

// App cache node info.
type App struct {
	AppName string `json:"appName"`
	Version string `json:"version"`
	StartAt uint64 `json:"startAt"`
	NodeID  string `json:"nodeID"`
}

// DBConfig defines for DB connection.
type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

// DSNFormat returns builder dsn string.
func (c DBConfig) DSNFormat() string {
	return fmt.Sprintf(MySQLDSNFormat, c.User, c.Password, c.Host, c.Port, c.Database)
}

// LivenessCheckConfig liveness check duration.
type LivenessCheckConfig struct {
	CheckDuration time.Duration `json:"check_duration"`
}

// Config global configs.
type Config struct {
	// OriginalPath is the path to the config file that was read. If
	// none was found, it's the path to the default config file that
	// was written.
	OriginalPath    string `json:"-"`
	TemplatePath    string `json:"template_path"`
	PIDFileLocation string `json:"pid_file_location"`

	LivenessCheck LivenessCheckConfig `json:"liveness_check"`

	MySQLConfig DBConfig `json:"mysql_config"`

	App App `json:"app"`
}

// Load will load a configuration file, trying each of the paths given
// and using the first one that is a regular file and can be opened.
//
// If none exists, a default config will be written to the first path in
// the list.
//
// An error will be returned only if any of the paths existed but was
// not a valid config file.
func Load(paths []string, conf *Config) error {
	var r io.Reader
	for _, path := range paths {
		f, err := os.Open(path)
		if err == nil {
			r = f
			conf.OriginalPath = path
			break
		}
		if os.IsNotExist(err) {
			continue
		}
		return err
	}
	if r == nil {
		path := paths[0]
		log.Warnf("No config file found, writing default to %s", path)
		if err := WriteDefault(path, conf); err != nil {
			return err
		}
		log.Info("Loading default configuration...")
		return Load([]string{path}, conf)
	}
	if err := json.NewDecoder(r).Decode(&conf); err != nil {
		return fmt.Errorf("couldn't unmarshal config: %v", err)
	}
	if err := envconfig.Process(envPrefix, conf); err != nil {
		return fmt.Errorf("failed to process config env vars: %v", err)
	}
	return nil
}

// WriteDefault will set conf to the default config and write it to disk
// in path, if the path is non-empty.
func WriteDefault(path string, conf *Config) error {
	_, b, _, _ := runtime.Caller(0)
	configPath := filepath.Dir(b)
	rootPath := filepath.Dir(configPath)
	Default.TemplatePath = filepath.Join(rootPath, "templates")

	*conf = Default
	if err := envconfig.Process(envPrefix, conf); err != nil {
		return err
	}
	if path == "" {
		return nil
	}
	return WriteConf(path, conf)
}

// WriteConf writes conf to specific file.
func WriteConf(path string, conf *Config) error {
	bs, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bs, 0644)
}
