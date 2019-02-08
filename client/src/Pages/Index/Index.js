import React, { Component } from "react";
import "./Index.css";
import { withRouter } from "react-router-dom";

class Index extends Component {
  constructor(props) {
    super(props);
    this.props = props;

    this.state = {
      players: 2
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleGame = this.handleGame.bind(this);
  }

  handleChange(e) {
    if (e.target.validity.valid && !isNaN(e.target.value)) {
      this.setState(
        Object.assign(this.state, {
          players: e.target.validity.valid ? e.target.value : this.state.players
        })
      );
    } else {
      this.setState(Object.assign(this.state, { players: 2 }));
    }
  }

  handleGame() {
    const self = this;
    fetch("http://localhost:7100/v1/game", {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json"
      },
      redirect: "follow",
      referrer: "no-referrer",
      body: JSON.stringify({ num_players: this.state.players })
    })
      .then(response => response.json())
      .then(g => {
        self.props.history.push("/" + g.id);
      });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <div>
            <h1>HANGMAN</h1>
            <h2>Type the number of players</h2>
          </div>
          <div>
            <input
              type="text"
              pattern="[0-9]*"
              onInput={this.handleChange}
              value={this.state.players}
            />
          </div>
          <div>
            <button type="button" onClick={this.handleGame}>
              START GAME
            </button>
          </div>
        </header>
      </div>
    );
  }
}

export default withRouter(Index);
