import React from "react";
import "./Players.css";

const Players = props => {
  if (props.players) {
    const playerList = Object.keys(props.players).map(p => {
      return (
        <div className="Players">
          <div className="Title">PLAYER {props.players[p].num}</div>
          <div className="LeftSize">Points:</div>
          <div className="RightSize">{props.players[p].points}</div>
          <div className="LeftSize">Turn:</div>
          <div className="RightSize">
            {props.players[p].turn ? "True" : "False"}
          </div>
          <div className="LeftSize">Tentatives:</div>
          <div className="RightSize">{props.players[p].tentatives}</div>
          <div className="LeftSize">Dead:</div>
          <div className="RightSize">
            {props.players[p].dead ? "True" : "False"}
          </div>
        </div>
      );
    });

    return <div>{playerList}</div>;
  }
  return <div />;
};

export default Players;
