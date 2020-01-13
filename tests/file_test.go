package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	d := ReadAll()
	assert.Equal(t, `module github.com/nikandfor/embed

go 1.13

require (
	github.com/golang/snappy v0.0.1
	github.com/nikandfor/cli v0.0.0-20191209221237-ad6c97f5afac
	github.com/nikandfor/json v0.0.0-20191226181838-acc5fc24730f // indirect
	github.com/nikandfor/tlog v0.3.0
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.3.0
)
`, string(d))
}
