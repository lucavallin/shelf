import React from "react";
import PropTypes from "prop-types";
import { Header, Title, Left, Icon, Right, Button, Body } from "native-base";

export default (AppHeader = ({ navigation, title }) => (
  <Header>
    <Left>
      <Button transparent onPress={() => navigation.openDrawer()}>
        <Icon name="menu" />
      </Button>
    </Left>
    <Body>
      <Title>{title}</Title>
    </Body>
    <Right />
  </Header>
));

AppHeader.propTypes = {
  title: PropTypes.string.isRequired
};
