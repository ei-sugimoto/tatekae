import { getCookie } from './cookie';

export const CheckExistToken = (): boolean => {
  const token = getCookie('token');
  if (token) {
    return true;
  }
  return false;
};
