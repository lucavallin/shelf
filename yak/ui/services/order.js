import http from '../util/http';

export async function createOrder(days) {
  return http.get(`/order/${days}`);
}
