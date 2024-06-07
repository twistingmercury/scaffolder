package cmd_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/twistingmercury/scaffolder/cmd"
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
	expected := "scaffolder version: n/a, build date: n/a"
	assert.Equal(t, expected, actual)
}

func TestNewTemplateInfo(t *testing.T) {
	const (
		gitPath     = `https://github.com/twistingmercury/fake-template.git`
		moduleName  = "unit/test"
		binName     = "app"
		rootDir     = "appDir"
		vendorName  = "Wallyworld"
		description = "paradise"
	)

	var (
		ti  cmd.TemplateInfo
		err error
	)

	t.Run("missing module name", func(t *testing.T) {
		_, err = cmd.NewTemplateInfo(gitPath, rootDir, "", binName, vendorName, description)
		assert.Error(t, err)
	})
	t.Run("missing bin name", func(t *testing.T) {
		_, err = cmd.NewTemplateInfo(gitPath, rootDir, moduleName, "", vendorName, description)
		assert.Error(t, err)
	})
	t.Run("missing git path", func(t *testing.T) {
		_, err = cmd.NewTemplateInfo("", rootDir, moduleName, binName, vendorName, description)
		assert.Error(t, err)
	})
	t.Run("missing root dir", func(t *testing.T) {
		_, err = cmd.NewTemplateInfo(gitPath, "", moduleName, binName, vendorName, description)
		assert.Error(t, err)
	})
	t.Run("missing vendor and description", func(t *testing.T) {
		ti, err = cmd.NewTemplateInfo(gitPath, rootDir, moduleName, binName, "", "")
		assert.NoError(t, err)
		assert.NotNil(t, ti)
		assert.Equal(t, "TODO: provide a vendor name", ti.VendorName.Value)
		assert.Equal(t, fmt.Sprintf("TODO: provide a description for %s", binName), ti.Description.Value)
	})
	t.Run("all values provided", func(t *testing.T) {
		ti, err = cmd.NewTemplateInfo(gitPath, rootDir, moduleName, binName, vendorName, description)
		assert.NoError(t, err)
		assert.NotNil(t, ti)
		assert.Equal(t, moduleName, ti.ModuleName.Value)
		assert.Equal(t, binName, ti.BinName.Value)
		assert.Equal(t, rootDir, ti.RootDir)
		assert.Equal(t, gitPath, ti.GitPath)
		assert.Equal(t, vendorName, ti.VendorName.Value)
		assert.Equal(t, description, ti.Description.Value)

		assert.Equal(t, "MODULE_NAME|{{module_name}}", ti.ModuleName.Regex().String())
		assert.Equal(t, "BIN_NAME|{{bin_name}}", ti.BinName.Regex().String())
		assert.Equal(t, "IMG_VENDOR_NAME|{{vendor_name}}", ti.VendorName.Regex().String())
		assert.Equal(t, "IMG_DESCRIPTION|{{description}}", ti.Description.Regex().String())
	})
}
