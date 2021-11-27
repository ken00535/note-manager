package usecase

import (
	"testing"

	"note-manager/pkg/infra/config"
	"note-manager/pkg/infra/logger"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateToken(t *testing.T) {
	type Repo struct{}
	config.Init(logger.NewMockLogger())
	u := NewAuthUsecase(Repo{})
	token, err := u.GetToken("ken")
	assert.NoError(t, err)
	err = u.ValidateToken(token)
	assert.NoError(t, err)
}
