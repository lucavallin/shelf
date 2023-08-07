import React, { Component } from "react";
import AppHeader from "../components/AppHeader";
import { getVideos } from "../api/youtube";
import VideoList from "../components/VideoList";

export default class Home extends Component {
  state = {
    videos: []
  };

  componentDidMount() {
    getVideos({ order: "date", maxResults: 10 }).then(videos => this.setState({ videos }));
  }

  render() {
    return (
      <>
        <AppHeader title="Home" {...this.props} />
        <VideoList videos={this.state.videos} {...this.props} />
      </>
    );
  }
}
