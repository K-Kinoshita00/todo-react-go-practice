import { useState } from 'react'
import type { Todo } from '../../../lib/openapi/gen/schema'
import { Card, CardContent, Typography, Button } from '@mui/material'
import CreateTodo from './CreateTodo'
import UpdateTodo from './UpdateTodo'
import DeleteTodo from './DeleteTodo'
import useTodoList from '../hooks/useTodoList'

const TodoList = (): React.JSX.Element => {
  const { data: todos, isLoading, isError, error } = useTodoList()

  const [openCreateTodo, setOpenCreateTodo] = useState(false)
  const [selectedTodo, setSelectedTodo] = useState<Todo | null>(null)
  const [deleteTodo, setDeleteTodo] = useState<Todo | null>(null)

  if (isError) {
    return <Typography color='error'>{error?.message}</Typography>
  }
  return (
    <>
      <Button type='button' onClick={() => setOpenCreateTodo(true)}>
        Todo 作成
      </Button>
      {isLoading ? (
        <Typography>loading...</Typography>
      ) : (
        todos?.data.map((todo) => (
          <Card key={todo.id} sx={{ margin: 2 }}>
            <CardContent>
              <Typography variant='h6'>{todo.title}</Typography>
              <Typography>{todo.status}</Typography>
              <Button type='button' onClick={() => setSelectedTodo(todo)}>
                編集
              </Button>
              <Button
                type='button'
                color='error'
                onClick={() => setDeleteTodo(todo)}
              >
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
        />
      )}
      {selectedTodo && (
        <UpdateTodo
          open={!!selectedTodo}
          onClose={() => setSelectedTodo(null)}
          todo={selectedTodo}
        />
      )}
      {deleteTodo && (
        <DeleteTodo
          open={!!deleteTodo}
          onClose={() => setDeleteTodo(null)}
          todo={deleteTodo}
        />
      )}
    </>
  )
}

export default TodoList
