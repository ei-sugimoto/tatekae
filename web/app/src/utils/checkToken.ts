export const CheckExistToken = (): boolean => {
  const token = localStorage.getItem('token');
  if (token) {
    return true;
  }
  return false;
};
