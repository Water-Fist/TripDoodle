var app = new Vue({
    el: '#app',
    data: {
        ws: null,
        serverUrl: "ws://localhost:8080/chat"
    },
    mounted: function() {
        this.connectToWebsocket()
    },
    methods: {
        connectToWebsocket() {
            this.ws = new WebSocket( this.serverUrl );
            this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
        },
        onWebsocketOpen() {
            console.log("connected to chat!");
        }
    }
})