import { Navigate, Outlet } from "react-router-dom";

const ProtectedRoutes = ({ isAuthenticated }) => {
  console.log(isAuthenticated);
  return isAuthenticated ? <Outlet /> : <Navigate to="/" />;
};

export default ProtectedRoutes;