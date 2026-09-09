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
import type { CreateTodo as CreateTodoParam } from '../../lib/openapi/gen/schema'

const createTodo = async (param: CreateTodoParam): Promise<Response> => {
  return await fetch('/todos', {
    method: 'POST',
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

type CreateTodoProps = {
  open: boolean
  onClose: () => void
  onCreated: () => void
}

const CreateTodo = ({
  open,
  onClose,
  onCreated,
}: CreateTodoProps): React.JSX.Element => {
  const [title, setTitle] = useState('')
  const [status, setStatus] = useState<CreateTodoParam['status']>('not_started')
  const [error, setError] = useState<string | null>(null)

  const handleCreate = useCallback(async () => {
    if (title.trim() === '') {
      setError('タイトルが未入力です')
      return
    }
    const param: CreateTodoParam = {
      title,
      status,
    }
    await createTodo(param)
    await onCreated()
  }, [title, status, onCreated])

  return (
    <Dialog open={open} onClose={onClose} fullWidth>
      <DialogTitle>Todo 作成</DialogTitle>
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
          Cancel
        </Button>
        <Button variant='contained' color='primary' onClick={handleCreate}>
          Create
        </Button>
      </DialogActions>
    </Dialog>
  )
}
export default CreateTodo
