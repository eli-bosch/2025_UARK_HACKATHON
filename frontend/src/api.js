export async function loadNoteData(userID) {
    try {
        const response = await fetch(`http://localhost:9010/api/notes/64edc2a5cba3f93e7f7e12b6`);
        const noteData = await response.json();

        return noteData;
    } catch (error) {
        console.error("Error loading notes:", error);
    }
}