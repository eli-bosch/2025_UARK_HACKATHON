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
            <div className='sixty'>
                <NoteDisplay notes={notes}/>
                <NoteDisplay notes={notes}/>
            </div>
            <div className='twenty'>
                <Sidebar/>
            </div>
        </div>
    );
}

export default SharedWorkspace;