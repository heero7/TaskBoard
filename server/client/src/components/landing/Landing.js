import React, { Component } from "react";
import { Button } from "react-materialize";
import { Link } from "react-router-dom";

import "./Landing.css";

class Landing extends Component {
  render() {
    return (
      <div>
        <HeroSection />
        <PhotoSection />
        <FeaturesSection />
        <SignInUpSection />
      </div>
    );
  }
}

const HeroSection = () => {
  return (
    <section className="hero">
      <div className="background-image" />
      <div className="hero-content-area">
        <h1>TaskBoard</h1>
        <h3>Tasks prioritized like a pro.</h3>
        <Button>
          <Link to={"/signup"}>Sign Up</Link>
        </Button>
      </div>
    </section>
  );
};

const PhotoSection = () => {
  return (
    <section className="destinations">
      <h3 className="title">Get the best out of your day</h3>
      <p>Organizing has never been easier</p>
      <hr />
      <ul className="grid">
        <li id="landing-block-1" className="large" />
        <li id="landing-block-2" className="small" />
        <li id="landing-block-3" className="small" />
        <li id="landing-block-4" className="large" />
      </ul>
    </section>
  );
};

const FeaturesSection = () => {
  return (
    <section class="packages">
      <h3 class="title">Features</h3>
      <p>
        We offer a variety of features <br />
        Here are a few important features.
      </p>
      <hr />

      <ul class="grid">
        <li>
          <i class="fa fa-list fa-4x" />
          <h4>Create your Lists</h4>
          <p>
            Write down your own lists to keep track of and add them to your
            board!
          </p>
        </li>
        <li>
          <i class="fa fa-trophy fa-4x" />
          <h4>Achievements</h4>
          <p>From your finished tasks, create meaningful achievements.</p>
        </li>
        <li>
          <i class="fa fa-file-text fa-4x" />
          <h4>Write to resumes</h4>
          <p>
            Take those achievements and write them to your resume in the form of
            a blurb. Don't let those tasks go unoticed
          </p>
        </li>
        <li>
          <i class="fa fa-user fa-4x" />
          <h4>Create your account</h4>
          <p>By creating your account you can do this on the go anywhere</p>
        </li>
      </ul>
    </section>
  );
};

const SignInUpSection = () => {
  return (
    <section class="contact">
      <h3 class="title">Learn More</h3>
      <p>Why just read about it, try it yourself today!</p>
      <hr />
      <div className="btns">
        <Button>
          <Link to={"/signin"}>Sign In</Link>
        </Button>
        <Button>
          <Link to={"/signup"}>Sign Up</Link>
        </Button>
      </div>
    </section>
  );
};

export default Landing;
