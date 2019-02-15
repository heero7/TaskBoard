import React from "react";
import { Router, Route } from "react-router-dom";

import { history } from "../util/history";

import PrivateRoute from "./reusables/PrivateRoute";

import Header from "./partials/Header";
import Landing from "./landing/Landing";
import SignIn from "./signin/SignInForm";
import SignUp from "./signup/SignUpForm";
// Protexted routes
import Dashboard from "./dashboard/Dashboard";

class Main extends React.Component {

   render() {
    return ( 
        <div>
            <Router history={history}>
                <div>
                    <Header />
                    <Route exact path="/" component={Landing}/>
                    <Route path="/signin" component={SignIn}/>
                    <Route path="/signup" component={SignUp}/>
                    <PrivateRoute path="/dashboard" component={Dashboard} />
                </div>
            </Router>
        </div>
    ); 
    }
}

export default Main;