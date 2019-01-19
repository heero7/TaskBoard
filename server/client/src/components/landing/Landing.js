import React, { Component } from "react";

import "./Landing.css";

class Landing extends Component {
  render() {
    return (
      <section class="hero">
        <div className="background-image"/>
        <div class="hero-content-area">
          <h1>TaskBoard</h1>
          <h3>Tasks prioritized like a pro.</h3>
          <a href="#" class="btn">
            Sign up now!
          </a>
        </div>
      </section>
    );
  }
}

export default Landing;
