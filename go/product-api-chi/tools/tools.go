//go:build tools
// +build tools

package tools

import (
	// Required for pre-commit hooks go-critic
	_ "github.com/quasilyte/go-ruleguard"
)
