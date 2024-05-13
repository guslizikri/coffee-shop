package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "abcd12345"
var hasedPassword string
var errors error

func TestHashPass(t *testing.T) {
	hasedPassword, errors = HashPass(password)
	assert.NoError(t, errors, "error occured while hasing password")
	assert.NotEqual(t, password, hasedPassword, "password tidak terhasing")
}

// untuk kasus ini test hash dan verify harus dijalankan barenagan,
// karena test verify memakai var hashedpassword yang sudah diubah saat test hashpass
func TestVerifyPassword(t *testing.T) {
	t.Run("verify success", func(t *testing.T) {
		var hasPassword = VerifyPass(hasedPassword, password)
		assert.Nil(t, hasPassword, "password salah")
	})

	t.Run("verify failed", func(t *testing.T) {
		var hasPassword = VerifyPass(hasedPassword, "12345")
		assert.Nil(t, hasPassword, "password masih salah")
	})
}
