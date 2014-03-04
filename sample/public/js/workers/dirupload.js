self.addEventListener('message', function(e) {
  var file = JSON.parse(e.data);

  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/files/upload", false);
  xhr.setRequestHeader('X-CSRF-Token',file.Csrf );
  xhr.overrideMimeType('text/plain; charset="x-user-defined-binary"');

  xhr.send(JSON.stringify(file));
  self.postMessage(JSON.stringify(file.Size));

  // fileUpload(file,file.Csrf);
}, false);

// not used right now
function fileUpload(file) {
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/files/upload", false);
  xhr.overrideMimeType('text/plain; charset="x-user-defined-binary"');
  xhr.setRequestHeader('X-CSRF-Token',CSRF );
  xhr.send(JSON.stringify(file));
  self.postMessage(JSON.stringify(file.Size));
}
