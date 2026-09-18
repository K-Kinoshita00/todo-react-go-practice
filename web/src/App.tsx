import { Route, Routes } from 'react-router'
import TodoPage from './pages/TodoPage'
import LoginPage from './pages/LoginPage'

function App() {
  return (
    <Routes>
      <Route path='/' element={<TodoPage />} />
      <Route path='/login' element={<LoginPage />} />
    </Routes>
  )
}

export default App
