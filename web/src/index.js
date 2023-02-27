import React from 'react';
import ReactDOM from 'react-dom/client';
import App from 'app';

import { iOSversion } from 'Services/Utils/stickymobile';

if (iOSversion().version > 14) { document.querySelectorAll('#page')[0].classList.add('min-ios15'); }

const root = ReactDOM.createRoot(document.getElementById('page'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);