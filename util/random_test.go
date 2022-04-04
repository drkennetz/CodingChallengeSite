package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	x := RandomInt(0, 100)
	require.NotZero(t, x)
}

func TestRandomString(t *testing.T) {
	x := RandomString(5)
	require.Equal(t, 5, len(x))
}

func TestRandomFullName(t *testing.T) {
	x := RandomFullName()
	require.NotEqual(t, "", x)
}

func TestIsValidEmail(t *testing.T) {
	require.True(t, IsValidEmail("dkennetz@funky.net"))
}

func TestRandomEmail(t *testing.T) {
	x := RandomEmail()
	require.True(t, IsValidEmail(x))
}

func TestRandomGrade(t *testing.T) {
	grade := RandomGrade()
	require.Contains(t, []string{"beginner", "junior", "midlevel", "senior"}, grade)
}

func TestRandomUsername(t *testing.T) {
	x := RandomUsername()
	require.Equal(t, 8, len(x))
}

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, config)
}
