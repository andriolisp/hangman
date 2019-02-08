import React, { Component } from "react";
import "./Game.css";
import WordReplacer from "./Components/WordReplacer";
import LetterChoice from "./Components/LetterChoice";
import ChosenWords from "./Components/ChosenWords";
import Players from "./Components/Players";
import { withRouter } from "react-router-dom";

class Game extends Component {
  constructor(props) {
    super(props);
    this.props = props;

    this.getGame = this.getGame.bind(this);
    this.handleWord = this.handleWord.bind(this);
    this.handleNewGame = this.handleNewGame.bind(this);

    this.state = {
      id: "",
      size: 0,
      remaining: 0,
      turn: 1,
      num_players: 0,
      winner: 0,
      replacers: [],
      players: {},
      message: ""
    };
  }

  componentDidMount() {
    this.getGame(this.props.match.params.id);
  }

  componentWillReceiveProps(nextProps) {
    this.props = nextProps;
    this.getGame(this.props.match.params.id);
  }

  handleWord(g) {
    this.setState(Object.assign(this.state, g));
  }

  handleNewGame() {
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
      body: JSON.stringify({ num_players: this.state.num_players })
    })
      .then(response => response.json())
      .then(g => {
        self.props.history.push("/" + g.id);
      });
  }

  getGame(ID) {
    const self = this;
    fetch(`http://localhost:7100/v1/game/${ID}`, {
      method: "GET",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json"
      },
      redirect: "follow",
      referrer: "no-referrer"
    })
      .then(response => response.json())
      .then(g => {
        if (!g.message) g.message = "";
        if (!g.details) g.details = [];

        self.setState(Object.assign(self.state, g));
      });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <div>
            <WordReplacer replacers={this.state.replacers} />
          </div>
          <div>
            <ChosenWords details={this.state.details} />
          </div>
          <div>
            <h3 className="Message">{this.state.message}</h3>
          </div>
          <div>
            <LetterChoice
              game={this.state}
              onLetterChosen={this.handleWord}
              onNewGame={this.handleNewGame}
            />
          </div>
          <div>
            <Players players={this.state.players} />
          </div>
        </header>
      </div>
    );
  }
}

export default withRouter(Game);
