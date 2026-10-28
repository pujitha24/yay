//go:build !integration

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Jguer/yay/v13/pkg/settings/parser"
)

func TestQualifyAURTargets(t *testing.T) {
	t.Parallel()

	targets := []string{"doas", "extra/foo", "aur/bar", "baz>=1.0"}

	assert.Equal(t, targets, QualifyAURTargets(parser.ModeAny, targets))
	assert.Equal(t, []string{"aur/doas", "extra/foo", "aur/bar", "aur/baz>=1.0"},
		QualifyAURTargets(parser.ModeAUR, targets))
	assert.Equal(t, []string{"doas", "extra/foo", "aur/bar", "baz>=1.0"}, targets)
}
