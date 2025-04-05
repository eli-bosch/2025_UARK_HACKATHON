import React from 'react';
import '../General.css';
import './Workspace.css';
import { loadNoteData } from '../api.js';
import NoteDisplay from '../components/NoteDisplay';
import { useState, useEffect } from 'react';

function Workspace() {
    const [notes, setNotes] = useState([]); // State to store the notes

    useEffect(() => {
        const username = "ebosch"; // Replace with actual user ID, probably from login or context
        const fetchNotes = async () => {
            const notesData = await loadNoteData(username); // Fetch the notes
            setNotes(notesData); // Update state with the loaded notes
        };
        fetchNotes();
    }, []); // Empty dependency array to run this only once after initial render

    return (
        <div className='workspace'>
            <NoteDisplay notes={notes}/>
            <div className='sidebar'>
                <h1>Title</h1>
                <button>Open</button>
            </div>
        </div>
    );
}

export default Workspace;