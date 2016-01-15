package neat

import "os"

// Host keeps host-specific information. A host may refer to something local or
// remote.
type Host interface {
	// Name is the name of the host (usually the hostname).
	Name() string

	// Addr is the address for contacting the host. If this is local and there
	// is no method of contacting the host, this may return an empty string.
	Addr() string
}

// NewHost creates a host with the given name and address.
func NewHost(name, addr string) Host {
	return &host{
		name: name,
		addr: addr,
	}
}

// NewLocalHost creates a host from the existing machine. This is a convenience
// method for looking up the hostname and creating a host with no address.
func NewLocalHost() (Host, error) {
	name, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return NewHost(name, ""), nil
}

type host struct {
	name string
	addr string
}

func (h *host) Name() string {
	return h.name
}

func (h *host) Addr() string {
	return h.addr
}
