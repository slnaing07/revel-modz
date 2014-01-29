#!/bin/bash
#
# tested in Ubuntu 13.10

# install 3rd party tools and frameworks
sudo apt-get install npm
sudo npm install -g grunt-cli highlight.js

sudo npm install -g grunt-contrib-watch grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify grunt-contrib-coffee grunt-markdown grunt-sass grunt-contrib-less


#
#  once project is created with 'revel new'
#  need to run:
#		>npm init
# 		>npm install grunt
#
