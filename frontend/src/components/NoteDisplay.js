import NoteCard from "./NoteCard";

function NoteDisplay({notes}) {
    return(
        <div className='notes'>
            {notes.map((note) => (
                <div className='note-container'>
                    <NoteCard note = {note} key = {note.$id} />
                </div>
            ))}
            <button className='new-note'>+</button>
        </div>
    );
}

export default NoteDisplay;

