package main

import (
	"context"
//	"dagger/mcp-demo/internal/dagger"
)

type McpDemo struct{}

// Returns lines that match a pattern in the files of the provided Directory
func (m *McpDemo) Hello(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"sh", "-c", "echo heeeey"}).
		Stdout(ctx)
}
