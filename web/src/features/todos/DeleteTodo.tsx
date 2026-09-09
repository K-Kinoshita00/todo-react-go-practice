import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Typography,
} from '@mui/material'
import { useCallback } from 'react'
import type { Todo } from '../../lib/openapi/gen/schema'

const deleteTodo = async (id: string): Promise<Response> => {
  return await fetch(`/todos/${id}`, {
    method: 'DELETE',
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error(`HTTP error! status: ${res.status}`)
      }
      return res
    })
    .catch((e) => {
      throw new Error(e.message)
    })
}

type DeleteTodoProps = {
  open: boolean
  onClose: () => void
  onDeleted: () => void
  todo: Todo
}

const DeleteTodo = ({
  open,
  onClose,
  onDeleted,
  todo,
}: DeleteTodoProps): React.JSX.Element => {
  const handleDelete = useCallback(async () => {
    await deleteTodo(todo.id)
    await onDeleted()
  }, [todo.id, onDeleted])

  return (
    <Dialog open={open} onClose={onClose} fullWidth>
      <DialogTitle>Todo 削除</DialogTitle>
      <DialogContent>
        <Typography>Todo {todo.title} を削除しますか？</Typography>
      </DialogContent>
      <DialogActions>
        <Button variant='contained' color='inherit' onClick={onClose}>
          キャンセル
        </Button>
        <Button variant='contained' color='error' onClick={handleDelete}>
          削除
        </Button>
      </DialogActions>
    </Dialog>
  )
}
export default DeleteTodo
