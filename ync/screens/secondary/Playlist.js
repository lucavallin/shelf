import React, { Component } from "react";
import AppHeader from "../../components/AppHeader";

export default class Playlist extends Component {
  state = {
    videos: []
  };

  render() {
    return (
      <>
        <AppHeader title="Playlist" {...this.props} />
      </>
    );
  }
}
