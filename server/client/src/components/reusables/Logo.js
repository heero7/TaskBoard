import React from "react";

class Logo extends React.Component {
    render() {
      let img = this.props.login == true
        ? <img src="https://img.icons8.com/dusk/64/000000/task.png" alt="" /> 
        :  <img src="https://img.icons8.com/dusk/64/000000/add-user-male.png" alt="" />;

      return (
        <div className="logo">
          {img}
          <span> {this.props.title} </span>
        </div>
      );
    }
  }

  export default Logo;