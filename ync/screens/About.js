import React, { Component } from "react";
import { Container, Content, Spinner } from "native-base";
import { Image, StyleSheet } from "react-native";
import { isEmpty, get } from "lodash";
import AppHeader from "../components/AppHeader";
import { getChannel } from "../api/youtube";
import Config from "../config";

export default class About extends Component {
  state = {
    channel: {}
  };

  componentDidMount() {
    getChannel().then(channel => this.setState({ channel }));
  }

  renderAbout = () => {
    const { brandingSettings } = this.state.channel;
    return <Image source={{ uri: brandingSettings.image.bannerImageUrl }} style={styles.image} />;
  };

  render() {
    const { channel } = this.state;
    const title = get(channel, "snippet.title", "About");

    return (
      <>
        <AppHeader title={title} {...this.props} />
        <Container>
          <Content padder>
            {isEmpty(channel) ? <Spinner color={Config.primaryColor} /> : this.renderAbout()}
          </Content>
        </Container>
      </>
    );
  }
}

const styles = StyleSheet.create({
  image: {
    width: 100,
    height: 100
  }
});
