Post-Redirect-Get
===================

This lesson introduces you to the [Post-Redirect-Get (PRG)](http://en.wikipedia.org/wiki/Post/Redirect/Get) pattern
for submitting forms.

The PRG pattern is used by web developers to prevent the resubmission of form data when refreshing the browser or opening a bookmark. For example, this can save you and your customers from the headache of an unintentional double purchase.

In PRG, a user fills out a form and submits it with a POST request.
The server receives the request, does whatever processing it needs to with
the form contents, and then uses a REDIRECT to point the user to a new page which is obtained with a GET request.
It is this REDIRECT-GET which prevents the POST from being double submitted.
Generally speaking, the page which the user is redirected to will contain some confirmation of the POST data being acceptable. In an upcoming lesson, we will explore how to deal with bad input provided by the user.

Create a new app
------------------

Let's start by making a new Revel app

``` Bash
revel new PRG github.com/iassic/revel-modz/skeleton
cd PRG
sh npminit.sh
cd ..
revel run PRG
```

You should now have a new Revel app named `PRG` running from your terminal. Check by directing your browser to `localhost:9000`.


Adding a form to Index.html
---------------------------

Open the file `PRG/app/views/App/Index.html`

Replace line: `{{template "templates/ipsum.html" .}}`

with:

``` HTML
<div class="row">
    <div class="large-6 large-centered columns">
        <div class="panel">
            <h5>My First Form</h5>
        </div>
    </div>
</div>
```

Refresh your app in the browser and you should see the changes.

Now let's add an actual form.

``` HTML
...
    <h5>My First Form</h5>
    <form action="/ipost" method="POST">
        Say something: <input type="text" name="said">
        <input type="submit">
	</form>
...
```

If you try to refresh again, you will see we temporarily broke the application. What we need is a handler for our new POST method and a new page to Redirect to.

Fixing this, or adding the POST-REDIRECT-GET loop involves several changes at this point.

Adding new handlers and routes
------------------------------

To add the new handlers, in app/controllers/app.go:

``` Go
func (c App) IndexPost(said string) revel.Result {
	fmt.Println("said:", said)
	return c.Redirect(routes.App.Result())
}

func (c App) Result() revel.Result {
	return c.Render()
}
```

To add the new view for App.Result(), in app/views/App
create a new file called `Results.html` with the following content.

``` HTML
{{set . "title" "Results"}}

{{set . "foundation_css" true}}
{{set . "foundation_js" true}}

{{template "templates/header.html" .}}
{{template "templates/menu.html" .}}

<div class="row">
    <div class="large-6 large-centered columns">
        <div class="panel">
            <h5>Results</h5>
        </div>
    </div>
</div>

{{template "templates/links.html" .}}
{{template "templates/footer.html" .}}

```


To add the new routes, in app/conf/routes:

add the following lines below the first route

```
POST 	/ipost 			App.IndexPost
GET 	/result 		App.Result
```









The final version of each file:
---------------------------------

PRG/app/controllers/app.go

``` Go
package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"PRG/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) IndexPost(said string) revel.Result {
	fmt.Println("said:", said)
	return c.Redirect(routes.App.Result())
}

func (c App) Result() revel.Result {
	return c.Render()
}
```

PRG/app/views/App/Index.html:

``` HTML
{{set . "title" "Results"}}

{{set . "foundation_css" true}}
{{set . "foundation_js" true}}

{{template "templates/header.html" .}}
{{template "templates/menu.html" .}}

<div class="row">
    <div class="large-6 large-centered columns">
        <div class="panel">
            <h5>Results</h5>
        </div>
    </div>
</div>

{{template "templates/links.html" .}}
{{template "templates/footer.html" .}}

```

PRG/app/views/App/Result.html:

``` HTML
{{set . "title" "Results"}}

{{set . "foundation_css" true}}
{{set . "foundation_js" true}}

{{template "templates/header.html" .}}
{{template "templates/menu.html" .}}

<div class="row">
    <div class="large-6 large-centered columns">
        <div class="panel">
            <h5>Results</h5>
        </div>
    </div>
</div>

{{template "templates/links.html" .}}
{{template "templates/footer.html" .}}

```

PRG/app/conf/routes:

```
# Routes
# This file defines all application routes (Higher priority routes first)
# ~~~~

module:testrunner

GET     /                                       App.Index

POST    /ipost									App.IndexPost

GET 	/result									App.Result

# github.com/iassic/revel-modz/module/assets   (/ipa)
module:ipa



# Don't ignore favicon requests
GET     /favicon.ico                            Static.Serve("public/img","favicon.png")

# Map static resources from the /app/public folder to the /public path
GET     /public/*filepath                       Static.Serve("public")

# 404 Catch all
*       /:controller/:action                    404

# Catch all
# *       /:controller/:action                    :controller.:action
```
