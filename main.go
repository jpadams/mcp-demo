package main

import (
	"context"
	//"dagger/mcp-demo/internal/dagger"
)

type McpDemo struct{}

// Returns "heeeey"
func (m *McpDemo) Hello(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"sh", "-c", "echo -n heeeey"}).
		Stdout(ctx)
}
