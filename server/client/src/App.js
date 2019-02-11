import React, { Component } from 'react';

import Header from "./components/partials/Header";
import Footer from "./components/partials/Footer";
import Main from "./components/Main";


class App extends Component {
  render() {
    return (
      <div>
        <Header />
        <Main />
      </div>
    );
  }
}

export default App;
