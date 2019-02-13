import React from "react";
import { BrowserRouter, Route } from "react-router-dom";

import { connect } from "react-redux";
import * as actions from "../actions";

import Header from "./partials/Header";
import Landing from "./landing/Landing";
import SignIn from "./signin/SignInForm";
import SignUp from "./signup/SignUpForm";

class Main extends React.Component {
    componentDidMount() {
        console.log(this.props);
    }
   render() {
    return ( 
        <div>
            <BrowserRouter>
                <div>
                    <Header />
                    <Route exact path="/" component={Landing}/>
                    <Route path="/signin" component={SignIn}/>
                    <Route path="/signup" component={SignUp}/>
                </div>
            </BrowserRouter>
        </div>
    ); 
    }
}

export default connect(null, actions)(Main);