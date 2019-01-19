import React, { Component } from 'react';

import Header from "./components/header/Header";
import Landing from "./components/landing/Landing";


class App extends Component {
  render() {
    return (
      <div>
        <Header />
        <Landing />
      </div>
    );
  }
}

export default App;
