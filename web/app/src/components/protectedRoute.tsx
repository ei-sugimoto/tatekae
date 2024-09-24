import React from 'react';
import { Navigate } from 'react-router-dom';
import { CheckExistToken } from '../utils/checkToken';
import { Box, Text } from '@chakra-ui/react';
import { useAtom } from 'jotai';
import { MeAtom } from '../utils/meAtom';

interface ProtectedRouteProps {
  element: React.ReactElement;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ element }) => {
  const [val] = useAtom(MeAtom);
  if (!CheckExistToken() || !val) {
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
