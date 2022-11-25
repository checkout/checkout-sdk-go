package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetEventTypes(t *testing.T) {
	response, err := PreviousApi().Events.RetrieveAllEventTypes()
	assert.Nil(t, err)
	assert.NotNil(t, response)

	eventsResponseVersion, err := PreviousApi().Events.RetrieveAllEventTypes(response.EventTypes[0].Version)
	assert.Nil(t, err)
	assert.NotNil(t, eventsResponseVersion)
	assert.Equal(t, response.EventTypes[0].EventTypes, eventsResponseVersion.EventTypes[0].EventTypes)
}
