import { API } from './api';

const post = async (url, data) => {
  const res = await API.post(url, data);
  return res.data
};

export const deerService = {
  post,
};
