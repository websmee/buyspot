import { useRoutes } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { getCurrentPrices } from 'Store/reducer';
import { useEffect } from 'react';

import Spot from 'Pages/spot';
import Orders from 'Pages/orders';
import Settings from 'Pages/settings';
import Profile from 'Pages/profile';
import Layout from 'Layouts/layout';
import Login from 'Pages/login';
import runOneSignal from "Utils/onesignal";

function App() {
  const dispatch = useDispatch();

  useEffect(() => {
    runOneSignal();
    dispatch(getCurrentPrices());
    const intervalId = setInterval(() => {
        dispatch(getCurrentPrices());
    }, 60000);

    return () => {
      clearInterval(intervalId);
    };
  }, [dispatch]);

  const routes = [
    {
      path: '/',
      element: <Layout />,
      children: [
        { index: true, element: <Spot /> },
        { path: '/orders', element: <Orders /> },
        { path: '/settings', element: <Settings /> },
        { path: '/profile', element: <Profile /> },
        { path: '/login', element: <Login /> },
      ],
    },
  ];

  return useRoutes(routes);
}

export default App;
