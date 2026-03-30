// Package reexec provides a minimal process re-execution facility for tests.
// It replaces github.com/docker/docker/pkg/reexec to eliminate the Docker
// dependency and its associated CVEs.
package reexec

import (
	"fmt"
	"os"
	"path/filepath"
)

var registeredInitializers = make(map[string]func())

// Register adds an initializer function that will be called if the process
// is launched with argv[0] matching the given name.
func Register(name string, initializer func()) {
	if _, exists := registeredInitializers[name]; exists {
		panic(fmt.Sprintf("reexec: already registered initializer with name %q", name))
	}
	registeredInitializers[name] = initializer
}

// Init checks if the current process's argv[0] matches a registered name
// and, if so, runs the corresponding initializer. Returns true if an
// initializer was found and executed.
func Init() bool {
	name := filepath.Base(os.Args[0])
	initializer, ok := registeredInitializers[name]
	if !ok {
		return false
	}
	initializer()
	return true
}

// Self returns the path to the current process's executable.
func Self() string {
	name, err := os.Executable()
	if err != nil {
		name = os.Args[0]
	}
	return name
}
