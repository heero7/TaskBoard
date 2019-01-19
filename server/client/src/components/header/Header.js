import React, { Component } from "react";

class Header extends Component {
    render() {
        return (
            <header>
                <h2><a href="#">Mountain Traveling</a></h2>
            <nav>
                <li><a href="#">SignIn/SignUp</a></li>
                <li><a href="#">About</a></li>
                <li><a href="#">Contact</a></li>
            </nav>
            </header>
        )
    }
}

export default Header;