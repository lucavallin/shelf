import React from "react";
import { Root } from "native-base";
import { createDrawerNavigator } from "react-navigation-drawer";
import { createAppContainer } from "react-navigation";

import SideBar from "./components/Sidebar";
import Home from "./screens/Home";
import Search from "./screens/Search";
import Latest from "./screens/Latest";
import Popular from "./screens/Popular";
import Playlists from "./screens/Playlists";
import About from "./screens/About";
import Video from "./screens/secondary/Video";
import Playlist from "./screens/secondary/Playlist";

const AppContainer = createAppContainer(
  createDrawerNavigator(
    {
      Home: { screen: Home },
      Search: { screen: Search },
      Latest: { screen: Latest },
      Popular: { screen: Popular },
      Playlists: { screen: Playlists },
      About: { screen: About },
      Video: { screen: Video },
      Playlist: { screen: Playlist }
    },
    {
      initialRouteName: "Home",
      contentComponent: props => <SideBar {...props} />
    }
  )
);

export default () => (
  <Root>
    <AppContainer />
  </Root>
);
