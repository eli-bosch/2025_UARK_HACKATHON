import './Sidebar.css';
import { useEffect, useRef, useState } from "react";

function Sidebar() {
    const [saving, setSaving] = useState(false);
    const keyUpTimer = useRef(null);

    const titleRef = useRef(null);

    const saveData = async (key, value) => {
        const payload = { [key]: JSON.stringify(value) };
        console.log("Save data called:", payload);
        try {
            // await db.notes.update(note.$id, payload);
        } catch (error) {
            console.error(error);
        }
        setSaving(false);
    };

    const handleKeyUp = async () => {
        setSaving(true);
        if (keyUpTimer.current) {
            clearTimeout(keyUpTimer.current);
        }

        keyUpTimer.current = setTimeout(() => {
            console.log("Timer started");
            saveData("workspaceTitle", titleRef.current.value);
        }, 2000);
    };


    return(
        <div className='sidebar'>
            <input 
                type='text'
                defaultValue='title'
                className='workspace-title'
                ref={titleRef}
            />
            <button className='open-connection'>Open</button>
            <button className='new-note'>+ New Note +</button>
        </div>
    );
}

export default Sidebar;