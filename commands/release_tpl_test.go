package commands

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestRenderReleaseTpl(t *testing.T) {
	msg, err := renderReleaseTpl("Creating", "#", "1.0", "shawncatz/bub", "master")
	assert.Equal(t, nil, err)

	expMsg := `
# Creating release 1.0 for shawncatz/bub from master
#
# Write a message for this release. The first block of
# text is the title and the rest is the description.`
	assert.Equal(t, expMsg, msg)
}
