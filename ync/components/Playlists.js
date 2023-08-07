import React from "react";
import PropTypes from "prop-types";
import { Image, TouchableHighlight, StyleSheet } from "react-native";
import { isEmpty } from "lodash";
import { Container, Content, Card, CardItem, Text, Spinner } from "native-base";
import decode from "decode-html";
import Config from "../config";

export default (Playlists = ({ navigation, playlists }) => (
  <Container>
    <Content padder>
      {isEmpty(playlists) ? (
        <Spinner color={Config.primaryColor} />
      ) : (
        playlists.map(playlist => (
          <TouchableHighlight
            button
            key={playlist.id}
            onPress={() => navigation.navigate("Playlist", { youtubeId: playlist.id })}
            underlayColor="#ffffff00"
          >
            <Card>
              <CardItem cardBody>
                <Image
                  source={{ uri: playlist.snippet.thumbnails.medium.url }}
                  style={{ height: 180, width: null, flex: 1 }}
                />
              </CardItem>
              <CardItem>
                <Text style={styles.text}>{decode(playlist.snippet.title)}</Text>
              </CardItem>
            </Card>
          </TouchableHighlight>
        ))
      )}
    </Content>
  </Container>
));

Playlists.propTypes = {
  playlists: PropTypes.array.isRequired
};

const styles = StyleSheet.create({
  image: {
    maxWidth: 150
  },
  textBox: {
    padding: 10
  },
  text: {
    fontWeight: "bold"
  }
});
