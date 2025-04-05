# 2025_UARK_HACKATHON
Repo for the University of Arkansas' 2025 Hackathon

# My final gift
For the note database plugin part, look to the workspace.js items in the frontend/src/pages folders
Inside you will see NoteDisplay has a "note" parameter, see fakeData.js under the assets folder to see the format this accepts data

That that "notes" object (different from the name of the parameter this is the actual object) gets passed down to NoteDisplay.js under components,
Here it gets mapped from notes onto each note inside, **This is where you would probably want to filter based on user information only displaying the notes for the currently logged in user**

Then inside NoteCard.js the contents of these notes are parsed into its header and body variables.

Inside NoteCard.js is also where the saving of the notes happens. There is a function called saveData which does exactly what it says.
Inside Deletebutton.js there is the functions for deleting notes.

Similarly there is a saveWorkspace title functionality inside the SideBar.js component.

I have not added create note functionality https://github.com/divanov11/Sticky-Notes-React/blob/master/src/components/AddButton.jsx This might be a good place to start with that

Inside the chatbar.js there is a similar mapping structure to NoteDisplay.js for displaying the chat messages, check fakeData2.js for the current structure of these.

Ill have my phone, feel free to call me while I am driving as long as its **not between 11am and 12pm**

# Initial Setup for other developers
run `npm install`

# Credits:
ChatGPT
freeCodeCamp -> https://sticky-fcc.vercel.app/
