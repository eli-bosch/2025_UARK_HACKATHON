import './Chatbar.css';
import SubmitTextOnEnter from './SubmitTextOnEnter';
import {fakeData2 as messages} from '../assets/fakeData2.js'

function Chatbar()
{
    return(
        <div>
            <h1 className='chatbar-label'>Chat Log</h1>
            {messages.map((message) => (
                <p className='message-container'>
                    {JSON.parse(message.sender) + ": " + JSON.parse(message.content)}
                </p>
            ))}
            <div className='message-area'>
                <form id='textbox'>
                    {/* <textarea id='chat-message' rows='4' cols='50' placeholder='Enter message here'></textarea> */}
                    <SubmitTextOnEnter/>
                </form>
            </div>
        </div>
    );
}

export default Chatbar;