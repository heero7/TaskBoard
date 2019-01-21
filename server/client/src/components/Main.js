import React from "react";
import { Switch, Route } from "react-router-dom";

import Landing from "./landing/Landing";
import SignIn from "./signin/SignInForm";

const Main = () => (
    <main>
        <Switch>
            <Route exact path="/" component={Landing}/>
            <Route path="/signin" component={SignIn}/>
        </Switch>
    </main>
)

export default Main;