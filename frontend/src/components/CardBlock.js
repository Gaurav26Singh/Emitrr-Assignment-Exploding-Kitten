
import React, { useState, useEffect } from 'react';

function CardBlock({ cardType, index, onCardFlip, flippedArray }) {
  const [flipped, setFlipped] = useState(flippedArray[index]);

  useEffect(() => {
    setFlipped(flippedArray[index]);
  }, [flippedArray, index]);

  const handleClick = () => {
    if (!flipped) {
      setFlipped(true);
      onCardFlip(index, cardType);
    }
  };

  const getBackImage = () => {
    switch (cardType) {
      case 'cat':
        return '😼';
      case 'defuse':
        return '🙅‍♂️';
      case 'shuffle':
        return '🔀';
      case 'exploding':
        return '💣';
      default:
        return '';
    }
  };

  return (
    <div className={`card-block ${flipped ? 'flipped' : ''}`} onClick={handleClick}>
      <div className="card-front">

      </div>
      <div className="card-back" >
        {flipped} {getBackImage()}
      </div>
    </div>
  );
}

export default CardBlock;
