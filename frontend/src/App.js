import {lazy, Suspense} from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route
} from 'react-router-dom';
import { useSelector } from 'react-redux'
import './App.scss';
import routes from './routes/routes';
import ProtectedRoutes from './routes/ProtectedRoutes';

const LoginPage = lazy(() => import('./features/OnBoarding/LoginPage'));

function App() {
  //Getting isAuthenticated store value from Authentication slice.
  const isAuthenticated = useSelector((state) => state.authentication.isAuthenticated)
  console.log('app', isAuthenticated)
  return (
    <Router>
      <Suspense fallback={<div>Loading...</div>}>
        <Routes>
            {/* <Route path="*" element={<NotFoundPage />} /> */}
              <Route path="/" element={<LoginPage />} />
              <Route element={<ProtectedRoutes isAuthenticated={isAuthenticated} />}>
                {routes.map(({component: Component, path, exact}, index) => (
                  <Route path={`/${path}`} key={index} exact={exact} element={<Component />}/>
                ))}
              </Route>
          </Routes>
        </Suspense>
  </Router>
  );
}

export default App;
