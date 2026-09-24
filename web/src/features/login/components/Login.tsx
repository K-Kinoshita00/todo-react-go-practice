import { Box, Button, TextField, Typography } from '@mui/material'
import { useCallback, useState } from 'react'
import { signIn } from '../../../lib/cognito'
import { setToken } from '../../../providers/auth'
import { useNavigate } from 'react-router'

const Login = () => {
  const navigate = useNavigate()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  const handleLogin = useCallback(async () => {
    if (email === '') {
      setError('メールアドレスを入力してください')
      return
    }
    if (password === '') {
      setError('パスワードを入力してください')
      return
    }
    try {
      const token = await signIn(email, password)
      setToken(token)
      navigate('/')
    } catch {
      setError('メールアドレス もしくは パスワード が異なっています')
    }
  }, [email, password, setError, navigate])

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        marginTop: '30vh',
        marginX: 'auto',
        width: '40%',
        gap: 1,
      }}
    >
      <Typography variant='h1' align='center'>
        Todo List
      </Typography>
      <Typography variant='h5'>ログイン</Typography>
      <TextField
        type='email'
        placeholder='email'
        value={email}
        onChange={(e) => {
          setEmail(e.target.value)
        }}
        fullWidth
      />
      <TextField
        type='password'
        placeholder='password'
        value={password}
        onChange={(e) => {
          setPassword(e.target.value)
        }}
        fullWidth
      />
      <Box>
        {error !== '' && <Typography color='error'>{error}</Typography>}
        <Button onClick={handleLogin}>ログイン</Button>
      </Box>
    </Box>
  )
}

export default Login
