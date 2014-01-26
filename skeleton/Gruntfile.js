

module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    
    jshint: {
      options: {
        jshintrc: 'js/.jshintrc'
      },
      gruntfile: {
        src: 'Gruntfile.js'
      },
      src: {
        src: ['js/*.js']
      },
      test: {
        src: ['js/tests/unit/*.js']
      }
    },

    concat: {
      options: {
        separator: ';'
      },

      foundation_js: {
        files: {
          'public/js/foundation.js': pkg.foundation_js,
        },
      },
      bootstrap_js: {
        files: {
          'public/js/bootstrap.js': pkg.bootstrap_js,
        },
      },
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      build: {
        src: 'src/<%= pkg.name %>.js',
        dest: 'build/<%= pkg.name %>.min.js'
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

    // qunit: {
    //   files: ['test/**/*.html']
    // },
    // jshint: {
    //   files: ['Gruntfile.js', 'src/**/*.js', 'test/**/*.js'],
    //   options: {
    //     // options here to override JSHint defaults
    //     globals: {
    //       jQuery: true,
    //       console: true,
    //       module: true,
    //       document: true
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
  // grunt.loadNpmTasks('grunt-sass');
  // grunt.loadNpmTasks('grunt-contrib-uglify');
  // grunt.loadNpmTasks('grunt-contrib-jshint');
  // grunt.loadNpmTasks('grunt-contrib-qunit');
  // grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-compass');

  // Register tasks.
  grunt.registerTask('test', ['concat','compass']);

  // grunt.registerTask('test', ['jshint', 'qunit']);
  // grunt.registerTask('default', ['jshint', 'qunit', 'concat', 'uglify']);

};