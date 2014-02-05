revel-modz/modules/assets
=================================

A collection of useful libraries and other public assets

Installation
-----------------

`go get github.com/iassic/revel-modz`


What's inside
-----------------

Frameworks

- bootstrap-3.0.3   (JS,CSS)
- foundation-5.0.3  (JS,CSS)

CSS

- normalize.css
- jquery-ui.min.css

Fonts

- whhq: font based icons from [webhostinghub.com](http://www.webhostinghub.com/glyphs/)

JS

- head.load.min.js
- jquery-1.10.2. {js,min.js,min.map}
- jquery-2.0.3. {js,min.js,min.map}
- jquery-ui.min.js
- modernizer.js

- fastclick.js
- higlight.pack.js
- placeholder.js
- move.js
- keypress-1.0.9.min.js
- [mightmouse.js](https://github.com/verdverm/mightymouse-js)

- d3.js


Usage
------------------

add the following line to your `app.conf`

```
module.ipa     = github.com/iassic/revel-modz/modules/assets
```

add the following line to your `routes` to import under the path `/ipa/*`

```
module:ipa
```

add the following line to your `routes` to import under the path `/<path>/ipa/*`

```
* /<path>/  module:ipa
```

