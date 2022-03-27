package suffixmap_test

import (
	"testing"

	"gotest.tools/assert"

	. "github.com/stackql/go-suffix-map/pkg/suffixmap"
)

type testingVal struct {
	K string
	V string
}

var (
	superSimpleTestSubstrate map[string]interface{} = map[string]interface{}{
		"providerX.serviceY.resourceZ.parameterAlpha": testingVal{K: "k", V: "v"},
		"providerA.serviceB.resourceC.parameterBeta":  testingVal{K: "kBeta", V: "vBeta"},
	}
)

func TestSimpleSuffixMap(t *testing.T) {
	sm := NewSuffixMap(nil)
	assert.Equal(t, sm.Size(), 0)
}

func TestBasicLookup(t *testing.T) {
	sm := NewSuffixMap(superSimpleTestSubstrate)
	assert.Equal(t, sm.Size(), 2)

	// simple lookup
	rv, ok := sm.Get("parameterAlpha")
	assert.Assert(t, ok)
	castResult := rv.(testingVal)
	assert.Assert(t, ok)
	assert.Equal(t, castResult.K, "k")
	assert.Equal(t, castResult.V, "v")

	// simple lookup again
	rv, ok = sm.Get("parameterBeta")
	assert.Assert(t, ok)
	castResult = rv.(testingVal)
	assert.Assert(t, ok)
	assert.Equal(t, castResult.K, "kBeta")
	assert.Equal(t, castResult.V, "vBeta")

	// failed lookup
	_, ok = sm.Get("a")
	assert.Assert(t, !ok)

	// longer qualifitication
	rv, ok = sm.Get("resourceC.parameterBeta")
	assert.Assert(t, ok)
	castResult = rv.(testingVal)
	assert.Assert(t, ok)
	assert.Equal(t, castResult.K, "kBeta")
	assert.Equal(t, castResult.V, "vBeta")

	// longer failed lookup
	_, ok = sm.Get("ourceC.parameterBeta")
	assert.Assert(t, !ok)

	// still longer qualifitication
	rv, ok = sm.Get("serviceB.resourceC.parameterBeta")
	assert.Assert(t, ok)
	castResult = rv.(testingVal)
	assert.Assert(t, ok)
	assert.Equal(t, castResult.K, "kBeta")
	assert.Equal(t, castResult.V, "vBeta")

	// still longer failed lookup
	_, ok = sm.Get("rviceB.resourceC.parameterBeta")
	assert.Assert(t, !ok)

	// longest qualifitication
	rv, ok = sm.Get("providerA.serviceB.resourceC.parameterBeta")
	assert.Assert(t, ok)
	castResult = rv.(testingVal)
	assert.Assert(t, ok)
	assert.Equal(t, castResult.K, "kBeta")
	assert.Equal(t, castResult.V, "vBeta")

	// longest failed lookup
	_, ok = sm.Get("oviderA.serviceB.resourceC.parameterBeta")
	assert.Assert(t, !ok)
}
