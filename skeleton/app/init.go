package app

import (
	"github.com/iassic/revel-modz/modules/grunt"
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

var compilers = []grunt.GruntCompiler{
	grunt.GruntCompiler{Name: "Gruntfile.js", Path: "Gruntfile.js", Grunt: "default"},

	grunt.GruntCompiler{Name: "Foundation JS", Path: "app/assets/js/foundation-5.0.3", Grunt: "uglify:foundation_js"},
	grunt.GruntCompiler{Name: "Foundation SASS", Path: "app/assets/sass/foundation-5.0.3", Grunt: "sass:foundation_css"},

	grunt.GruntCompiler{Name: "MyApp JS", Path: "app/assets/js/myapp", Grunt: "uglify:myapp_js"},
	grunt.GruntCompiler{Name: "MyApp SASS", Path: "app/assets/sass/myapp", Grunt: "sass:myapp_sass"},
}
