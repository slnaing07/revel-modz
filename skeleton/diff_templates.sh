#!/bin/bash

files=(
	# init.sh
	# Gruntfile.js

	conf/app.conf

	app/controllers/admin.go
	app/controllers/app.go
	app/controllers/user.go

	app/views/templates/header.html
)

for file in ${files[@]}
do
	echo "${file}"
	echo "------------------------"
	colordiff ${file}.template ../sample/${file} 
	echo
	echo
	echo
done
