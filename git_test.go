package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestParseOrigin(t *testing.T) {
	for _, test := range []struct {
		origin   string
		expected string
	}{
		{
			origin:   "git@github.com:organization/myrepo.git",
			expected: "organization/myrepo",
		},
		{
			origin:   "git@github.com:organization/myrepo",
			expected: "organization/myrepo",
		},
		{
			origin:   "https://github.com/organization/myrepo.git",
			expected: "organization/myrepo",
		},
		{
			origin:   "https://github.com/organization/myrepo",
			expected: "organization/myrepo",
		},
		{
			origin:   "https://github.com:80/organization/myrepo",
			expected: "organization/myrepo",
		},
		{
			origin:   "https://github.com:80/organization/myrepo\n",
			expected: "organization/myrepo",
		},
	} {
		t.Run(test.origin, func(t *testing.T) {
			result, err := parseOrigin(test.origin)
			require.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}
