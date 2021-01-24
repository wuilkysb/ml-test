package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer_WhenServerIsNotNilAndSameType(t *testing.T) {
	server := NewServer()

	assert.NotNil(t, server)
	assert.Equal(t, fmt.Sprintf("%T", echo.New()), fmt.Sprintf("%T", server))
}
