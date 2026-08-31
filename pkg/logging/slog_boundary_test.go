// Copyright (c) 2026 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package logging

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	goruntime "runtime"
	"strconv"
	"testing"
)

func TestRepositoryDoesNotEmitThroughSlogDirectly(t *testing.T) {
	_, currentFile, _, ok := goruntime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller() failed")
	}
	repositoryRoot := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", ".."))
	loggingDir := filepath.Clean(filepath.Join(repositoryRoot, "pkg", "logging"))
	bannedFunctions := map[string]bool{
		"Default": true, "SetDefault": true, "Log": true, "LogAttrs": true,
		"Debug": true, "Info": true, "Warn": true, "Error": true,
	}

	err := filepath.WalkDir(repositoryRoot, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			switch entry.Name() {
			case ".git", ".cache":
				return filepath.SkipDir
			}
			if filepath.Clean(path) == loggingDir {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}

		fileSet := token.NewFileSet()
		parsed, err := parser.ParseFile(fileSet, path, nil, parser.SkipObjectResolution)
		if err != nil {
			return err
		}
		slogNames := make(map[string]bool)
		for _, spec := range parsed.Imports {
			importPath, err := strconv.Unquote(spec.Path.Value)
			if err != nil || importPath != "log/slog" {
				continue
			}
			name := "slog"
			if spec.Name != nil {
				name = spec.Name.Name
			}
			if name == "." {
				t.Errorf("%s: dot import of log/slog bypasses the logging boundary", relativePath(repositoryRoot, path))
				continue
			}
			slogNames[name] = true
		}
		if len(slogNames) == 0 {
			return nil
		}

		ast.Inspect(parsed, func(node ast.Node) bool {
			selector, ok := node.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			identifier, ok := selector.X.(*ast.Ident)
			if !ok || !slogNames[identifier.Name] {
				return true
			}
			if selector.Sel.Name == "Logger" || bannedFunctions[selector.Sel.Name] {
				position := fileSet.Position(selector.Pos())
				t.Errorf("%s:%d: direct slog %s use bypasses pkg/logging", relativePath(repositoryRoot, path), position.Line, selector.Sel.Name)
			}
			return true
		})
		return nil
	})
	if err != nil {
		t.Fatalf("inspect repository: %v", err)
	}
}

func relativePath(root, path string) string {
	relative, err := filepath.Rel(root, path)
	if err != nil {
		return path
	}
	return filepath.ToSlash(relative)
}
