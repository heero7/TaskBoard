import React from "react";

import { connect } from "react-redux";

import { Link } from "react-router-dom";

import { Navbar } from "react-materialize";

class Header extends React.Component {
  renderNavLinks() {
    switch(this.props.auth) {
      case null:
        return;
      case false:
        return [
          <li><Link to={"/"}>Home</Link></li>,
          <li><Link to={"/signin"}>Sign In</Link></li>,
          <li><Link to={"/signup"}>Sign Up</Link></li>
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

export default connect(mapStateToProps)(Header);
