import { Typography } from '@mui/material'
import TodoList from '../features/todos/components/TodoList'

const TodoPage = () => {
  return (
    <>
      <Typography variant='h1'>Todo List</Typography>
      <TodoList />
    </>
  )
}

export default TodoPage
