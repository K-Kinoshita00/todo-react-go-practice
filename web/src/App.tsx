import { Route, Routes, Navigate } from 'react-router'
import TodoPage from './pages/TodoPage'
import LoginPage from './pages/LoginPage'
import RequireAuth from './providers/auth/RequireAuth'

function App() {
  return (
    <Routes>
      <Route element={<RequireAuth />}>
        <Route path='/' element={<Navigate to='/todos' replace />} />
        <Route path='/todos' element={<TodoPage />} />
      </Route>
      <Route path='/login' element={<LoginPage />} />
    </Routes>
  )
}

export default App
