import './NoteCard.css';
import { useEffect, useRef, useState } from "react";
// import Trash from "../icons/Trash";
import { autoGrow, bodyParser } from "../utils";
// import { db } from "../appwrite/databases";
import Spinner from "../icons/Spinner";
import DeleteButton from "../components/DeleteButton";
// import { useContext } from "react";
// import { NotesContext } from "../context/NotesContext";

const NoteCard = ({ note }) => {
    const cardRef = useRef(null);

    //let colors = {};
    let header = "";

    /* try {
    colors = note.colors ? JSON.parse(note.colors) : {};
    } catch (err) {
    console.error("Invalid colors JSON", err);
    } */

    try {
    header = note.Header;
    } catch (err) {
    console.error("Invalid header JSON", err);
    }


    // const { setSelectedNote } = useContext(NotesContext);

    const [saving, setSaving] = useState(false);
    const keyUpTimer = useRef(null);
    const body = bodyParser(note.Body);

    const textAreaRef = useRef(null);
    const headerAreaRef = useRef(null);

    useEffect(() => {
        autoGrow(textAreaRef);
        // setZIndex(cardRef.current);
    }, []);



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
            saveData("body", textAreaRef.current.value);
            saveData("header", headerAreaRef.current.value);
        }, 2000);
    };

    return (
        <div
            ref={cardRef}
            className="card"
        >
            <div
                className="card-header"
            >
                <DeleteButton noteId={note.$id} />

                <input
                    type='text'
                    onKeyUp={handleKeyUp}
                    ref={headerAreaRef}
                    defaultValue={header}
                ></input>

        
                <div className="card-saving">
                    {(saving && 
                        <div>
                            <Spinner></Spinner>
                            <span>
                                Saving...
                            </span>
                        </div>
                    )}
                    
                </div>
            </div>
            <div className="card-body">
                <textarea
                    onKeyUp={handleKeyUp}
                    onFocus={() => {
                        // setZIndex(cardRef.current);
                        // setSelectedNote(note);
                    }}
                    onInput={() => {
                        autoGrow(textAreaRef);
                    }}
                    ref={textAreaRef}
                    defaultValue={body}
                ></textarea>
            </div>
        </div>
    );
};

export default NoteCard;