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

Now we need to add some imports for our new handlers:

`"fmt"` for the `Println` function and `"PRG/app/routes"` for the `routes.App.Result()`
inside of `c.Redirect(...)`

Change the line `import "github.com/revel/revel"`

``` Go
import (
"fmt"

"github.com/revel/revel"

"PRG/app/routes"
)
```


To add the new view for App.Result(), in app/views/App
create a new file called `Result.html` with the following content.

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


To add the new routes, in PRG/conf/routes:

add the following lines below the first route

```
POST 	/ipost 			App.IndexPost
GET 	/result 		App.Result
```


And now we should be good, try refreshing your browser.
No need to restart the webserver!


Getting the message to show up in Result.html
-------------------------------------------

Now we are going to use Revel's template system and Flash mechanism in order to have the text entered and submitted show up in the Result page.

The first item of business is to add a template for displaying the message. Templates in Revel are derived from Go which is basically a text replacement engine. In Revel's case, we are using the 'html/template' package which html-escapes the substituted input.
This prevents the user from entering malicious input such as `<script src="http://badguy.com/mynastyscript.js">`

In `Result.html`, add the following line below `<h5>Results</h5>`:

```
you said: {{.flash.message}}
```

and in `app.go`, replace `fmt.Println(...)` in `IndexPost()` with

``` Go
c.Flash.Out["message"] = said
```

Notice that `.message` in the template matches `"message"` in the POST handler.

Once you have made these changes, try loading the index page and submitting a message. You should see that message displayed on the results page.

The final version of each file:
---------------------------------

PRG/app/controllers/app.go:

``` Go
package controllers

import (
"github.com/revel/revel"
"PRG/app/routes"
)

type App struct {
*revel.Controller
}

func (c App) Index() revel.Result {
return c.Render()
}

func (c App) IndexPost(said string) revel.Result {
c.Flash.Out["message"] = said
return c.Redirect(routes.App.Result())
}

func (c App) Result() revel.Result {
return c.Render()
}
```

PRG/app/views/App/Index.html:

``` HTML
{{set . "title" "Index"}}

{{set . "foundation_css" true}}
{{set . "foundation_js" true}}

{{template "templates/header.html" .}}
{{template "templates/menu.html" .}}

<div class="row">
    <div class="large-6 large-centered columns">
        <div class="panel">
            <h5>My First Form</h5>
            <form action="/ipost" method="POST">
              Say something: <input type="text" name="said">
              <input type="submit">
          </form>
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
            You said: {{.flash.message}}
        </div>
    </div>
</div>

{{template "templates/links.html" .}}
{{template "templates/footer.html" .}}
```

PRG/conf/routes:

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
