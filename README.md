revel-modz
==========

skeletons & modules for the Revel Framework

this is a work in progress and requires my patch to the `revel new` command
which allows third party skeletons

see: github.com/robfig/revel/pull/472


Installation
--------------

The following instructions will setup a new app with
a skeleton and a grunt file for asset management

`go get github.com/iassic/revel-modz`

`cd $GOPATH/github.com/iassic/revel-modz/modules/asset-mgmt/scripts`

`sh install_dependencies.sh`


Usage
---------------

`revel new <APP_NAME> github.com/iassic/revel-modz/skeleton`

`cd <APP_NAME>`

`sh npminit.sh`
