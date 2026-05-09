import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080',
});

export const shortenURL = async (url: string) => {
  const { data } = await api.post('/shorten', { url });
  return data;
};

export const fetchRecentLinks = async () => {
  // Assuming a hypothetical endpoint for list
  const { data } = await api.get('/links');
  return data;
};

export default api;