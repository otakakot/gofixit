package gofixit

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

func Run(fn func(ast.Node, *token.FileSet) bool) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	target := "."

	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	files, err := collect(ctx, target)
	if err != nil {
		panic(err)
	}

	count := 0

	for _, file := range files {
		if fixed, err := process(ctx, file, fn); err != nil {
			panic(err)
		} else if fixed {
			count++
		}
	}

	fmt.Printf("%d files fixed.\n", count)
}

func collect(
	ctx context.Context,
	target string,
) ([]string, error) {
	files := []string{}

	info, err := os.Stat(target)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		if err := filepath.WalkDir(target, func(path string, dir fs.DirEntry, err error) error {
			if ctx.Err() != nil {
				return ctx.Err()
			}

			if err != nil {
				return err
			}

			if dir.IsDir() && dir.Name() == "vendor" {
				return filepath.SkipDir
			}

			if !dir.IsDir() && strings.HasSuffix(path, ".go") {
				files = append(files, path)
			}

			return nil
		}); err != nil {
			return nil, err
		}
	} else if strings.HasSuffix(target, ".go") {
		abs, err := filepath.Abs(target)
		if err != nil {
			return nil, fmt.Errorf("get absolute path: %w", err)
		}

		files = append(files, abs)
	}

	return files, nil
}

func process(
	ctx context.Context,
	filename string,
	fn func(ast.Node, *token.FileSet) bool,
) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	src, err := os.ReadFile(filename)
	if err != nil {
		return false, fmt.Errorf("read error: %w", err)
	}

	if bytes.Contains(src, []byte("DO NOT EDIT.")) {
		return false, nil
	}

	fset := token.NewFileSet()

	astN, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return false, fmt.Errorf("parse error: %w", err)
	}

	ast.Inspect(astN, func(n ast.Node) bool {
		if ctx.Err() != nil {
			return false
		}

		return fn(n, fset)
	})

	buf := bytes.Buffer{}

	if err := printer.Fprint(&buf, fset, astN); err != nil {
		return false, fmt.Errorf("printer: %w", err)
	}

	if bytes.Equal(src, buf.Bytes()) {
		return false, nil
	}

	formatted := bytes.Buffer{}

	if err := format.Node(&formatted, fset, astN); err != nil {
		return false, fmt.Errorf("format: %w", err)
	}

	res, err := format.Source(formatted.Bytes())
	if err != nil {
		return false, fmt.Errorf("format: %w", err)
	}

	if err := os.WriteFile(filename, res, 0644); err != nil {
		return false, fmt.Errorf("write: %w", err)
	}

	fmt.Println("fixed:", filename)

	return true, nil
}
