import React from "react";
import { Link } from "react-router-dom";
import { connect } from "react-redux";
import { userActions } from "../../actions/userActions";

import Input from "../reusables/Input";
import Logo from "../reusables/Logo";

import { Button } from "react-materialize";
import "./SignInForm.css";

class SignInForm extends React.Component {
  constructor(props) {
    super(props);

    // If we are coming back here by anymeans
    // reset the login status to original state
    this.props.dispatch(userActions.logout());

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  /**
   * Handle the submitting of the data and
   * call dispatch the action to login. First,
   * check if we have a user and password entered
   * @param {Event} e
   */
  handleSubmit(e) {
    e.preventDefault();
    let data = new FormData(e.target);
    let credentials = {
      email: "",
      password: ""
    };

    // todo: check to see if the user values are not valid

    if (!data.get("username") || !data.get("password")) {
      console.error("Missing field, check form fields");
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
          <Logo title="Login" login={true}/>
          <form onSubmit={this.handleSubmit} method="POST">
            <Input type="text" name="username" placeholder="username" />
            <Input type="password" name="password" placeholder="password" />
            <Button className="modal-button"> Sign In</Button>
          </form>
          <Link  to={"/signup"}><Button className="modal-button">Sign Up</Button></Link>
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
