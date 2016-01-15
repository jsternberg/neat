package neat_test

import (
	"os"
	"testing"

	"github.com/jsternberg/neat"
)

func TestHost_NewHost(t *testing.T) {
	host := neat.NewHost("localhost", "127.0.0.1:2300")
	if host.Name() != "localhost" {
		t.Errorf("expected localhost, got %s", host.Name())
	}
	if host.Addr() != "127.0.0.1:2300" {
		t.Errorf("expected 127.0.0.1:2300, got %s", host.Addr())
	}
}

func TestHost_NewLocalHost(t *testing.T) {
	host, err := neat.NewLocalHost()
	if err != nil {
		t.Fatal(err)
	}

	name, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	} else if name != host.Name() {
		t.Errorf("expected hostname %s, got %s", name, host.Name())
	}

	if host.Addr() != "" {
		t.Errorf("expected no address, got %s", host.Addr())
	}
}
