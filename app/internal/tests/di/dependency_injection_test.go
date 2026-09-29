package di_test

import (
	"testing"

	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestEnsureMultipleCallsToDiUseSameInstance(t *testing.T) {
	dbCon := di.NewDbConnectionDi()
	println(dbCon.ID.ID())
	dbCon2 := di.NewDbConnectionDi()
	println(dbCon2.ID.ID())
	assert.Equal(t, dbCon, dbCon2)
}
