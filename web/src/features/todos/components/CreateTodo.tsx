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
import type { CreateTodo as CreateTodoParam } from '../../../lib/openapi/gen/schema'
import useCreateTodo from '../hooks/useCreateTodo'

type CreateTodoProps = {
  open: boolean
  onClose: () => void
}

const CreateTodo = ({ open, onClose }: CreateTodoProps): React.JSX.Element => {
  const [title, setTitle] = useState('')
  const [status, setStatus] = useState<CreateTodoParam['status']>('not_started')
  const [error, setError] = useState<string | null>(null)

  const { mutateAsync: createTodo, isPending } = useCreateTodo()

  const handleCreate = useCallback(async () => {
    if (title.trim() === '') {
      setError('タイトルが未入力です')
      return
    }
    const param: CreateTodoParam = {
      title,
      status,
    }
    try {
      await createTodo(param)
      onClose()
    } catch (e) {
      console.error(e)
      setError(String(e))
    }
  }, [title, status, createTodo, onClose])

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
        <Button variant='contained' color='inherit' onClick={onClose} disabled={isPending}>
          キャンセル
        </Button>
        <Button variant='contained' color='primary' onClick={handleCreate} disabled={isPending}>
          作成
        </Button>
      </DialogActions>
    </Dialog>
  )
}
export default CreateTodo
