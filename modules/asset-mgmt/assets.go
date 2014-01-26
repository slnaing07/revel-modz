package assetmgmt

import (
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/revel"
)

type AssetCompiler struct {
	Name  string
	Path  string
	Grunt string
}

func (AC AssetCompiler) Refresh() *revel.Error {
	// It's start-up or a file changed.  Re-compile...
	revel.INFO.Println("Compiling: ", AC.Name)

	os.Chdir(revel.BasePath)

	out, err := exec.Command("grunt", AC.Grunt).Output()
	if err != nil {
		revel.ERROR.Println("Failed to compile", AC.Path, err)
		return nil
	}
	revel.TRACE.Println(string(out))

	os.Chdir(revel.SourcePath)
	return nil
}

func (c *AssetCompiler) WatchDir(info os.FileInfo) bool {
	// Watch all directories, except the ones starting with a dot.
	return !strings.HasPrefix(info.Name(), ".")
}

func (c *AssetCompiler) WatchFile(basename string) bool {
	// Watch all files, except the ones starting with a dot.
	return !strings.HasPrefix(basename, ".")
}
