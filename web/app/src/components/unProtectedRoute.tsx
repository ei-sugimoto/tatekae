import { Navigate } from 'react-router-dom';
import { CheckExistToken } from '../utils/checkToken';

interface ProtectedRouteProps {
  element: React.ReactElement;
}

const UnProtectedRoute: React.FC<ProtectedRouteProps> = ({ element }) => {
  if (CheckExistToken()) {
    return <Navigate to='/dashboard' replace />;
  }
  return element;
};

export default UnProtectedRoute;
