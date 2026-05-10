import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || '',
});

api.interceptors.request.use((config) => {
  const apiKey = import.meta.env.VITE_API_KEY;
  if (apiKey) {
    config.headers['X-API-Key'] = apiKey;
  }
  return config;
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
