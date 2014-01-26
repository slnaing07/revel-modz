var myfiles = {
  "foundation_js": [
  "public/js/foundation-5.0.3/foundation.abide.js",
  "public/js/foundation-5.0.3/foundation.accordion.js",
  "public/js/foundation-5.0.3/foundation.alert.js",
  "public/js/foundation-5.0.3/foundation.clearing.js",
  "public/js/foundation-5.0.3/foundation.dropdown.js",
  "public/js/foundation-5.0.3/foundation.interchange.js",
  "public/js/foundation-5.0.3/foundation.joyride.js",
  "public/js/foundation-5.0.3/foundation.js",
  "public/js/foundation-5.0.3/foundation.magellan.js",
  "public/js/foundation-5.0.3/foundation.offcanvas.js",
  "public/js/foundation-5.0.3/foundation.orbit.js",
  "public/js/foundation-5.0.3/foundation.reveal.js",
  "public/js/foundation-5.0.3/foundation.tab.js",
  "public/js/foundation-5.0.3/foundation.tooltip.js",
  "public/js/foundation-5.0.3/foundation.topbar.js"
  ],
  "bootstrap_js": [
  "public/js/bootstrap-3.0.3/transition.js",
  "public/js/bootstrap-3.0.3/alert.js",
  "public/js/bootstrap-3.0.3/button.js",
  "public/js/bootstrap-3.0.3/carousel.js",
  "public/js/bootstrap-3.0.3/collapse.js",
  "public/js/bootstrap-3.0.3/dropdown.js",
  "public/js/bootstrap-3.0.3/modal.js",
  "public/js/bootstrap-3.0.3/tooltip.js",
  "public/js/bootstrap-3.0.3/popover.js",
  "public/js/bootstrap-3.0.3/scrollspy.js",
  "public/js/bootstrap-3.0.3/tab.js",
  "public/js/bootstrap-3.0.3/affix.js"
  ],

  "myapp_js": [
    "public/js/myapp/myapp_sample1.js",
    "public/js/myapp/myapp_sample2.js"
  ],

  "myapp_md": [
    "app/views/markdown/sample1.md",
    "app/views/markdown/sample2.md"
  ],

  "myapp_coffee": [
    "public/coffee/myapp/myapp_sample1.coffee",
    "public/coffee/myapp/myapp_sample2.coffee"
  ]
}

module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    
    jshint: {
      options: {
        jshintrc: 'public/js/.jshintrc'
      },
      gruntfile: {
        src: 'Gruntfile.js'
      },
      src: {
        src: ['public/js/**/*.js']
      }
    },

    concat: {
      options: {
        separator: ';'
      },

      foundation_js: {
        files: {
          'public/js/foundation.js': myfiles.foundation_js,
        },
      },
      bootstrap_js: {
        files: {
          'public/js/bootstrap.js': myfiles.bootstrap_js,
        },
      },
      myapp_js: {
        files: {
          'public/js/<%= pkg.name %>.js': myfiles.myapp_js,
        },
      },
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      foundation_js: {
        src: 'public/js/foundation.js',
        dest: 'public/js/fountation.min.js'
      },
      bootstrap_js: {
        src: 'public/js/bootstrap.js',
        dest: 'public/js/bootstrap.min.js'
      },
      myapp_js: {
        src: 'public/js/<%= pkg.name %>.js',
        dest: 'public/js/<%= pkg.name %>.min.js'
      }
    }

    // sass: {
    //   options: {
    //     includePaths: ['bower_components/foundation/scss']
    //   },
    //   dist: {
    //     options: {
    //       outputStyle: 'compressed'
    //     },
    //     files: {
    //       'css/app.css': 'scss/app.scss'
    //     }        
    //   }
    // },

    // watch: {
    //   gruntfile: {
    //     files: 'Gruntfile.js',
    //     tasks: ['jshint:gruntfile'],
    //   },

    //   sass: {
    //     files: 'scss/**/*.scss',
    //     tasks: ['sass']
    //   }
    // }


  });

  grunt.event.on('watch', function(action, filepath, target) {
    grunt.log.writeln(target + ': ' + filepath + ' has ' + action);
  });


  // Load the plugin that provides the tasks.
  // grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  // grunt.loadNpmTasks('grunt-sass');
  // grunt.loadNpmTasks('grunt-contrib-qunit');
  // grunt.loadNpmTasks('grunt-contrib-watch');

  // Register tasks.
  grunt.registerTask('default', ['concat','uglify']);
  // grunt.registerTask('test', ['jshint', 'qunit']);

};