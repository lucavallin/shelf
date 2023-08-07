# Youtube Native Channel

Youtube Native Channel (with React Native)

## Config

Rename `.config.js` to `config.js` and enter the required configuration.

## Gotchas

- The <Wrapper /> component wraps the whole app, it gives a margin the size of the status bar (on Android) to overcome a bug in expo/react-navigation that makes the view overlap with the status bar.

## TODO

- Check API calls so that only necessary content is retrieved
- Wire up to Sentry for proper error logging
- Update all deps
- Translations
- Possibly split components
- Tests
