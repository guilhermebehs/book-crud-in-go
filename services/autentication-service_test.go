package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAndValidateJWTToken(t *testing.T) {

	t.Run("should create and validate token successfully", func(t *testing.T) {

		os.Setenv("AUTHENTICATED_USER", "admin")
		os.Setenv("AUTHENTICATED_PASS", "admin")

		authenticationService := CreateAuthenticationService()
		token, err := authenticationService.Authenticate("admin", "admin")
		assert.Equal(t, err, nil)
		assert.NotEqual(t, token, "")
		assert.Equal(t, authenticationService.Validate(token), true)
	})

	t.Run("should return error when creating token with wrong credentials", func(t *testing.T) {

		os.Setenv("AUTHENTICATED_USER", "admin")
		os.Setenv("AUTHENTICATED_PASS", "admin")

		authenticationService := CreateAuthenticationService()
		token, err := authenticationService.Authenticate("admin2", "admin2")
		assert.Equal(t, err.Error(), "Unauthorized")
		assert.Equal(t, token, "")
	})

	t.Run("should return error when validating invalid token", func(t *testing.T) {

		os.Setenv("AUTHENTICATED_USER", "admin")
		os.Setenv("AUTHENTICATED_PASS", "admin")

		authenticationService := CreateAuthenticationService()
		token, err := authenticationService.Authenticate("admin", "admin")
		assert.Equal(t, err, nil)
		assert.NotEqual(t, token, "")
		assert.Equal(t, authenticationService.Validate("something else"), false)
	})
}
