import React from "react";
import { Platform, Dimensions, StyleSheet, Linking } from "react-native";
import { Content, Text, List, ListItem, Icon, Left } from "native-base";
import { channelUrl } from "../config";

const links = [
  {
    name: "Ricerca",
    route: "Search",
    icon: "search"
  },
  {
    name: "Ultimi video",
    route: "Latest",
    icon: "videocam"
  },
  {
    name: "Più popolari",
    route: "Popular",
    icon: "star"
  },
  {
    name: "Playlists",
    route: "Playlists",
    icon: "list-box"
  },
  {
    name: "Iscriviti al canale",
    route: () => Linking.openURL(config.channelUrl),
    icon: "logo-youtube"
  },
  {
    name: "Chi siamo",
    route: "About",
    icon: "help"
  }
];

export default (Sidebar = ({ navigation }) => (
  <Content bounces={false} style={{ flex: 1, backgroundColor: "#fff", top: -1 }}>
    <List
      dataArray={links}
      renderRow={link => (
        <ListItem
          button
          noBorder
          onPress={
            typeof link.route === "function" ? link.route : () => navigation.navigate(link.route)
          }
        >
          <Left>
            <Icon
              active
              name={link.icon}
              style={{
                color: "#777",
                fontSize: 26,
                width: 30,
                textAlign: "center"
              }}
            />
            <Text style={styles.text}>{link.name}</Text>
          </Left>
        </ListItem>
      )}
    />
  </Content>
));

const deviceHeight = Dimensions.get("window").height;
const deviceWidth = Dimensions.get("window").width;
const styles = StyleSheet.create({
  drawerCover: {
    alignSelf: "stretch",
    height: deviceHeight / 3.5,
    width: null,
    position: "relative",
    marginBottom: 10
  },
  drawerImage: {
    position: "absolute",
    left: Platform.OS === "android" ? deviceWidth / 10 : deviceWidth / 9,
    top: Platform.OS === "android" ? deviceHeight / 13 : deviceHeight / 12,
    width: 210,
    height: 75,
    resizeMode: "cover"
  },
  text: {
    fontWeight: Platform.OS === "ios" ? "500" : "400",
    fontSize: 16,
    marginLeft: 20
  },
  badgeText: {
    fontSize: Platform.OS === "ios" ? 13 : 11,
    fontWeight: "400",
    textAlign: "center",
    marginTop: Platform.OS === "android" ? -3 : undefined
  }
});
