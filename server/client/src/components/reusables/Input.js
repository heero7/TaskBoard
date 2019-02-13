import React from "react";

class Input extends React.Component {
    render() {
      return (
        <div className="Input">
          <input
            type={this.props.type}
            name={this.props.name}
            placeholder={this.props.placeholder}
            required
            autoComplete="false"
          />
          <label htmlFor={this.props.name} />
        </div>
      );
    }
  }

  export default Input;