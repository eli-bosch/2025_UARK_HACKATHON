import NoteCard from "./NoteCard";
import './NoteDisplay.css';

function NoteDisplay({ notes }) {
    if (!Array.isArray(notes)) {
        return <div>Loading...</div>;
    }

    return (
        <div className='notes'>
            {notes.map((note, index) => {
                const key = note?.$id || note?.id || index;

                return (
                    <div className='note-container' key={key}>
                            <NoteCard note={note} />
                    </div>
                );
            })}
        </div>
    );
}

export default NoteDisplay;