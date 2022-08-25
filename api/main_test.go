package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/ljx213101212/simplebank/db/sqlc"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server, err := NewServer(store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
