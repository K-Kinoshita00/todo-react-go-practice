import './App.css'
import { Typography } from '@mui/material'
import TodoList from './features/todos/list'

function App() {
  return (
    <>
      <Typography variant='h1'>Todo List</Typography>
      <TodoList />
    </>
  )
}

export default App
