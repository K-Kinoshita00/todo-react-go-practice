import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Typography,
} from '@mui/material'
import { useCallback, useState } from 'react'
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
  const [error, setError] = useState<string | null>(null)
  const { mutateAsync: deleteTodo, isPending } = useDeleteTodo()

  const handleDelete = useCallback(async () => {
    try {
      await deleteTodo(todo.id)
      onClose()
    } catch (e) {
        console.error(e)
        setError(String(e))
    }
  }, [todo.id, deleteTodo, onClose])

  return (
    <Dialog open={open} onClose={onClose} fullWidth>
      <DialogTitle>Todo 削除</DialogTitle>
      <DialogContent>
        <Typography>Todo {todo.title} を削除しますか？</Typography>
        {error && <Typography color='error'>{error}</Typography>}
      </DialogContent>
      <DialogActions>
        <Button variant='contained' color='inherit' onClick={onClose} disabled={isPending}>
          キャンセル
        </Button>
        <Button variant='contained' color='error' onClick={handleDelete} disabled={isPending} >
          削除
        </Button>
      </DialogActions>
    </Dialog>
  )
}
export default DeleteTodo
