import qs from 'qs';
import { API } from './api';

const get = async (url, data) => {
  const res = await API.get(`${url}?${qs.stringify(data)}`);
  return res.data;
};

// 获取分页列表
const getPageList = async (url, data) => {
  if (data.page_size === undefined) data.page_size = 10;
  if (data.p === undefined) data.p = 1;
  return get(url, data);
};

const post = async (url, data) => {
  const res = await API.post(url, data);
  return res.data;
};

export const deerService = {
  get,
  getPageList,
  post,
};
