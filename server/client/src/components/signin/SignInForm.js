import React from "react";
import "./SignInForm.css";

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
          autocomplete="false"
        />
        <label for={this.props.name} />
      </div>
    );
  }
}

// Modal
class Modal extends React.Component {
  handleSubmit(e) {
    e.preventDefault();
    console.log(e.target.value);
  }

  render() {
    return (
      <div className="Modal">
        <Logo />
        <form onSubmit={this.props.onSubmit}>
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
