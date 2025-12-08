import qs from 'qs';
import { API } from './api';

const get = async (url, data) => {
  const res = await API.get(`${url}?${qs.stringify(data)}`);
  return res.data
};

const post = async (url, data) => {
  const res = await API.post(url, data);
  return res.data
};

export const deerService = {
  get,
  post,
};
