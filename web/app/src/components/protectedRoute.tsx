import React from 'react';
import { Navigate } from 'react-router-dom';
import { CheckExistToken } from '../utils/checkToken';

interface ProtectedRouteProps {
  element: React.ReactElement;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ element }) => {
  if (!CheckExistToken()) {
    return <Navigate to='/' replace />;
  }
  return element;
};

export default ProtectedRoute;
