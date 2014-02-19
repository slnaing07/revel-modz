#!/bin/bash

# npminit.sh

npm init
npm install grunt
npm install grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify grunt-contrib-coffee grunt-markdown grunt-contrib-sass grunt-contrib-less

appname="soe_smple"

mv app/assets/js/myapp app/assets/js/$appname
mv app/assets/js/$appname/myapp_util.js app/assets/js/$appname/${appname}_util.js

mv app/assets/sass/myapp app/assets/sass/$appname
mv app/assets/sass/myapp.sass app/assets/sass/$appname.sass
mv app/assets/sass/$appname/_myapp_user.sass app/assets/sass/$appname/_${appname}_user.sass
mv app/assets/sass/$appname/_myapp_common.sass app/assets/sass/$appname/_${appname}_common.sass
