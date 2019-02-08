import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import Index from "./Pages/Index/Index";
import Game from "./Pages/Game/Game";
import * as serviceWorker from "./serviceWorker";

import { BrowserRouter as Router, Route } from "react-router-dom";

ReactDOM.render(
  <Router>
    <div>
      <Route exact path="/" component={Index} />
      <Route path="/:id" component={Game} />
    </div>
  </Router>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
