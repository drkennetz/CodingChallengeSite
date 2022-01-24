package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	// main entry point of all tests of a specific golang package
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}