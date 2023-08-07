import React, { Component } from "react";
import AppHeader from "../components/AppHeader";
import { getPlaylists } from "../api/youtube";
import List from "../components/Playlists";

export default class Playlists extends Component {
  state = {
    playlists: []
  };

  componentDidMount() {
    getPlaylists().then(playlists => this.setState({ playlists }));
  }

  render() {
    return (
      <>
        <AppHeader title="Playlists" {...this.props} />
        <List playlists={this.state.playlists} {...this.props} />
      </>
    );
  }
}
