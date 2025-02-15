package auth_test

import (
	"testing"

	"github.com/dtegapp/nexus/v3/router/auth"
	"github.com/dtegapp/nexus/v3/wamp"
	"github.com/stretchr/testify/require"
)

func TestAnonAuth(t *testing.T) {
	anonAuth := auth.AnonymousAuth{
		AuthRole: "guest",
	}

	details := wamp.Dict{
		"authid":      "someone",
		"authmethods": []string{"anonymous"}}
	welcome, err := anonAuth.Authenticate(wamp.ID(101), details, nil)
	require.NoError(t, err)

	require.NotNil(t, welcome, "received nil welcome msg")
	require.Equal(t, wamp.WELCOME, welcome.MessageType())
	s, _ := wamp.AsString(welcome.Details["authmethod"])
	require.Equal(t, "anonymous", s, "invalid authmethod in welcome details")
	s, _ = wamp.AsString(welcome.Details["authrole"])
	require.Equal(t, "guest", s, "incorrect authrole in welcome details")
}
