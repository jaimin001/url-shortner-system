import axios from 'axios';

const api = axios.create({
  baseURL: '',
});

export const shortenURL = async (url: string) => {
  const { data } = await api.post('/shorten', { url: url });
  return data;
};

export const fetchRecentLinks = async () => {
  const { data } = await api.get('/links');
  return data;
};

export default api;