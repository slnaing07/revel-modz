function mousePlay() {

	mightymouse.AddHandler("#btn_home","...", handle_click);

}

function handle_click(e) {
	console.log(e);
	console.log(e.data);
	alert("something was clicked");
}