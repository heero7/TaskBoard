import React from "react";
import { Switch, Route } from "react-router-dom";

import Landing from "./landing/Landing";
import SignIn from "./signin/SignInForm";
import SignUp from "./signup/SignUpForm";

const Main = () => (
    <main>
        <Switch>
            <Route exact path="/" component={Landing}/>
            <Route path="/signin" component={SignIn}/>
            <Route path="/signup" component={SignUp}/>
        </Switch>
    </main>
)

export default Main;