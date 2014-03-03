function initComm() {

	$("#send_msg_btn").click(function(){
		websock_comm.Send($('#comm_input').val());
	})

	websock_comm.AddHandler("error",error_handler)
	websock_comm.AddHandler("ack",ack_handler)
	websock_comm.Connect(window.location.host, "/comm");
}

function error_handler(msg) {
	console.log("Recv'd ERROR: " + msg)
}

function ack_handler(msg) {
	console.log("Recv'd ACK: " + msg)
}
