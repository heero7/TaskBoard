import React from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";

import { Navbar } from "react-materialize";
import { userActions } from "../../actions";

import { Button } from "react-materialize";

class Header extends React.Component {
  constructor(props) {
    super(props);
    this.handleLogout = this.handleLogout.bind(this);
  }

  handleLogout() {
    const { dispatch } = this.props;
    dispatch(userActions.logout());
  }


  renderNavLinks() {
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
          <li key={4} onClick={this.handleLogout}><Button>Logout</Button></li>,
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
