import React from "react";
import { Link } from "react-router-dom";

import { Navbar } from "react-materialize";

// todo: create method to render correct links

const Header = () => {
  return (
    <Navbar right>
      <li>
        <Link to={"/"}>Home</Link>
      </li>
      <li>
        <Link to={"/signin"}>Sign In</Link>
      </li>
      <li>
        <Link to={"/signup"}>Sign Up</Link>
      </li>
    </Navbar>
  );
};

export default Header;
