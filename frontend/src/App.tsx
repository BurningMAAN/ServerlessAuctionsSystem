import { BrowserRouter, Route, Switch } from 'react-router-dom';
import  Dashboard from './Pages/dashboard';
import MyInventory from './Pages/myInventory';
import AuctionView from './Pages/auction';
import MyAuctions from './Pages/myAuctions';
import AuthenticationImage from './Pages/login';
import { NotificationsProvider } from '@mantine/notifications';
import BugalterDashboard from './Pages/buhalter';

function App() {
  return (
    <NotificationsProvider>
    <div className="App">
      <header className="App-header">
        <BrowserRouter>
          <Switch>
            <Route path="/login" component={AuthenticationImage}></Route>
            <Route path="/myInventory" component={MyInventory}></Route>
            <Route path="/generateData" component={BugalterDashboard}></Route>
            <Route path="/myAuctions" component={MyAuctions}></Route>
            <Route path="/auctions/:auctionID" component={AuctionView}></Route>
            <Route path="/" component={Dashboard}></Route>
          </Switch>
        </BrowserRouter>
      </header>
    </div>
    </NotificationsProvider>
  );
}

export default App;
