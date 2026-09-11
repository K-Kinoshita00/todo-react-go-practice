import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Typography,
} from '@mui/material'
import { useCallback } from 'react'
import type { Todo } from '../../../lib/openapi/gen/schema'
import useDeleteTodo from '../hooks/useDeleteTodo'

type DeleteTodoProps = {
  open: boolean
  onClose: () => void
  todo: Todo
}

const DeleteTodo = ({
  open,
  onClose,
  todo,
}: DeleteTodoProps): React.JSX.Element => {
  const { mutate: deleteTodo, error: deleteTodoError } = useDeleteTodo()

  const handleDelete = useCallback(async () => {
    await deleteTodo(todo.id)
    if (deleteTodoError) {
      console.error(deleteTodoError)
      return
    }
    onClose()
  }, [todo.id, deleteTodo, onClose, deleteTodoError])

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
