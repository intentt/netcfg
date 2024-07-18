package netcfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyPrefix(t *testing.T) {
	serverConfiguration := &ServerConfiguration{
		Prefix: "/",
	}

	assert.Equal(
		t,
		"/",
		serverConfiguration.PrefixedPath("/"),
	)

	assert.Equal(
		t,
		"/bar/",
		serverConfiguration.PrefixedPath("/bar/"),
	)

	assert.Equal(
		t,
		"/foo",
		serverConfiguration.PrefixedPath("/foo"),
	)
}

func TestServerConfigurationGeneratesRoutes(t *testing.T) {
	serverConfiguration := &ServerConfiguration{
		Prefix: "/test",
	}

	assert.Equal(
		t,
		"/test/",
		serverConfiguration.PrefixedPath("/"),
	)

	assert.Equal(
		t,
		"/test/bar/",
		serverConfiguration.PrefixedPath("/bar/"),
	)

	assert.Equal(
		t,
		"/test/foo",
		serverConfiguration.PrefixedPath("/foo"),
	)
}
