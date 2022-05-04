package api

import (
	"testing"

	db "github.com/drkennetz/codingchallenge/db/sqlc"
	"github.com/drkennetz/codingchallenge/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserAPI(t *testing.T) {
	return

}

func randomUser(t *testing.T) (user db.User, password string) {
	acc := randomAccount()

	password = util.RandomString(9)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	user = db.User{
		AccountID: acc.ID,
		AdminUser: false,
		Username: "tester",
		Password: hashedPassword,
		Grade: db.DevLevelBeginner,
	}
	return
}