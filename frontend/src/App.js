import './App.css';

function App() {

  const sendPostRequest = async () => {
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
    };
      

  return (
    <div>
      <h1>Go POST Test</h1>
      <button onClick={sendPostRequest}>SEND POST</button>
    </div>
  );
}

export default App;
