import React from 'react';
import logo from './logo.svg';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import  Dashboard from './Pages/dashboard';
import MyInventory from './Pages/myInventory';
import AuctionView from './Pages/auction';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <BrowserRouter>
          <Switch>
            <Route path="/myInventory" component={MyInventory}></Route>
            <Route path="/auctions/:id" component={AuctionView}></Route>
            <Route path="/" component={Dashboard}></Route>
          </Switch>
        </BrowserRouter>
      </header>
    </div>
  );
}

export default App;
