import React from "react";
import { connect } from "react-redux";
import { userActions } from "../../actions";

import Input from "../reusables/Input";

import "./SignUpForm.css";


class SignUpForm extends React.Component {
    constructor() {
        super();
        this.handleSignUp = this.handleSignUp.bind(this);
    }

    handleSignUp(e) {
        e.preventDefault();
        let data = new FormData(e.target);
        let credentials = {
            email : "",
            password : ""
        };

        // todo : add validators 
        if (data.get("password").localeCompare(data.get("confirm")) !== 0) {
            console.error("Passwords aren't equal!");
            return;
        }
        credentials.email = data.get("email");
        credentials.password = data.get("password");

        const { dispatch } = this.props;
        dispatch(userActions.register(credentials));
    }

    render() {
        return (
            <div className="background">
                <div className="Modal">
                    <form onSubmit={this.handleSignUp} method="POST">
                        <Input type="text" name="email" placeholder="email"/>
                        <Input type="password" name="password" placeholder="password"/>
                        <Input type="password" name="confirm" placeholder="password" />
                        <button>Sign Up!</button>
                    </form>
                </div>
            </div>
        );
    }
}

function mapStateToProps(state) {
    const { registering } = state.registration;
    return {
        registering
    };
}

export default connect(mapStateToProps)(SignUpForm);