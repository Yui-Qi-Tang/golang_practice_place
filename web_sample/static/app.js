$(function () {
    var ws;
    if (window.WebSocket === undefined) {
        $("#container").append("Your browser does not support WebSockets");
        return;
    } else {
        ws = initWS();
    }
    function initWS() {
        var socket = new WebSocket("wss://localhost:8081/socket/handler");
        var container = $("#container");
        socket.onopen = function() {
            //container.append("<p>Socket is open</p>");
            console.log("web Socket open!!");
        };
        socket.onmessage = function (e) {
            container.append("<p> Got some shit:" + e.data + "</p>");
        }
        socket.onclose = function () {
            container.append("<p>Socket closed</p>");
        }
        return socket;
    }
    $("#sendBtn").click(function (e) {
        e.preventDefault();
        ws.send(JSON.stringify({ Num: parseInt($("#numberfield").val()) }));
    });
});