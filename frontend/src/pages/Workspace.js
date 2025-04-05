import React from 'react';
import '../General.css';
import './Workspace.css';
import {fakeData as notes} from '../assets/fakeData.js';
import NoteDisplay from '../components/NoteDisplay';
import Sidebar from '../components/Sidebar.js';

function Workspace() {
    return (
        <div className='workspace'>
            <div className='eighty'>
                <NoteDisplay notes={notes}/>
            </div>
            <div className='twenty'>
                <Sidebar/>
            </div>
        </div>
    );
}

export default Workspace;