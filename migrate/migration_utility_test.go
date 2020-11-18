package migrate

import (
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	log.Info("TestA")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	// Do something here before testing.
	ExecuteUp()
	log.Info("\033[1;36m{}\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here after testing.
	ExecuteDrop()
	log.Info("\033[1;36m{}\033[0m", "> Teardown completed")
}
