package app

import (
	AM "github.com/iassic/revel-modz/modules/asset-mgmt"
	"github.com/robfig/revel"
)

// in your app initialization..
func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	revel.OnAppStart(func() {
		for _, AC := range compilers {
			revel.INFO.Printf("Listening: %q\n", AC.Path)
			revel.MainWatcher.Listen(AC, revel.AppName+"/"+AC.Path)
		}
	})

}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var compilers = []AM.AssetCompiler{
	AM.AssetCompiler{
		Name:  "Gruntfile.js",
		Path:  "Gruntfile.js",
		Grunt: "default",
	},

	AM.AssetCompiler{
		Name:  "Foundation JS Folder",
		Path:  "public/js/foundation-5.0.3",
		Grunt: "uglify:foundation_js",
	},
	AM.AssetCompiler{
		Name:  "Foundation SASS Folder",
		Path:  "public/sass/foundation-5.0.3",
		Grunt: "sass:foundation_sass",
	},

	AM.AssetCompiler{
		Name:  "Bootstrap JS Folder",
		Path:  "public/js/bootstrap-3.0.3",
		Grunt: "uglify:bootstrap_js",
	},
	AM.AssetCompiler{
		Name:  "Bootstrap LESS Folder",
		Path:  "public/less/bootstrap-3.0.3",
		Grunt: "less:bootstrap_less",
	},

	AM.AssetCompiler{
		Name:  "MyApp Coffee Folder",
		Path:  "public/coffeescript",
		Grunt: "coffee:myapp_coffee",
	},
	AM.AssetCompiler{
		Name:  "MyApp JS Folder",
		Path:  "public/js/myapp",
		Grunt: "uglify:myapp_js",
	},
	AM.AssetCompiler{
		Name:  "MyApp SASS Folder",
		Path:  "public/sass/myapp",
		Grunt: "sass:myapp_sass",
	},
	AM.AssetCompiler{
		Name:  "MyApp LESS Folder",
		Path:  "public/less/myapp",
		Grunt: "less:myapp_less",
	},
	AM.AssetCompiler{
		Name:  "MyApp Markdown Folder",
		Path:  "views/markdown",
		Grunt: "markdown:myapp_md",
	},
}
