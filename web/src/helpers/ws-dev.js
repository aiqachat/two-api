import { API, updateAPI } from './api';
import { setUserData } from './data';

export const autoLogin = async () => {
  const turnstileToken = 11;
  const res = await API.post(`/api/user/login?turnstile=${turnstileToken}`, {
    username: 'admin',
    password: '88888888',
  });
  const { success, message, data } = res.data;
  if (!success) {
    console.error(message);
    return;
  }
  setUserData(data);
  updateAPI();
};

export const wsDev = {
  autoLogin,
};
