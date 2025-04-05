import { useEffect, useState } from 'react';
import { BSON } from 'bson';
import NoteCard from './NoteCard';

const BsonNoteLoader = ({ bsonBuffer }) => {
  const [note, setNote] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (!bsonBuffer) return;

    try {
      const parsed = BSON.deserialize(new Uint8Array(bsonBuffer));
      setNote(parsed);
    } catch (e) {
      console.error("Failed to parse BSON", e);
      setError("Invalid BSON format.");
    }
  }, [bsonBuffer]);

  if (error) return <div>{error}</div>;
  if (!note) return <div>Loading BSON note...</div>;

  return <NoteCard note={note} />;
};

export default BsonNoteLoader; 