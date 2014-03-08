revel-modz
==========

Modules, samples, and a skeleton for the Revel Framework


Get the dependencies
--------------

grunt:

``` Bash
sudo apt-get install nodejs npm ruby
sudo gem install sass
sudo npm install -g grunt-cli highlight.js
sudo npm install -g grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify grunt-contrib-coffee grunt-markdown grunt-contrib-sass grunt-contrib-less
```

databases:

one of [postgres, mysql, sqlite]


Installation
--------------

`go get -u` revel-modz

``` Bash
go get -u github.com/iassic/revel-modz
```


Usage
---------------
The following instructions will setup a new app from the revel-modz skeleton

``` Bash
revel new <APP_NAME> github.com/iassic/revel-modz/skeleton
cd <APP_NAME>
bash init.sh  (hit ctrl-c when prompted) [you will see a bunch of errors initially]
cd ..
```

create a new database and update the db.spec in app.conf 
add the following environment variables to your `.profile` or `.bashrc`
```
export DB_DEV_USER='username'
export DB_DEV_PASS='userpass'
export DB_DEV_NAME='databasename'
export DB_PROD_USER='username'
export DB_PROD_PASS='userpass'
export DB_PROD_NAME='databasename'
```



```
revel run <APP_NAME>
```

now point your browser at `localhost:9000`


Features
----------------

Front-end:

- Foundation 5.1.1
- Headjs for asynchronous loading of assets
- Many JS/CSS goodies in revel-modz/modules/assets
- Templated includes for per page assets
- A `appendjs` template function for inserting JS code

Back-end:

- JS/SASS app resources initialized in app/assets
- Hot Code watch and recompile of app/assets with Grunt
- ORM with github.com/jinzhu/gorm

Security:

- User Authentication
- CSRF protection
- `X-Frame-Options` `X-XXS-Protection` and `X-Content-Type-Options` headers


Samples
----------------

The skeleton includes everything

more samples coming soon
