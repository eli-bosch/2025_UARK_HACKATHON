import React from 'react';
import '../General.css';
import './SharedWorkspace.css';
import {fakeData as notes} from '../assets/fakeData.js';
import NoteDisplay from '../components/NoteDisplay';
import Chatbar from '../components/Chatbar.js';

function SharedWorkspace() {
    return (
        <div className='workspace'>
            <div className='chatbar'>
                <Chatbar/>
            </div>
            <NoteDisplay notes={notes}/>
            <div className='sidebar'>
                <h1>Title</h1>
                <button>Open</button>
            </div>
        </div>
    );
}

export default SharedWorkspace;