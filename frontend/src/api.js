export async function loadNoteData(username) {
    try {
        const response = await fetch('http://localhost:9010/note/user', {
            method: 'POST', // POST request, not GET
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username: username }) // Send username in the body
        });

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        const noteData = await response.json();
        return noteData;
    } catch (error) {
        console.error("Error loading notes:", error);
    }
}