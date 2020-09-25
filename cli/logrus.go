package cli

import (
	"bytes"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	logLevel string
)

func init() {
	log.SetFormatter(&SimpleFormatter{})
	log.SetOutput(os.Stdout)
}

//SimpleFormatter implements the Logrus Formatter interface
type SimpleFormatter struct {
}

//Format is the logrus Formatter implementatino: just print message to output
func (s *SimpleFormatter) Format(e *log.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if e.Buffer != nil {
		b = e.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(e.Message)
	b.WriteByte('\n')
	return b.Bytes(), nil

}

func setLogLevel(logLevel string) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("Unsupported loglevel '%s': only 'trace', 'debug', 'info', 'warn', 'error', 'fatal' and 'panic' log levels are supported", logLevel)
	}
	log.SetLevel(level)
	log.Debugf("log level set to '%s'", level.String())
	return nil
}
