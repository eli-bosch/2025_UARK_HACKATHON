import NoteCard from "./NoteCard";

function NoteDisplay({ notes }) {
    if (!Array.isArray(notes)) {
        return <div>Loading...</div>;
    }

    return (
        <div className='notes'>
            {notes.map((note, index) => {
                const key = note?.$id || note?.id || index;

                //const isBson = note instanceof Uint8Array || note instanceof ArrayBuffer;

                return (
                    <div className='note-container' key={key}>
                            <NoteCard note={note} />
                    </div>
                );
            })}
            <button className='new-note'>+</button>
        </div>
    );
}

export default NoteDisplay;