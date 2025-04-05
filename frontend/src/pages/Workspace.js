import React from 'react';
import '../General.css';
import './Workspace.css';
import {fakeData as notes} from '../assets/fakeData.js';
import NoteDisplay from '../components/NoteDisplay';

function Workspace() {
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