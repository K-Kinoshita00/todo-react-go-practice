import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  Select,
  MenuItem,
  Typography,
} from '@mui/material'
import { useState, useCallback } from 'react'
import type { Todo } from '../../../lib/openapi/gen/schema'
import useUpdateTodo from '../hooks/useUpdateTodo'

type UpdateTodoProps = {
  open: boolean
  onClose: () => void
  todo: Todo
}

const UpdateTodo = ({
  open,
  onClose,
  todo,
}: UpdateTodoProps): React.JSX.Element => {
  const [title, setTitle] = useState(todo.title)
  const [status, setStatus] = useState(todo.status)
  const [error, setError] = useState<string | null>(null)

  const { mutateAsync: updateTodo, isPending } = useUpdateTodo()

  const handleUpdate = useCallback(async () => {
    if (title.trim() === '') {
      setError('タイトルが未入力です')
      return
    }
    try {
      await updateTodo({ id: todo.id, title, status })
      onClose()
    } catch (e) {
      console.error(e)
      setError(String(e))
    }
  }, [title, status, todo.id, updateTodo, onClose])

  return (
    <Dialog open={open} onClose={onClose} fullWidth>
      <DialogTitle>Todo 更新</DialogTitle>
      <DialogContent>
        <TextField
          fullWidth
          label='Title'
          sx={{ marginY: 2 }}
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
        <Select
          fullWidth
          value={status}
          onChange={(e) => setStatus(e.target.value)}
          sx={{ marginY: 2 }}
        >
          <MenuItem value='not_started'>NOT_STARTED</MenuItem>
          <MenuItem value='in_progress'>IN_PROGRESS</MenuItem>
          <MenuItem value='completed'>COMPLETED</MenuItem>
          <MenuItem value='archive'>ARCHIVE</MenuItem>
        </Select>
        {error && <Typography color='error'>{error}</Typography>}
      </DialogContent>
      <DialogActions>
        <Button
          variant='contained'
          color='inherit'
          onClick={onClose}
          disabled={isPending}
        >
          キャンセル
        </Button>
        <Button
          variant='contained'
          color='primary'
          onClick={handleUpdate}
          disabled={isPending}
        >
          更新
        </Button>
      </DialogActions>
    </Dialog>
  )
}
export default UpdateTodo
