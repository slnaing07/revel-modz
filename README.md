revel-modz
==========

skeletons & modules for the Revel Framework

this is a work in progress and requires my patch to the `revel new` command
which allows third party skeletons

see: [revel pull #472](github.com/robfig/revel/pull/472)

``` Bash
go get github.com/robfig/revel
cd github.com/robfig/revel
git remote add iassic github.com/iassic/revel
git checkout feature/new-cmd-skeleton-arg
```

Installation
--------------

The following instructions will setup a new app with
a skeleton and a grunt file for asset management

``` Bash
go get github.com/iassic/revel-modz
cd github.com/iassic/revel-modz/modules/asset-mgmt/scripts
sh install_dependencies.sh
```


Usage
---------------

``` Bash
revel new <APP_NAME> github.com/iassic/revel-modz/skeleton
cd <APP_NAME>
sh npminit.sh
revel run <APP_NAME>
```
