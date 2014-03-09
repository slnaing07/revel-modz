function initAnalyticsView() {

    // function error_handler(msg) {
    //     console.log("Recv'd ERROR: " + msg)
    // }

    // function ack_handler(msg) {
    //     console.log("Recv'd ACK: " + msg)
    // }

    // $("#send_msg_btn").click(function() {
    //     websock_comm.Send($('#comm_input').val());
    // })

    // websock_comm.AddHandler("error", error_handler)
    // websock_comm.AddHandler("ack", ack_handler)
    // websock_comm.Connect(window.location.host, "/comm");



    // setup filter form submit button
    $("#analytics-pagereqs-filter-button").on("click", handle_filter_button_click);
}

function handle_filter_button_click(e) {
	// dev debug printing
    console.log("filter button was clicked");
    console.log(e);


    // post filter values to server
    var group = $("#analytics-pagereqs-field-group").val();
    var id = $("#analytics-pagereqs-field-id").val();
    console.log("group", group)
    console.log("id", id)


    // receive results from server
    var data = {
        message: "hello world",
    };


    // clear the current table results



    // render results in panel

	    // for each element of the results {
    var text = "my <%message%> template."
    var template = Hogan.compile(text, {
        delimiters: '<% %>'
    });
    var output = template.render(data);

    console.log(output);

    // instead of logging, append to 'table'
    // $("#analytics-pagereqs-results").append(output)

    // }
}

// {{range $index, $result := .results}}
// <div class="row">
//     <div class="large-12 columns">
//         result row
//         <div class="row">
//             <div class="small-1 columns">
//                 {{$index}}
//             </div>
//             <div class="small-1 columns">
//                 {{$result.Method}}
//             </div>
//             <div class="small-2 columns">
//                 {{$result.Time}}
//             </div>
//             <div class="small-2 columns">
//                 Time
//             </div>
//             <div class="small-2 columns">
//                 Time
//             </div>
//             <div class="small-2 columns">
//                 Time
//             </div>
//             <div class="small-2 columns">
//                 Time
//             </div>
//         </div>
//     </div>
// </div>
// {{end}}