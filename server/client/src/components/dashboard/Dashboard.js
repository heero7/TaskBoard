import React from "react";
import { connect } from "react-redux";

class Dashboard extends React.Component {
    render() {
        return (
            <div>
                <h1>Tasks</h1>
                <br />
            </div>
        );
    }
}

function mapStateToProps(state) {
    const { authentication } = state;
    const { user } = authentication;
    return {
        user
    };
}



export default connect(mapStateToProps)(Dashboard);