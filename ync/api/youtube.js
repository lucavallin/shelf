import queryString from "query-string";
import axios from "axios";
import Config from "../config";

const http = axios.create({
  baseURL: "https://www.googleapis.com/youtube/v3"
});

const createEndpoint = (endpoint, params) =>
  `/${endpoint}?key=${Config.apiKey}&${queryString.stringify(params)}`;

export const getVideos = params =>
  http
    .get(
      createEndpoint("search", {
        channelId: Config.channelId,
        part: "snippet,id",
        ...params
      })
    )
    .then(res => res.data.items.reduce((acc, video) => acc.concat([video]), []))
    .catch(err => console.log(err));

export const getPlaylists = () =>
  http
    .get(
      createEndpoint("playlists", {
        channelId: Config.channelId,
        part: "snippet,contentDetails",
        order: "date",
        maxResults: 30
      })
    )
    .then(res => res.data.items)
    .catch(err => console.log(err));

export const getChannel = () =>
  http
    .get(
      createEndpoint("channels", {
        id: Config.channelId,
        part: "snippet,brandingSettings"
      })
    )
    .then(res => res.data.items[0])
    .catch(err => console.log(err));
