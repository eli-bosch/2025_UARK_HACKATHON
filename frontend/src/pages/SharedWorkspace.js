import React from 'react';
import '../General.css';
import './SharedWorkspace.css';
import {fakeData as notes} from '../assets/fakeData.js';
import NoteDisplay from '../components/NoteDisplay';
import Chatbar from '../components/Chatbar.js';
import Sidebar from '../components/Sidebar.js';

function SharedWorkspace() {
    return (
        <div className='workspace'>
            <div className='chatbar'>
                <Chatbar/>
            </div>
            <div className='notes-section'>
                <div className='user-notes'>
                    <NoteDisplay notes={notes}/>
                </div>
                <div className='guest-notes'>
                    <NoteDisplay notes={notes}/>
                </div>
            </div>
            <div className='summary-sidebar'>
                <Sidebar/>
            </div>
        </div>
    );
}

export default SharedWorkspace;