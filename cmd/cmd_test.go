package cmd_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twistingmercury/scaffolder/cmd"
	"github.com/twistingmercury/scaffolder/conf"
	"io"
	"strings"
	"testing"
)

func TestRootCmd(t *testing.T) {
	cmd.Initialize()
	b := bytes.NewBufferString("")
	cmd.RootCmd().SetOut(b)
	cmd.RootCmd().SetArgs([]string{""})
	cmd.Execute()
	require.NotNil(t, cmd.RootCmd())
	assert.Contains(t, cmd.RootCmd().UseLine(), cmd.Usage)
}

func TestRootCmdHelpFlag(t *testing.T) {
	rCmd := cmd.NewRootCmd()
	require.NotNil(t, rCmd)
	b := bytes.NewBufferString("")
	rCmd.SetOut(b)
	rCmd.SetArgs([]string{"--help"})
	err := rCmd.Execute()
	require.NoError(t, err)

	out, err := io.ReadAll(b)
	require.NoError(t, err)
	assert.Greater(t, len(out), 0)
}

func TestVersionCmd(t *testing.T) {
	vcmd := cmd.NewVersionCmd()
	require.NotNil(t, vcmd)
	b := bytes.NewBufferString("")
	vcmd.SetOut(b)
	err := vcmd.Execute()
	require.NoError(t, err)
	bits, err := io.ReadAll(b)
	actual := strings.TrimSpace(string(bits))
	expected := "scaffolder version " + conf.BuildVersion
	assert.Equal(t, expected, actual)
}
