package netcfg

import (
	"fmt"
	"strings"
)

var pathCutset string = "/"

type ServerConfiguration struct {
	HttpAddress *HttpAddressConfiguration
	Prefix      string
}

func (self *ServerConfiguration) PrefixedPath(path string) string {
	trimmedPrefix := strings.TrimLeft(self.Prefix, pathCutset)
	trimmedPath := strings.TrimLeft(path, pathCutset)

	if trimmedPrefix == "" {
		return fmt.Sprintf("/%s", trimmedPath)
	}

	return fmt.Sprintf("/%s/%s", trimmedPrefix, trimmedPath)
}
