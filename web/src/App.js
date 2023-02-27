import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';

import SpotHeader from 'Layouts/Header/spotHeader';
import OrdersHeader from 'Layouts/Header/ordersHeader';
import Footer from 'Layouts/Footer/index';
import Spot from 'Pages/spot';
import Orders from 'Pages/orders';

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path='/' element={<><SpotHeader /><Spot /></>}></Route>
        <Route exact path='/orders' element={<><OrdersHeader /><Orders /></>}></Route>
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;
