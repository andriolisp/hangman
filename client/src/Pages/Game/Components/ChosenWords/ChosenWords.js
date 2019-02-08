import React from "react";
import "./ChosenWords.css";

const ChosenWords = props => {
  if (props.details && props.details.length > 0) {
    const letters = props.details.map(d => {
      return <li className="ChosenWords">{d.letter}</li>;
    });

    return (
      <div>
        <ul className="ChosenWords">{letters}</ul>
      </div>
    );
  }
  return <div />;
};

export default ChosenWords;
