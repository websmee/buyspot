import { useRoutes } from 'react-router-dom';

import Spot from 'Pages/spot';
import Orders from 'Pages/orders';
import Settings from 'Pages/settings';
import Profile from 'Pages/profile';
import Layout from 'Layouts/layout';

function App() {
  const routes = [
    {
      path: '/',
      element: <Layout />,
      children: [
        { index: true, element: <Spot /> },
        { path: '/orders', element: <Orders /> },
        { path: '/settings', element: <Settings /> },
        { path: '/profile', element: <Profile /> },
      ],
    },
  ];

  return useRoutes(routes);
}

export default App;
