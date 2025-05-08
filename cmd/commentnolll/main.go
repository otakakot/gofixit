package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
)

func main() {
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
		if fixed, err := process(ctx, file); err != nil {
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
) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	src, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return false, err
	}

	lines := []string{}

	scanner := bufio.NewScanner(bytes.NewReader(src))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	insertLines := map[int]struct{}{}

	for _, cg := range node.Comments {
		needInsert := false
		for _, c := range cg.List {
			for line := range strings.SplitSeq(c.Text, "\n") {
				if len(line) > 120-3 {
					needInsert = true
					break
				}
			}
			if needInsert {
				break
			}
		}
		if needInsert {
			pos := fset.Position(cg.Pos())
			insertLines[pos.Line] = struct{}{}
		}
	}

	if len(insertLines) == 0 {
		return false, nil
	}

	var out []string

	for i, line := range lines {
		if _, ok := insertLines[i+1]; ok {
			out = append(out, "//nolint:lll")
		}
		out = append(out, line)
	}

	output := strings.Join(out, "\n")

	res, err := format.Source([]byte(output))
	if err != nil {
		return false, fmt.Errorf("format: %w", err)
	}

	if err := os.WriteFile(filename, res, 0644); err != nil {
		return false, err
	}

	return true, nil
}
