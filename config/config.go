package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/kelseyhightower/envconfig"
)

var(
	log 
)

const (
	// MySQLDSNFormat string format
	MySQLDSNFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local"
)

// DBConfig defines for DB connection.
type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func (c *DBConfig) format() string {
	return fmt.Sprintf(MySQLDSNFormat, c.User, c.Password, c.Host, c.Port, c.Database)
}

// Config global configs.
type Config struct {
	MySQLConfig DBConfig `json:"mysql_config"`
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
	if err := processCustom(envPrefix, conf, loadZipkin, loadJaeger); err != nil {
		return fmt.Errorf("failed to process config custom loader: %v", err)
	}
	return nil
}
