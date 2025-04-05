import React from 'react';
import '../General.css';
import './Workspace.css';
import {fakeData as notes} from '../assets/fakeData.js';
import NoteCard from '../components/NoteCard';

function Workspace() {
    return (
        <div className='workspace'>
            <div className='notes'>
                {notes.map((note) => (
                    <div className='note-container'>
                        <NoteCard note = {note} key = {note.$id} />
                    </div>
                ))}
                <button className='new-note'>+</button>
            </div>
            <div className='sidebar'>
                <h1>Title</h1>
                <button>Open</button>
            </div>
        </div>
    );
}

export default Workspace;