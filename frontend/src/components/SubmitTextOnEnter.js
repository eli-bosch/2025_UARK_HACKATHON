import React, { useState } from 'react';

function SubmitTextOnEnter() {
  const [value, setValue] = useState('');

  const handleKeyDown = (e) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault(); // prevent newline
      handleSubmit();
    }
  };

  const handleSubmit = () => {
    if (value.trim()) {
      console.log('Submitted:', value);
      setValue('');
    }
  };

  return (
    <div>
      <textarea
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onKeyDown={handleKeyDown}
        rows={4}
        placeholder="Enter message here"
        style={{ width: '283px', padding: '8px' }}
      />
    </div>
  );
}

export default SubmitTextOnEnter;