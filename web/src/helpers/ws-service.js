import qs from 'qs';
import { API } from './api';

const get = async (url, data) => {
  const res = await API.get(`${url}?${qs.stringify(data)}`);
  return res.data;
};

// 获取分页列表
const getPageList = async (url, { pageSize, pageNumber, ...data }) => {
  data.page_size = pageSize || 10;
  data.p = pageNumber || 1;
  const res = await get(url, data);
  const { items, page_size, total, page } = res.data;
  Object.assign(res.data, {
    list: items,
    total,
    pageSize: page_size,
    pageNumber: page,
  });
  return res;
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
