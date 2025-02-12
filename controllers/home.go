package controllers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := `
  <!DOCTYPE html>
  <html>
  <head>
    <title>Upload CSV</title>
  </head>
  <body>
    <h1>Upload CSV File</h1>
    <form enctype="multipart/form-data">
      <input type="file" id="fileInput" name="file" accept=".csv">
      <br><br>
    </form>
    <button id="submit" onclick="submitForm()">Upload</button>

    <h1>WebSocket Client</h1>
    <div>
        <input id="messageInput" type="text" placeholder="Enter message" />
        <button id="sendBtn">Send</button>
    </div>
    <div id="messages" style="margin-top: 20px;">
        <h3>Messages:</h3>
    </div>
  </body>
    <script>
        const token= sessionStorage.getItem('token');
        if(token==null || token=='') {
          token= randomString(10)
          sessionStorage.setItem('token', token)
        }
        const ws = new WebSocket('ws://localhost:1900/ws?key='+token); // Replace with your server URL

        const messageInput = document.getElementById('messageInput');
        const sendBtn = document.getElementById('sendBtn');
        const messagesDiv = document.getElementById('messages');

        // Event: Connection opened
        ws.onopen = () => {
            console.log('Connected to WebSocket server');
            addMessage('Connected to server');
        };

        // Event: Message received
        ws.onmessage = (event) => {
            console.log('Message received:', event.data);
            addMessage('Server: '+event.data);
        };

        // Event: Connection closed
        ws.onclose = () => {
            console.log('Disconnected from server');
            addMessage('Disconnected from server');
        };

        // Event: Error occurred
        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
            addMessage('Error occurred. Check console for details.');
        };

        // Send message on button click
        sendBtn.addEventListener('click', () => {
            const message = messageInput.value;
            if (message) {
                ws.send(message);
                addMessage('You: '+message);
                messageInput.value = ''; // Clear input
            }
        });

        // Add message to the messages div
        function addMessage(msg) {
            const p = document.createElement('p');
            p.textContent = msg;
            messagesDiv.appendChild(p);
        }

        function submitForm() {
          const formData = new FormData();
          formData.append("file", document.querySelector("#fileInput").files[0]);

          fetch("http://localhost:1900/upload?key="+token, {
              method: "POST",
              body: formData
          })
          .then(response => alert(response))
          .catch(error => console.error("Error:", error));
        }

    </script>
  </html>`
	w.Write([]byte(tmpl))
}
