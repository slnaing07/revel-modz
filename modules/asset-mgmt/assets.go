package app

import (
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/revel"
)

func init() {
	compilers := []AssetCompiler{
		AssetCompiler{
			Name: "Foundation JS Folder",
			Exec: "uglifyjs",
			Path: "public/js/foundation",
		},
		AssetCompiler{
			Name: "Foundation SASS Folder",
			Exec: "compass...",
			Path: "public/sass/myapp",
		},
		AssetCompiler{
			Name: "MyApp JS Folder",
			Exec: "uglifyjs...",
			Path: "public/js/myapp",
		},
		AssetCompiler{
			Name: "MyApp SASS Folder",
			Exec: "compass...",
			Path: "public/sass/myapp",
		},
		AssetCompiler{
			Name: "MyApp Coffee Folder",
			Exec: "coffee...",
			Path: "public/coffeescript",
		},
		AssetCompiler{
			Name: "MyApp Markdown Folder",
			Exec: "blackfriday...",
			Path: "views/markdown",
		},
		AssetCompiler{
			Name: "MyApp Template Folder",
			Exec: "mustache...",
			Path: "views/templates",
		},
	}

	revel.OnAppStart(func() {
		for _, AC := range compilers {
			revel.MainWatcher.Listen(AC, AC.Path)
		}
	})

}

type AssetCompiler struct {
	Name string
	Exec string
	Path string
}

func (AC *AssetCompiler) Refresh() *revel.Error {
	// It's start-up or a file changed.  Re-compile...
	println(AC.Name)

	// Revel's working directory, should be the app root
	orig, _ := os.Getwd()

	// change to the assets directory
	// TODO change to os.FilePathSeparater, or whatever it's called
	os.Chdir(orig + AC.Path)

	// compile js files
	for _, p := range c.Assets {
		cmd_str := p.Script
		_, err := exec.Command(cmd_str).Output()
		if err != nil {
			glog.Errorln("Failed to compile", p.Path, "JS\n", err)
			return nil
		}
	}
	os.Chdir(orig)
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
