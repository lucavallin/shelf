import http from '../util/http';

export async function postHerd(xml) {
  return http.post(`/load`, xml, {
    headers: { 'Content-Type': 'text/plain' },
  });
}

export async function getHerd(days) {
  return http.get(`/herd/${days}`);
}

export async function getStock(days) {
  return http.get(`/stock/${days}`);
}
