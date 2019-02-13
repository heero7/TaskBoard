import React from "react";

import Input from "../reusables/Input";

import "./SignUpForm.css";

import axios from "axios";

class Modal extends React.Component {
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

        axios.post("http://localhost:8080", credentials)
        .then(res => console.log(data))
        .catch(err => console.log(err));
    }

    render() {
        return (
            <div className="Modal">
                <form onSubmit={this.handleSignUp} method="POST">
                    <Input type="text" name="email" placeholder="email"/>
                    <Input type="password" name="password" placeholder="password"/>
                    <Input type="password" name="confirm" placeholder="password" />
                    <button>Sign Up!</button>
                </form>
            </div>
        );
    }
}

class SignUpForm extends React.Component {
    render() {
        return (
            <div className="background">
                <Modal onSubmit={this.handleSignUp}/>
            </div>
        );
    }
}

export default SignUpForm;