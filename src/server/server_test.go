package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xvello/letsblockit/src/filters"
)

// Checks that the server can instantiate all its components.
// This includes parsing all pages and filters.
func TestServerDryRun(t *testing.T) {
	server := NewServer(&Options{DryRun: true})
	assert.Equal(t, ErrDryRunFinished, server.Start())
}

func BenchmarkServerStart(b *testing.B) {
	o := &Options{DryRun: true}
	for n := 0; n < b.N; n++ {
		_ = NewServer(o).Start()
	}
}

func BenchmarkLoadPages(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = loadPages()
	}
}

func BenchmarkLoadFilters(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = filters.LoadFilters()
	}
}
