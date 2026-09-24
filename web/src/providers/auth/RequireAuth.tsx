import { Navigate, Outlet } from 'react-router'
import { getToken } from '.'

const RequireAuth = (): React.ReactElement => {
  if (getToken() === null) {
    return <Navigate to='/login' replace />
  }
  return <Outlet />
}

export default RequireAuth
