#!/bin/bash

# npminit.sh

npm init
npm install grunt grunt-contrib-watch
npm install grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify 
npm install grunt-contrib-coffee grunt-markdown grunt-sass grunt-contrib-less
grunt
