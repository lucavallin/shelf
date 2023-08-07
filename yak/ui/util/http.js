const axios = require('axios');

const instance = axios.create({
  // This should be moved to configuration
  baseURL: 'http://localhost/yak-webshop',
  timeout: 1000,
});

module.exports = instance;
