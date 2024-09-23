import * as React from 'react';
import { ChakraProvider } from '@chakra-ui/react';
import * as ReactDOM from 'react-dom/client';
import App from './App';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { DashBoard } from './pages/dashBeard';
import ProtectedRoute from './components/protectedRoute';
import { ProjectDetail } from './pages/projectDetail';
import UnProtectedRoute from './components/unProtectedRoute';
import { Provider } from 'jotai';
import { store } from './utils/meAtom';

const rootElement = document.getElementById('root');

const router = createBrowserRouter([
  {
    path: '/',
    element: <UnProtectedRoute element={<App />} />,
  },
  {
    path: '/dashboard',
    element: <ProtectedRoute element={<DashBoard />} />,
  },
  {
    path: '/project/:id',
    element: <ProtectedRoute element={<ProjectDetail />} />,
  },
]);

if (rootElement) {
  ReactDOM.createRoot(rootElement).render(
    <React.StrictMode>
      <ChakraProvider>
        <Provider store={store}>
          <RouterProvider router={router} />
        </Provider>
      </ChakraProvider>
    </React.StrictMode>
  );
} else {
  console.error('Failed to find the root element');
}
