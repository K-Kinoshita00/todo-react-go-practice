import { useState, useEffect, useCallback } from 'react'
import type {
  Todo,
  ResponseTodoWithPagination,
} from '../../lib/openapi/gen/schema'
import { Card, CardContent, Typography, Button } from '@mui/material'
import CreateTodo from './CreateTodo'
import UpdateTodo from './UpdateTodo'
import DeleteTodo from './DeleteTodo'

type TodoListBody = ResponseTodoWithPagination['content']['application/json']

const useTodoList = () => {
  const [todos, setTodos] = useState<Todo[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  const refetchTodoList = async (): Promise<void> => {
    return await fetch('/todos')
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
  }

  useEffect(() => {
    refetchTodoList()
  }, []) // 1 回だけ
  return { todos, loading, error, refetch: refetchTodoList }
}

const TodoList = (): React.JSX.Element => {
  const { todos, loading, error, refetch } = useTodoList()

  const [openCreateTodo, setOpenCreateTodo] = useState(false)
  const [selectedTodo, setSelectedTodo] = useState<Todo | null>(null)
  const [deleteTodo, setDeleteTodo] = useState<Todo | null>(null)

  const onCreated = useCallback(async () => {
    await refetch()
    setOpenCreateTodo(false)
  }, [refetch, setOpenCreateTodo])

  const onUpdated = useCallback(async () => {
    await refetch()
    setSelectedTodo(null)
  }, [refetch, setSelectedTodo])

  const onDeleted = useCallback(async () => {
    await refetch()
    setDeleteTodo(null)
  }, [refetch, setDeleteTodo])

  if (error) {
    return <Typography color='error'>{error}</Typography>
  }
  return (
    <>
      <Button type='button' onClick={() => setOpenCreateTodo(true)}>
        Todo 作成
      </Button>
      {loading ? (
        <Typography>loading...</Typography>
      ) : (
        todos.map((todo) => (
          <Card key={todo.id} sx={{ margin: 2 }}>
            <CardContent>
              <Typography variant='h6'>{todo.title}</Typography>
              <Typography>{todo.status}</Typography>
              <Button type='button' onClick={() => setSelectedTodo(todo)}>
                編集
              </Button>
              <Button type='button' color='error' onClick={() => setDeleteTodo(todo)}>
                削除
              </Button>
            </CardContent>
          </Card>
        ))
      )}
      {openCreateTodo && (
        <CreateTodo
          open={openCreateTodo}
          onClose={() => setOpenCreateTodo(false)}
          onCreated={onCreated}
        />
      )}
      {selectedTodo && (
        <UpdateTodo
          open={!!selectedTodo}
          onClose={() => setSelectedTodo(null)}
          onUpdated={onUpdated}
          todo={selectedTodo}
        />
      )}
      {deleteTodo && (
        <DeleteTodo
          open={!!deleteTodo}
          onClose={() => setDeleteTodo(null)}
          onDeleted={onDeleted}
          todo={deleteTodo}
        />
      )}
    </>
  )
}

export default TodoList
