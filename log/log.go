package log

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	log          = logrus.New() // for production.
	rawLog       = logrus.New()
	translations = make(map[string]string)
)

// LoadTranslations takes a map[string]interface and flattens it to map[string]string
// Because translations have been loaded - we internally override log the formatter
// Nested entries are accessible using dot notation.
// example:   `{"foo": {"bar": "baz"}}`
// flattened: `foo.bar: baz`
func LoadTranslations(thing map[string]interface{}) {
	formatter := new(prefixed.TextFormatter)
	formatter.TimestampFormat = `2006-01-02 15:04:05`
	formatter.FullTimestamp = true
	log.Formatter = &TranslationFormatter{formatter}
	translations, _ = Flatten(thing)
}

// TranslationFormatter Formatter with *prefixed.TextFormatter.
type TranslationFormatter struct {
	*prefixed.TextFormatter
}

// RawFormatter Formatter without any formatter.s
type RawFormatter struct{}

// Format returns byte array from logrus.Entry.
func (f *RawFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

func init() {
	formatter := new(prefixed.TextFormatter)
	formatter.TimestampFormat = `2006-01-02 15:04:05`
	formatter.FullTimestamp = true

	log.Formatter = formatter
	rawLog.Formatter = new(RawFormatter)
}

// Get returns the initialization of logrus.Logger.
func Get() *logrus.Logger {
	switch strings.ToLower(os.Getenv("GALAXY_LOGLEVEL")) {
	case "error":
		log.Level = logrus.ErrorLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "debug":
		log.Level = logrus.DebugLevel
	default:
		log.Level = logrus.InfoLevel
	}
	return log
}

// GetRaw returns the rawLog.
func GetRaw() *logrus.Logger {
	return rawLog
}
