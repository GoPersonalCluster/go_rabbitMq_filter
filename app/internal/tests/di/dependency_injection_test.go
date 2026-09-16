package di_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/di"
	"github.com/stretchr/testify/assert"
)

func EnsureMultipleCallsToDiUseSameInstance(t *testing.T) {
	dbCon := di.NewDbConnectionDi()
	print(dbCon.ID)
	dbCon2 := di.NewDbConnectionDi()
	print(dbCon2.ID)
	assert.Equal(t, dbCon, dbCon2)
}
