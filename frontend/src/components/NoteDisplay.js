import NoteCard from "./NoteCard";

function NoteDisplay({notes}) {
    if (!Array.isArray(notes)) {
        // change to blank starting note if no previous note data exists (new user)
        return <div>Loading...</div>; // Show loading message if notes is not an array
    }
    return(
        <div className='notes'>
            {notes.map((note) => (
                <div className='note-container' key = {note.$id || note.id || Math.random()}>
                    <NoteCard note = {note}  />
                </div>
            ))}
            <button className='new-note'>+</button>
        </div>
    );
}

export default NoteDisplay;

