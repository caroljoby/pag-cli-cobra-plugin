package tests

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yathendra/cobrapagcliplugin/cmd"
)

func TestStatusCmd_ShowsSession(t *testing.T) {
	buf := &bytes.Buffer{}
	root := cmd.NewPagCmd(buf)
	root.SetArgs([]string{"session", "status", ""})
	err := root.Execute()
	if err != nil {
		t.Fatal(err)
	}
	got := buf.String()
	if !strings.Contains(got, "Status of session 123") {
		t.Fatalf("unexpected output: %q", got)
	}
}
