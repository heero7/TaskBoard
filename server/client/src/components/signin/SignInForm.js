import React from "react";
import { connect } from "react-redux";
import { userActions } from "../../actions/userActions";

import Input from "../reusables/Input";
import Logo from "../reusables/Logo";

import "./SignInForm.css";

class SignInForm extends React.Component {
  constructor(props) {
    super(props);

    // If we are coming back here by anymeans
    // reset the login status to original state
    this.props.dispatch(userActions.logout());

    this.handleSubmit = this.handleSubmit.bind(this);
  }


  handleSubmit(e) {
    e.preventDefault();
    let data = new FormData(e.target);
    let credentials = {
      email : "",
      password : ""
    };

    // todo: check to see if the user values are not valid

    if (!data.get("username") || !data.get("password")) {
      console.error("Missing field");
      return;
    }

    credentials.email = data.get("username");
    credentials.password = data.get("password");
    
    const { dispatch } = this.props;
    dispatch(userActions.login(credentials.email, credentials.password));
  }

  render() {
    return (
      <div className="background">
        <div className="Modal">
          <Logo />
          <form onSubmit={this.handleSubmit} method="POST">
            <Input type="text" name="username" placeholder="username" />
            <Input type="password" name="password" placeholder="password" />
            <button> Sign In</button>
          </form>
          <div className="social-signin">
            <button className="fb" onClick={this.props.onClick}>
              <i className="fa fa-facebook" aria-hidden="true" />
            </button>
            <button className="tw" onClick={this.props.onClick}>
              <i className="fa fa-twitter" aria-hidden="true" />
            </button>
          </div>
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { loggingIn } = state.authentication;
  return {
    loggingIn
  };
}

export default connect(mapStateToProps)(SignInForm);
