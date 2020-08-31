package config

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/99designs/gqlgen/internal/code"
)

func TestAutobinding(t *testing.T) {
	t.Run("with file path", func(t *testing.T) {
		cfg := Config{
			Models: TypeMap{},
			AutoBind: []string{
				"../chat",
			},
			Packages: &code.Packages{},
		}

		cfg.Schema = gqlparser.MustLoadSchema(&ast.Source{Name: "TestAutobinding.schema", Input: `
			scalar Banned
			type Message { id: ID }
		`})

		require.EqualError(t, cfg.autobind(), "unable to load ../chat - make sure you're using an import path to a package that exists")
	})
}
