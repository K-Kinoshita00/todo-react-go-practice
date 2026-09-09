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
import type { Todo, UpdateTodo as UpdateTodoParam } from '../../lib/openapi/gen/schema'

const updateTodo = async (id: string, param: UpdateTodoParam): Promise<Response> => {
  return await fetch(`/todos/${id}`, {
    method: 'PATCH',
    body: JSON.stringify(param),
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

type UpdateTodoProps = {
  open: boolean
  onClose: () => void
  onUpdated: () => void
  todo: Todo
}

const UpdateTodo = ({ open, onClose, onUpdated, todo }: UpdateTodoProps): React.JSX.Element => {
  const [title, setTitle] = useState(todo.title)
  const [status, setStatus] = useState(todo.status)
  const [error, setError] = useState<string | null>(null)

  const handleUpdate = useCallback(async () => {
    if (title.trim() === '') {
      setError('タイトルが未入力です')
      return
    }
    await updateTodo(todo.id, { title, status })
    await onUpdated()
  }, [title, status, todo.id, onUpdated])

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
        <Button variant='contained' color='inherit' onClick={onClose}>
          キャンセル
        </Button>
        <Button variant='contained' color='primary' onClick={handleUpdate}>
          更新
        </Button>
      </DialogActions>
    </Dialog>
  )
}
export default UpdateTodo