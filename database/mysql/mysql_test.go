package mysql

import (
	"testing"

	"github.com/chanhteam/go-utils/logger"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionShouldReturnNil(t *testing.T) {
	t.Skip()
	logger.NewDefault()
	db := GetConnection("127.0.0.1", "db_service", "root", "123456", 10, 10)
	assert.NotNil(t, db)
}
