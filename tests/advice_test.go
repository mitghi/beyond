package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/beyond/advice"
	"testing"
)

func Test_Advice(t *testing.T) {
	assert := assert.New(t)
	packages := testPackages()
	assert.NotNil(packages)
	advices := advice.GetAdvices(packages)
	assert.NotNil(advices)
	assert.NotNil(advices.List())
	assert.Len(advices.List(), 6)
	assertAdvice(assert, advices.List()[1], `NewEmptyAround`, "github.com/wesovilabs/beyond/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[1], `NewEmptyAround`, "github.com/wesovilabs/beyond/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[2], `NewComplexAround`, "github.com/wesovilabs/beyond/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[3], "NewEmptyAround", "github.com/wesovilabs/beyond/testdata/advice", true, true)
	assertAdvice(assert, advices.List()[4], "NewTracingAdvice", "github.com/wesovilabs/beyond/testdata/advice", true, false)
	assertAdvice(assert, advices.List()[5], "NewComplexBefore", "github.com/wesovilabs/beyond/testdata/advice", true, false)

}

func assertAdvice(assert *assert.Assertions, advice *advice.Advice, name, pkg string, hasBefore, hasReturning bool) {
	assert.Equal(advice.Pkg(), pkg)
	assert.Equal(advice.Name(), name)
	assert.Equal(advice.HasBefore(), hasBefore)
	assert.Equal(advice.HasReturning(), hasReturning)
}
