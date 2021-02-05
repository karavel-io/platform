package plan

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlan_GetComponent(t *testing.T) {
	c1 := Component{
		name:         "c1",
		version:      "0.1.0",
		dependencies: []string{"c2", "c3"},
	}

	p := New()

	p.AddComponent(c1)

	assert.Equal(t, &c1, p.GetComponent("c1"))
	assert.Nil(t, p.GetComponent("c2"))
}

func TestPlan_HasComponent(t *testing.T) {
	c1 := Component{
		name:         "c1",
		version:      "0.1.0",
		dependencies: []string{"c2", "c3"},
	}

	p := New()

	p.AddComponent(c1)

	assert.True(t, p.HasComponent("c1"))
	assert.False(t, p.HasComponent("c2"))
}

func TestPlan_Validate(t *testing.T) {
	c1 := Component{
		name:         "c1",
		version:      "0.1.0",
		dependencies: []string{"c2", "c3"},
	}

	c2 := Component{
		name:         "c2",
		version:      "0.1.0",
		dependencies: []string{"c1", "c3"},
	}

	c3 := Component{
		name:    "c3",
		version: "0.1.0",
	}

	p := New()

	p.AddComponent(c1)
	p.AddComponent(c2)
	p.AddComponent(c3)

	assert.NoError(t, p.Validate())
}

func TestPlan_ValidateMissingDep(t *testing.T) {
	c1 := Component{
		name:         "c1",
		version:      "0.1.0",
		dependencies: []string{"c2"},
	}

	p := New()

	p.AddComponent(c1)

	err := p.Validate()
	if assert.Error(t, err) {
		assert.Equal(t, "missing dependency: component 'c1' requires 'c2'", err.Error())
	}
}

func TestPlan_IntegrationsProcessing(t *testing.T) {
	c1 := Component{
		name:    "c1",
		version: "0.1.0",
		integrationsDeps: map[string][]string{
			"cool.feature": {"c2"},
		},
	}

	c2 := Component{
		name:    "c2",
		version: "0.1.0",
		integrationsDeps: map[string][]string{
			"cool.feature": {"c3", "c1"},
		},
	}

	p := New()

	p.AddComponent(c1)
	p.AddComponent(c2)

	assert.NoError(t, p.Validate())

	c1p := p.GetComponent("c1")
	c2p := p.GetComponent("c2")

	assert.True(t, c1p.integrations["cool.feature"])
	assert.False(t, c2p.integrations["cool.feature"])
}