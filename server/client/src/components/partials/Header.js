import React from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";

import { Navbar } from "react-materialize";

class Header extends React.Component {
  renderNavLinks() {
    console.log(this.props.authentication.loggedIn);
    switch(this.props.authentication.loggedIn) {
      case undefined:
      case false:
        return [
          <li key={1}><Link to={"/"}>Home</Link></li>,
          <li key={2}><Link to={"/signin"}>Sign In</Link></li>,
          <li key={3}><Link to={"/signup"}>Sign Up</Link></li>
        ];
      default: 
        return [
          <li key={4}><Link to={"/logout"}>Logout</Link></li>,
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

function mapStateToProps(state) {
  const { authentication } = state;
  return {
    authentication
  };
}

export default connect(mapStateToProps)(Header);
