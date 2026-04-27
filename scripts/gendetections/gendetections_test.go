package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeScript(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o755); err != nil {
		t.Fatal(err)
	}
}

func Test_parseScriptHeaders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		content         string
		wantShortname   string
		wantDescription string
	}{
		{
			name:            "both headers",
			content:         "#!/bin/sh\n# Shortname: rootkit\n# Description: Does rootkit things.\necho hello\n",
			wantShortname:   "rootkit",
			wantDescription: "Does rootkit things.",
		},
		{
			name:            "missing shortname",
			content:         "#!/bin/sh\n# Description: Some desc.\necho hello\n",
			wantShortname:   "",
			wantDescription: "Some desc.",
		},
		{
			name:            "missing description",
			content:         "#!/bin/sh\n# Shortname: foo\necho hello\n",
			wantShortname:   "foo",
			wantDescription: "",
		},
		{
			name:            "no headers",
			content:         "#!/bin/sh\necho hello\n",
			wantShortname:   "",
			wantDescription: "",
		},
		{
			name:            "stops at non-comment line",
			content:         "#!/bin/sh\n# Shortname: before\necho break\n# Description: after\n",
			wantShortname:   "before",
			wantDescription: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			path := filepath.Join(dir, "test.sh")
			if err := os.WriteFile(path, []byte(tt.content), 0o644); err != nil {
				t.Fatal(err)
			}

			shortname, description, err := parseScriptHeaders(path)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if shortname != tt.wantShortname {
				t.Errorf("shortname = %q, want %q", shortname, tt.wantShortname)
			}
			if description != tt.wantDescription {
				t.Errorf("description = %q, want %q", description, tt.wantDescription)
			}
		})
	}
}

func Test_scanDirectory(t *testing.T) {
	t.Parallel()

	t.Run("valid scripts", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "alpha.sh",
			"#!/bin/sh\n# Shortname: alpha\n# Description: Alpha detection.\necho a\n")
		writeScript(t, dir, "beta.sh",
			"#!/bin/sh\n# Shortname: beta\n# Description: Beta detection.\necho b\n")

		cmds, err := scanDirectory(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(cmds) != 2 {
			t.Fatalf("got %d entries, want 2", len(cmds))
		}
		if cmds[0].Path != "/alpha" || cmds[0].Command != "./bin/alpha.sh" {
			t.Errorf("entry 0 = %+v", cmds[0])
		}
		if cmds[1].Path != "/beta" || cmds[1].Command != "./bin/beta.sh" {
			t.Errorf("entry 1 = %+v", cmds[1])
		}
	})

	t.Run("skips missing shortname", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "good.sh",
			"#!/bin/sh\n# Shortname: good\n# Description: Good.\necho g\n")
		writeScript(t, dir, "bad.sh",
			"#!/bin/sh\n# Description: No shortname.\necho b\n")

		cmds, err := scanDirectory(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(cmds) != 1 {
			t.Fatalf("got %d entries, want 1", len(cmds))
		}
		if cmds[0].Path != "/good" {
			t.Errorf("expected /good, got %s", cmds[0].Path)
		}
	})

	t.Run("skips missing description", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "good.sh",
			"#!/bin/sh\n# Shortname: good\n# Description: Good.\necho g\n")
		writeScript(t, dir, "bad.sh",
			"#!/bin/sh\n# Shortname: nodesc\necho b\n")

		cmds, err := scanDirectory(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(cmds) != 1 {
			t.Fatalf("got %d entries, want 1", len(cmds))
		}
	})

	t.Run("skips non-sh files", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "good.sh",
			"#!/bin/sh\n# Shortname: good\n# Description: Good.\necho g\n")
		writeScript(t, dir, "readme.txt",
			"# Shortname: fake\n# Description: Not a script.\n")

		cmds, err := scanDirectory(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(cmds) != 1 {
			t.Fatalf("got %d entries, want 1", len(cmds))
		}
	})

	t.Run("skips directories", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "good.sh",
			"#!/bin/sh\n# Shortname: good\n# Description: Good.\necho g\n")
		if err := os.Mkdir(filepath.Join(dir, "subdir.sh"), 0o755); err != nil {
			t.Fatal(err)
		}

		cmds, err := scanDirectory(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if len(cmds) != 1 {
			t.Fatalf("got %d entries, want 1", len(cmds))
		}
	})

	t.Run("error on duplicate shortnames", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "alpha.sh",
			"#!/bin/sh\n# Shortname: same\n# Description: First.\necho a\n")
		writeScript(t, dir, "beta.sh",
			"#!/bin/sh\n# Shortname: same\n# Description: Second.\necho b\n")

		_, err := scanDirectory(dir)
		if err == nil {
			t.Fatal("expected error for duplicate shortnames")
		}
		if got := err.Error(); !strings.Contains(got, "duplicate path") {
			t.Errorf("error = %q, want it to mention duplicate path", got)
		}
	})

	t.Run("error on empty results", func(t *testing.T) {
		t.Parallel()
		dir := t.TempDir()
		writeScript(t, dir, "nope.sh", "#!/bin/sh\necho no headers\n")

		_, err := scanDirectory(dir)
		if err == nil {
			t.Fatal("expected error for no valid entries")
		}
	})

	t.Run("error on nonexistent directory", func(t *testing.T) {
		t.Parallel()
		_, err := scanDirectory("/nonexistent/path")
		if err == nil {
			t.Fatal("expected error for nonexistent directory")
		}
	})
}
