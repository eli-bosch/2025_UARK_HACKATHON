import NoteCard from "./NoteCard";
import './NoteDisplay.css';

function NoteDisplay({notes}) {
    return(
        <div className='notes'>
            {notes.map((note) => (
                <div className='note-container'>
                    <NoteCard note = {note} key = {note.$id} />
                </div>
            ))}
        </div>
    );
}

export default NoteDisplay;

