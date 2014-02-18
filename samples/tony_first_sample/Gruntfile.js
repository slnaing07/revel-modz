var files = {
  "foundation_js": [
  "app/assets/js/foundation-5.0.3/foundation.abide.js",
  "app/assets/js/foundation-5.0.3/foundation.accordion.js",
  "app/assets/js/foundation-5.0.3/foundation.alert.js",
  "app/assets/js/foundation-5.0.3/foundation.clearing.js",
  "app/assets/js/foundation-5.0.3/foundation.dropdown.js",
  "app/assets/js/foundation-5.0.3/foundation.interchange.js",
  "app/assets/js/foundation-5.0.3/foundation.joyride.js",
  "app/assets/js/foundation-5.0.3/foundation.js",
  "app/assets/js/foundation-5.0.3/foundation.magellan.js",
  "app/assets/js/foundation-5.0.3/foundation.offcanvas.js",
  "app/assets/js/foundation-5.0.3/foundation.orbit.js",
  "app/assets/js/foundation-5.0.3/foundation.reveal.js",
  "app/assets/js/foundation-5.0.3/foundation.tab.js",
  "app/assets/js/foundation-5.0.3/foundation.tooltip.js",
  "app/assets/js/foundation-5.0.3/foundation.topbar.js"
  ],
  
  "tony_first_sample_js": [
    "app/assets/js/tony_first_sample/tony_first_sample_util.js"
  ]
}

module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    
    jshint: {
      options: {
        jshintrc: 'app/assets/js/.jshintrc'
      },
      gruntfile: {
        src: 'Gruntfile.js'
      },
      src: {
        src: ['app/assets/js/**/*.js']
      }
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n',
        sourceMap: true,
        report: 'min'
      },
      foundation_js: {
        files: {
          'public/js/foundation-custom.min.js': files.foundation_js
        }
      },
      tony_first_sample_js: {
        files: {
          'public/js/tony_first_sample.min.js': files.tony_first_sample_js
        }
      }
    },

    sass: {
      foundation_css: {
        options: {
          style: 'compressed'
        },
        files: {
          'public/css/foundation-custom.min.css': 'app/assets/sass/foundation_custom.scss'
        }        
      }
      ,
      tony_first_sample_css: {
        options: {
          style: 'expanded',
          lineNumbers: true
        },
        files: {
          'public/css/tony_first_sample.css': 'app/assets/sass/tony_first_sample.sass'
        }        
      }
    },


  });

  grunt.event.on('watch', function(action, filepath, target) {
    grunt.log.writeln(target + ': ' + filepath + ' has ' + action);
  });


  // Load the plugin that provides the tasks.
  // grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-sass');

  // Register tasks.
  grunt.registerTask('default', ['uglify','sass']);
  grunt.registerTask('test', ['jshint', 'qunit']);

};