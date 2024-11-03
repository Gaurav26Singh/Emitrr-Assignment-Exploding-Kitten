
import React, { useState } from 'react';
import axios from 'axios';
import Info from './Info';

function NameForm({ onNameSubmit }) {
  const [name, setName] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const [infobox,setinfobox] = useState(false);

  const handleinfo = ()=>{
    setinfobox(!infobox);
  }

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if(name==="") throw new Error("Name is required");
      await axios.post('http://localhost:8080/api/user', { name });
      onNameSubmit(name); 
      setErrorMessage('');
    } catch (error) {
      console.error('Error submitting name:', error);
      setErrorMessage(`Failed to submit name. Please try again.(${error})`);
    }
  };

  return (

    <div id="nf">
      {infobox===false ? <button id='info' onClick={handleinfo}>i</button>:<button id='info' onClick={handleinfo}>x</button>}
      {infobox && <Info/>}
    <form onSubmit={handleSubmit} style={{ textAlign: 'center' }}>
      <div className="nameformTitle">
        <h1 style={{letterSpacing:'3px',fontSize:'35px'}}>Exploding Kittens😸</h1>
      </div>
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Enter Your Name"
        style={{
          backgroundColor: '#323437',
          border: '1px solid rgb(255, 221, 25)',
          color: 'white',
          padding: '8px',
          borderRadius: '5px',
          marginBottom: '10px',
          width: '200px',
          outline: 'none',
          marginTop:'15px'
        }}
      />
      <br />
      <button
        type="submit"
        style={{
          backgroundColor: 'rgb(255, 221, 25)',
          color: 'black',
          padding: '10px 30px',
          border: 'none',
          borderRadius: '5px',
          cursor: 'pointer',
          fontSize: '16px',
          fontWeight: 'bolder',
          marginTop:'10px'
        }}
      >
        Start
      </button>
      {errorMessage && <p style={{ color: 'rgb(255, 221, 25)', marginTop:'15px', fontSize:'18px'}}>{errorMessage}</p>}
    </form>
    </div>
  );
}

export default NameForm;
