import React from "react";
import "./SignInForm.css";

// TEST 
import axios from "axios";

class Logo extends React.Component {
  render() {
    return (
      <div className="logo">
        <i className="fa fa-bug" aria-hidden="true" />
        <span> Login </span>
      </div>
    );
  }
}

class Input extends React.Component {
  render() {
    return (
      <div className="Input">
        <input
          type={this.props.type}
          name={this.props.name}
          placeholder={this.props.placeholder}
          required
          autoComplete="false"
        />
        <label htmlFor={this.props.name} />
      </div>
    );
  }
}

// Modal
class Modal extends React.Component {
  constructor() {
    super();
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

    console.log("Valid fields entered..");

    credentials.email = data.get("username");
    credentials.password = data.get("password");

    axios.post("http://localhost:8080/api/v1/signin", credentials)
    .then(res => { console.log(res.data); })
    .catch(err => { console.log(err); });
  }

  render() {
    return (
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
        <a href="#">Lost your password ?</a>
      </div>
    );
  }
}


class SignInForm extends React.Component {
  render() {
    return (
      <div className="background">
        <Modal onSubmit={this.handleSubmit} />
      </div>
    );
  }
}

export default SignInForm;
