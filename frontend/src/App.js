import './App.css';
import { useEffect } from 'react';
import { initWebSocket, sendMessageToServer } from './ws'; 

function App() {
  useEffect(() => {
    initWebSocket(); 
  }, [])

 /* const sendPostRequest = async () => {
     try {
        const res = await fetch("http://localhost:9010/test", {
          method: "POST",
          headers: {
            "Content-Type": "text/plain",
          },
          body: "hi",
        });

        if (!res.ok) throw new Error('Network response was not ok');

        const text = await res.text();
        alert(text);
      } catch (error) {
        console.error('Error during fetch:', error)
      }
    }; */
      

  return (
    <div>
      <h1>WebSocket Test</h1>
      <input type="text" id="chatInput" placeholder="Type a message" />
      <button onClick={sendMessageToServer}>Send</button>
      <ul id="messages"></ul>
    </div>
  );
}

export default App;
