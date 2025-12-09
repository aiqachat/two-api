import { API, updateAPI } from './api';
import { setUserData } from './data';
import { debounce } from 'lodash';

export const autoLogin = debounce(
  async () => {
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
  },
  200,
  {
    leading: true,
    trailing: false,
  },
);

export const wsDev = {
  autoLogin,
};
