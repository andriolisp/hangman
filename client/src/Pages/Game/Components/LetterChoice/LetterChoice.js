import React from "react";
import "./LetterChoice.css";

class LetterChoice extends React.Component {
  constructor(props) {
    super(props);

    this.props = props;
    this.handleInput = this.handleInput.bind(this);
    this.handleWord = this.handleWord.bind(this);
    this.handleNewGame = this.handleNewGame.bind(this);
    this.state = {
      letter: ""
    };
  }

  handleInput(e) {
    this.setState(Object.assign(this.state, { letter: e.target.value }));
  }

  handleWord() {
    const self = this;
    fetch(`http://localhost:7100/v1/game/${self.props.game.id}`, {
      method: "PUT",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json"
      },
      redirect: "follow",
      referrer: "no-referrer",
      body: JSON.stringify({ letter: self.state.letter })
    })
      .then(response => response.json())
      .then(g => {
        if (self.props.onLetterChosen) {
          self.props.onLetterChosen(g);
          self.setState(Object.assign(self.state, { letter: "" }));
        } else {
          console.log(g);
        }
      });
  }

  handleNewGame() {
    if (this.props.onNewGame) {
      this.props.onNewGame();
      this.setState(Object.assign(this.state, { letter: "" }));
    }
  }

  render() {
    let message = `Player ${this.props.game.turn}: `;
    if (this.props.game.winner > 0) {
      message = "";
    }

    return (
      <diV>
        <div className="Message">{message}</div>
        <div className="Message">
          <input
            type="text"
            pattern="[A-Za-z]*"
            value={this.state.letter}
            onInput={this.handleInput}
            maxLength={1}
            disabled={this.props.game.winner !== 0}
            className="Letter"
          />
        </div>
        <div className="Message">
          <button
            type="button"
            onClick={this.handleWord}
            disabled={this.state.letter.length === 0}
            className="Check"
          >
            Check
          </button>
        </div>
        <div className="Message">
          <button type="button" onClick={this.handleNewGame} className="Check">
            New Game
          </button>
        </div>
      </diV>
    );
  }
}

export default LetterChoice;
