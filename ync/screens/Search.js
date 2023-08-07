import React, { Component } from "react";
import { Container, Header, Item, Icon, Button, Text, Input, Spinner } from "native-base";
import { isEmpty } from "lodash";
import { getVideos } from "../api/youtube";
import VideoList from "../components/VideoList";
import Config from "../config";

export default class Search extends Component {
  state = {
    query: "",
    searching: false,
    videos: []
  };

  search = () => {
    this.setState({ searching: true }, () =>
      getVideos({ q: this.state.query })
        .then(videos => this.setState({ videos }))
        .then(() => this.setState({ searching: false }))
    );
  };

  render() {
    return (
      <Container>
        <Header searchBar rounded>
          <Item>
            <Icon name="search" />
            <Input
              placeholder="Search"
              value={this.state.query}
              onChangeText={query => this.setState({ query })}
              onSubmitEditing={this.search}
            />
            <Icon name="videocam" />
          </Item>
          <Button transparent>
            <Text>Search</Text>
          </Button>
        </Header>
        {!isEmpty(this.state.videos) ? (
          <VideoList videos={this.state.videos} {...this.props} />
        ) : null}
      </Container>
    );
  }
}
