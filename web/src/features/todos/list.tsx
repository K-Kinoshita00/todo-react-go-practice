import { useState, useEffect } from 'react'
import type {
  Todo,
  ResponseTodoWithPagination,
} from '../../lib/openapi/gen/schema'
import { Card, CardContent, Typography } from '@mui/material'

type TodoListBody = ResponseTodoWithPagination['content']['application/json']

function useTodoList() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  useEffect(() => {
    fetch('/todos')
      .then((res) => {
        if (!res.ok) {
          throw new Error(`HTTP error! status: ${res.status}`)
        }
        return res.json()
      })
      .then((body: TodoListBody) => {
        setTodos(body.data)
      })
      .catch((e) => {
        setError(e.message)
      })
      .finally(() => {
        setLoading(false)
      })
  }, []) // 1 回だけ
  return { todos, loading, error }
}

const TodoList: React.FC = () => {
  const { todos, loading, error } = useTodoList()
  if (error) {
    return <Typography color='error'>{error}</Typography>
  }
  return (
    <>
      {loading ? (
        <Typography>loading...</Typography>
      ) : (
        todos.map((todo) => (
          <Card key={todo.id} sx={{ margin: 2 }}>
            <CardContent>
              <Typography variant='h6'>{todo.title}</Typography>
              <Typography>{todo.status}</Typography>
            </CardContent>
          </Card>
        ))
      )}
    </>
  )
}

export default TodoList
