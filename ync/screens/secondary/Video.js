import React, { Component } from "react";
import AppHeader from "../../components/AppHeader";

export default class Video extends Component {
  render() {
    return (
      <>
        <AppHeader title="Video" {...this.props} />
      </>
    );
  }
}
