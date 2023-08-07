import React from "react";
import { StyleSheet } from "react-native";
import { Card } from "native-base";
import YouTube from "react-native-youtube";

export default class Video extends React.Component {
  render() {
    return (
      <Card style={styles.container}>
        <YouTube
          videoId={this.props.navigation.state.params.youtubeId}
          play
          fullscreen
          loop={false}
          apiKey="blabla"
          onReady={e => this.setState({ isReady: true })}
          onChangeState={e => this.setState({ status: e.state })}
          onChangeQuality={e => this.setState({ quality: e.quality })}
          onError={e => this.setState({ error: e.error })}
          style={{ alignSelf: "stretch", height: 300 }}
        />
      </Card>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#fff"
  }
});
