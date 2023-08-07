import React, { Component } from "react";
import AppHeader from "../components/AppHeader";
import { getVideos } from "../api/youtube";
import VideoList from "../components/VideoList";

export default class Latest extends Component {
  state = {
    videos: []
  };

  componentDidMount() {
    getVideos({ order: "date", maxResults: 30 }).then(videos => this.setState({ videos }));
  }

  render() {
    return (
      <>
        <AppHeader title="Ultimi video" {...this.props} />
        <VideoList videos={this.state.videos} {...this.props} />
      </>
    );
  }
}
