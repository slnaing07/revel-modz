revel-modz
==========

skeletons & modules for the Revel Framework

this is a work in progress and requires my patch to the `revel new` command
which allows third party skeletons

see: [revel pull #472](https://github.com/robfig/revel/pull/472)

``` Bash
go get github.com/robfig/revel
cd $GOPATH/src/github.com/robfig/revel
git remote add iassic https://github.com/iassic/revel
git pull iassic
git checkout feature/new-cmd-skeleton-arg
cd revel
go install
```


Get the dependencies
--------------

grunt:

``` Bash
sudo apt-get install nodejs npm ruby
sudo gem install sass
sudo npm install -g grunt-cli highlight.js
sudo npm install -g grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify grunt-contrib-coffee grunt-markdown grunt-contrib-sass grunt-contrib-less
```

databases: [postgres, mysql, sqlite] CouchDB

Postgres:

...

Mysql:

...

Sqlite:

...

CouchDB:

...

Installation
--------------

`go get` revel-modz

``` Bash
go get github.com/iassic/revel-modz
```


Usage
---------------
The following instructions will setup a new app from the revel-modz skeleton

``` Bash
revel new <APP_NAME> github.com/iassic/revel-modz/skeleton
cd <APP_NAME>
sh npminit.sh  (hit ctrl-c when prompted) [you will see a bunch of errors initially]
cd ..
revel run <APP_NAME>
```

now point your browser at `localhost:9000`
