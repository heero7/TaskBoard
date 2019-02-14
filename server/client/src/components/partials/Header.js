import React from "react";


import { Link } from "react-router-dom";

import { Navbar } from "react-materialize";

class Header extends React.Component {
  renderNavLinks() {
    console.log(this.props);
    switch(this.props.auth) {
      case undefined:
        return [
          <li key={1}><Link to={"/"}>Home</Link></li>,
          <li key={2}><Link to={"/signin"}>Sign In</Link></li>,
          <li key={3}><Link to={"/signup"}>Sign Up</Link></li>
        ];
      default: 
        return [
          <li><Link to={"/logout"}>Logout</Link></li>,
        ];
    }
  }

  render() {
    return (
      <Navbar right>
        {this.renderNavLinks()}
      </Navbar>
    );
  }
}

function mapStateToProps({ auth }) {
  return { auth };
}

export default Header;
