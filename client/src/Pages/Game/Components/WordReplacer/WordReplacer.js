import React from "react";
import "./WordReplacer.css";

const WordReplacer = props => {
  if (props.replacers && props.replacers.length > 0) {
    return <h2 className="WordReplacer">{props.replacers.join("  ")}</h2>;
  }
  return <div />;
};

export default WordReplacer;
