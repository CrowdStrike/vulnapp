package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type commandEntry struct {
	Path        string `json:"path"`
	Command     string `json:"command"`
	Description string `json:"description,omitempty"`
}

// parseScriptHeaders reads the leading comment block of a detection script and
// returns its Shortname and Description headers. A Description may wrap onto
// following comment lines that are indented past the usual "# " (two or more
// spaces, or a tab, after the "#"); those lines are joined with single spaces.
func parseScriptHeaders(path string) (shortname, description string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", fmt.Errorf("open %s: %w", path, err)
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	inDescription := false
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			break
		}
		if inDescription {
			if text, ok := continuationText(line); ok {
				description += " " + text
				continue
			}
			inDescription = false
		}
		if v, ok := strings.CutPrefix(line, "# Shortname:"); ok {
			shortname = strings.TrimSpace(v)
		} else if v, ok := strings.CutPrefix(line, "# Description:"); ok {
			description = strings.TrimSpace(v)
			inDescription = true
		}
	}
	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("read %s: %w", path, err)
	}
	return shortname, description, nil
}

// continuationText reports whether a comment line continues the previous
// header and, if so, returns its trimmed text.
func continuationText(line string) (string, bool) {
	rest := strings.TrimPrefix(line, "#")
	if !strings.HasPrefix(rest, "  ") && !strings.HasPrefix(rest, "\t") {
		return "", false
	}
	text := strings.TrimSpace(rest)
	return text, text != ""
}

func scanDirectory(dir string) ([]commandEntry, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read directory %s: %w", dir, err)
	}

	seen := map[string]string{}
	var cmds []commandEntry
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sh") {
			continue
		}

		path := filepath.Join(dir, e.Name())
		shortname, description, err := parseScriptHeaders(path)
		if err != nil {
			return nil, err
		}
		if shortname == "" || description == "" {
			continue
		}

		urlPath := "/" + shortname
		if prev, ok := seen[urlPath]; ok {
			return nil, fmt.Errorf("duplicate path %q: %s and %s", urlPath, prev, e.Name())
		}
		seen[urlPath] = e.Name()

		cmds = append(cmds, commandEntry{
			Path:        urlPath,
			Command:     "./" + filepath.Join("bin", e.Name()),
			Description: description,
		})
	}

	if len(cmds) == 0 {
		return nil, fmt.Errorf("no valid detection scripts found in %s", dir)
	}

	return cmds, nil
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: gendetections <scripts-dir>")
	}

	cmds, err := scanDirectory(os.Args[1])
	if err != nil {
		return err
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(cmds); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "gendetections: %s\n", err)
		os.Exit(1)
	}
}
