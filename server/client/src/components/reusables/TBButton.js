import React from "react";
import { Button } from "react-materialize";

import "../styles/Button.css";

const TBButton = props => {
    return (
        <Button className="tb-button">{props.title}</Button>
    );
}

export default TBButton;