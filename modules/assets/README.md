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
- slickgrid         (here because it has a lot of stuff)
- fancytree-2.0.0-5 (here because it has a lot of stuff)

CSS

- normalize.css
- jquery-ui.min.css

Fonts

- whhq: font based icons from [webhostinghub.com](http://www.webhostinghub.com/glyphs/)

JS

- head.load.min.js
- modernizer.js
- jquery-1.10.2. {js,min.js,min.map}
- jquery-2.0.3. {js,min.js,min.map}

JQuery plugins

- jquery-ui.min.js
- jquery-autocomplete.js
- jquery-cookie.js
- jquery-event.drag-2.2.js
- jquery-event.drop-2.2.js
- jquery-dynatree.min.js
- [jquery-fancytree.js](https://github.com/mar10/fancytree/)
- jquery-sparkline.min.js

Other JS libraries

- fastclick.js
- higlight.pack.js
- placeholder.js
- move.js
- keypress-1.0.9.min.js
- [mightmouse.js](https://github.com/verdverm/mightymouse-js)

- d3.js
- dygraph.min.js
- LaTeXMathML.js
- mathquill.min.js


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

