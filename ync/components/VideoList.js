import React from "react";
import PropTypes from "prop-types";
import { TouchableHighlight, Image, StyleSheet } from "react-native";
import { isEmpty } from "lodash";
import { Text, Content, Container, Card, CardItem, Left, Right, Spinner } from "native-base";
import decode from "decode-html";
import Config from "../config";

export default (VideoList = ({ navigation, videos }) => (
  <Container>
    <Content padder>
      {isEmpty(videos) ? (
        <Spinner color={Config.primaryColor} />
      ) : (
        videos.map(video => (
          <TouchableHighlight
            button
            key={video.id.videoId}
            onPress={() => navigation.navigate("Video", { youtubeId: video.id.videoId })}
            underlayColor="#ffffff00"
          >
            <Card>
              <CardItem cardBody>
                <Left style={styles.image}>
                  <Image
                    source={{ uri: video.snippet.thumbnails.medium.url }}
                    style={{ height: 100, width: null, flex: 1 }}
                  />
                </Left>
                <Right style={styles.textBox}>
                  <Text style={styles.text}>{decode(video.snippet.title)}</Text>
                </Right>
              </CardItem>
            </Card>
          </TouchableHighlight>
        ))
      )}
    </Content>
  </Container>
));

VideoList.propTypes = {
  videos: PropTypes.array.isRequired
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
