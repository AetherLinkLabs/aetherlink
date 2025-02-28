package gateway

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSubdomain(t *testing.T) {
	subdomain, err := getSubdomain("chat.example.com")
	require.NoError(t, err)
	require.Equal(t, "chat", subdomain)

	subdomain, err = getSubdomain("ai.chat.example.com")
	require.NoError(t, err)
	require.Equal(t, "ai.chat", subdomain)

	subdomain, err = getSubdomain("example.com")
	require.NoError(t, err)
	require.Equal(t, "", subdomain)

	subdomain, err = getSubdomain("a.b.c.example.com")
	require.NoError(t, err)
	require.Equal(t, "a.b.c", subdomain)

	subdomain, err = getSubdomain("http://aether.localhost:8081/hello.html")
	require.NoError(t, err)
	require.Equal(t, "aether", subdomain)

	subdomain, err = getSubdomain("http://aether.ai3.localhost:8081/hello.html")
	require.NoError(t, err)
	require.Equal(t, "aether.ai3", subdomain)
}
