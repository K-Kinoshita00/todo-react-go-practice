import { Route, Routes } from 'react-router'
import TodoPage from './pages/TodoPage'
import LoginPage from './pages/LoginPage'
import RequireAuth from './providers/auth/RequireAuth'

function App() {
  return (
    <Routes>
      <Route element={<RequireAuth />}>
        <Route path='/' element={<TodoPage />} />
      </Route>
      <Route path='/login' element={<LoginPage />} />
    </Routes>
  )
}

export default App
