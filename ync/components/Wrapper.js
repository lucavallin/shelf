import React from "react";
import { StyleSheet, View } from "react-native";
import { Constants } from "expo";
import { Container } from "native-base";

export default (Wrapper = ({ children }) => (
  <Container>
    <View style={styles.statusBar} />
    {children}
  </Container>
));

const styles = StyleSheet.create({
  statusBar: {
    backgroundColor: "red",
    height: Constants.statusBarHeight
  }
});
