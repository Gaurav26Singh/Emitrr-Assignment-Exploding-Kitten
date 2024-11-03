
import React from 'react';


function Navbar({ userName, points }) {

  const handlerestart = ()=>{
    const answer = window.confirm("Are you sure, you want to exit the game?");
    if (answer) {
        window.location.reload();
    }
    
  }
  return (
    <div className="navbar">
      <div className="navbar-left">
        <h1 style={{cursor:'pointer'}} onClick={handlerestart}>Exploding Kittens 😸</h1>
      </div>
      <div className="navbar-right">
        <p>User: {userName}</p>
        <p>Points: {points}</p>
      </div>
    </div>
  );
}

export default Navbar;
