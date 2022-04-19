package test

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/JohnMurray/yaii/yaai/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type snapshotOpts struct {
	generate    bool
	snapshotDir string
}

func lexerSnapshot(t *testing.T, name string, lexer *parser.YaaiLexer) error {
	opts := &snapshotOpts{
		snapshotDir: os.Getenv("YAAI_SNAPSHOT_DIR"),
		generate:    os.Getenv("YAAI_GENERATE") != "",
	}

	tl := &testListener{}
	lexer.AddErrorListener(tl)

	generate := func() ([]byte, error) {
		buf := bytes.NewBuffer([]byte{})
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				buf.WriteString("EOF\n")
				break
			}
			buf.WriteString(fmt.Sprintf("%s(%s) -> ",
				lexer.SymbolicNames[t.GetTokenType()],
				t.GetText()))
		}

		if len(tl.errors) > 0 {
			// If we have lexer errors, simply return the first
			return nil, tl.errors[0]
		}
		return buf.Bytes(), nil
	}

	return snapshot(t, name, generate, opts)
}

func snapshot(t *testing.T, name string, generate func() ([]byte, error), options *snapshotOpts) error {
	if options == nil {
		return fmt.Errorf("provided options are nil")
	}

	snapshotFilename := fmt.Sprintf("%s.snapshot", name)
	path := path.Join(options.snapshotDir, snapshotFilename)

	// generate the data that will either be used to write or test against an existing
	// snapshot file.
	data, err := generate()
	if err != nil {
		return fmt.Errorf("could not generate data to compare/write: %v", err)
	}

	if options.generate {
		return snapshotWrite(path, data)
	} else {
		return snapshotDiff(t, path, data)
	}

}

func snapshotWrite(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("issue writing snapshot to path %s: %v", path, err)
	}
	return nil
}

func snapshotDiff(t *testing.T, path string, data []byte) error {
	// check if the file exists
	fInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("snapshot file not found %s: %v", path, err)
	}

	if fInfo.IsDir() {
		return fmt.Errorf("snapshot path points to directory: %s", path)
	}

	// read in the snapshot file
	snapshot, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read snapshot file %s: %v", path, err)
	}

	// compare the daters
	if !bytes.Equal(snapshot, data) {
		printDiff(t, snapshot, data)
		return fmt.Errorf("")
	}

	return nil
}

// Print the diff between two strings
func printDiff(t *testing.T, a, b []byte) {
	diff := diffmatchpatch.New()
	diffs := diff.DiffMain(string(a), string(b), true)

	t.Log(diff.DiffPrettyText(diffs))
}
