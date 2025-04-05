let socket 

function initWebSocket(userID = "sampleUSER") {
    socket = new WebSocket ('ws://localhost:8080/ws')

    socket.onopen = () => {
        console.log('WebSocket connection established')

        // confirm user is online
        if (socket.readyState === socket.OPEN) {
            const idMessage = JSON.stringify({ userID} );
            socket.send(idMessage);
        }
    }

    socket.onmessage = (event) => {
        const msg = JSON.parse(event.data)
        displayMessage(msg.message)
    }

    socket.onclose = () => {
        console.log('WebSocket connection closed!')
    }

    socket.onerror = (error) => {
        console.error('WebSocket error:', error)
    }
}

// probably will redesign
function sendMessageToServer() {
    const input = document.getElementById('chatInput');
    const message = input.value;
    if (socket && socket.readyState === WebSocket.OPEN) {
        const msg = JSON.stringify({ message });
        socket.send(msg);
        input.value = '';
    }
}

function displayMessage(text) {
    const ul = document.getElementById('messages');
    const li = document.createElement('li');
    li.textContent = text;
    ul.appendChild(li);
}

export { initWebSocket, sendMessageToServer }