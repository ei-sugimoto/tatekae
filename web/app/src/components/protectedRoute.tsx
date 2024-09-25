import React from 'react';
import { Navigate } from 'react-router-dom';
import { CheckExistToken } from '../utils/checkToken';
import { Box, Text } from '@chakra-ui/react';
import { useAtom } from 'jotai';
import { MeAtom } from '../utils/meAtom';
import { Client } from '../use-client';
import { UserService } from '../gen/proto_user/v1/user_connect';
import { deleteCookie, getCookie } from '../utils/cookie';
import { ConnectError } from '@connectrpc/connect';

interface ProtectedRouteProps {
  element: React.ReactElement;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ element }) => {
  const [val, setVal] = useAtom(MeAtom);

  const setAuth = async () => {
    const res = await GetAuth();
    if (res instanceof ConnectError) {
      deleteCookie('token');
      deleteCookie('id');
      return;
    }
    setVal((prev) => ({ ...prev, id: res.id, name: res.username }));
  };

  if (!val && CheckExistToken()) {
    setAuth();
  }

  if (!CheckExistToken()) {
    return <Navigate to='/' replace />;
  }
  return (
    <>
      <Box position={'fixed'} top={0} right={0}>
        <Text fontSize={'xl'}>username : {val?.name ?? 'No name'}</Text>
      </Box>
      {element}
    </>
  );
};

export default ProtectedRoute;

const GetAuth = async () => {
  try {
    const id = getCookie('id');
    const res = await Client(UserService).getByID({ id: Number(id) });
    return res;
  } catch (e) {
    return ConnectError.from(e);
  }
};
