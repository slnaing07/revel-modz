var files = {
  "foundation_js": [
  "app/assets/js/foundation/foundation.abide.js",
  "app/assets/js/foundation/foundation.accordion.js",
  "app/assets/js/foundation/foundation.alert.js",
  "app/assets/js/foundation/foundation.clearing.js",
  "app/assets/js/foundation/foundation.dropdown.js",
  "app/assets/js/foundation/foundation.interchange.js",
  "app/assets/js/foundation/foundation.joyride.js",
  "app/assets/js/foundation/foundation.js",
  "app/assets/js/foundation/foundation.magellan.js",
  "app/assets/js/foundation/foundation.offcanvas.js",
  "app/assets/js/foundation/foundation.orbit.js",
  "app/assets/js/foundation/foundation.reveal.js",
  "app/assets/js/foundation/foundation.tab.js",
  "app/assets/js/foundation/foundation.tooltip.js",
  "app/assets/js/foundation/foundation.topbar.js"
  ],
  
  "soe_smple_js": [
    "app/assets/js/soe_smple/soe_smple_util.js"
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
      soe_smple_js: {
        files: {
          'public/js/soe_smple.min.js': files.soe_smple_js
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
      soe_smple_css: {
        options: {
          style: 'expanded',
          lineNumbers: true
        },
        files: {
          'public/css/soe_smple.css': 'app/assets/sass/soe_smple.sass'
        }        
      }
    },


  });

  grunt.event.on('watch', function(action, filepath, target) {
    grunt.log.writeln(target + ': ' + filepath + ' has ' + action);
  });


  // Load the plugin that provides the tasks.
  grunt.loadNpmTasks('grunt-contrib-jshint');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-sass');

  // Register tasks.
  grunt.registerTask('default', ['uglify','sass']);
  grunt.registerTask('test', ['jshint', 'qunit']);

};