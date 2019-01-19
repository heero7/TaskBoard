import React, { Component } from "react";

import "./Landing.css";

class Landing extends Component {
  render() {
    return (
      <div>
        <HeroSection />
        <ExampleSection />
      </div>
    );
  }
}

const HeroSection = () => {
  return (
    <section className="hero">
      <div className="background-image"/>
      <div className="hero-content-area">
        <h1>TaskBoard</h1>
        <h3>Tasks prioritized like a pro.</h3>
        <a href="#" class="btn">
          Sign up now!
        </a>
      </div>
    </section>
  )
}

const ExampleSection = () => {
  const img1 = {
    backgroundImage: "url(http://www.hotelalexzermatt.com/wp-content/uploads/2015/05/cr-visualimpact.ch-Rainer-Eder-1.jpg)"
  }
  const img2 = {
    backgroundImage: "url(https://wallpapercave.com/wp/cYQcruq.jpg)"
  }
  const img3 = {
    backgroundImage: "url(http://www.redresort.com/wp-content/uploads/2016/04/tabphoto-summer-1.jpg)"
  }
  const img4 = {
    backgroundImage: "url(https://www.thephotoargus.com/wp-content/uploads/2013/01/mountain-02.jpg)"
  }
  return (
    <section className="destinations">
      <h3 className="title">Get the best out of your day</h3>
      <p>Organizing has never been easier</p>
      <hr/>
      <ul class="grid">
        <li className="large" style={img2}></li>
          <li className="small" style={img1}></li>
          <li className="small" style={img3}></li>
          <li className="large" style={img4}></li>
      </ul>
    </section>
  )
}

export default Landing;
