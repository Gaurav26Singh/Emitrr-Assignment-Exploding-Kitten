import React from 'react'

const Info = () => {
  return (
    <div id='infobox'>
        <p>This will be an online single-player card game that consists of 4 different types of cards</p>
       <p> - Cat card 😼</p>
       <p> - Defuse card 🙅‍♂️</p>
        <p>- Shuffle card 🔀</p>
       <p> - Exploding kitten card 💣</p>
       <br/>
      
    <p>There will be a button to start the game. When the game is started there will be a deck of 5 cards ordered randomly. Each time user clicks on the deck a card is revealed and that card is removed from the deck. <br/>
        A player wins the game once he draws all 5 cards from the deck and there is no card left to draw. </p>

<h4 style={{marginTop:'10px',marginBottom:'10px'}}>Rules</h4>

    <ul className="rules">
        <li>If the card drawn from the deck is a cat card, then the card is removed from the deck.</li>
        <li> If the card is a defusing card, then the card is removed from the deck. This card can be used to defuse one bomb that may come in subsequent cards drawn from the deck.</li>
        <li>If the card is exploding kitten (bomb) then the player loses the game.</li>
        <li> If the card is a shuffle card, then the game is restarted and the deck is filled with 5 cards again.</li>
    </ul>

    </div>
  )
}

export default Info
