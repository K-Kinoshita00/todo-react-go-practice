import { Box, Button } from '@mui/material'
import { useCallback } from 'react'
import { clearToken } from '../providers/auth'
import { useNavigate } from 'react-router'

const Sidebar = () => {
  const navigate = useNavigate()

  const signOut = useCallback(() => {
    clearToken()
    navigate('/login')
  }, [navigate])

  return (
    <Box
      sx={{
        bgcolor: 'grey.300',
        width: '10vw',
        height: '100%',
        display: 'block',
        padding: 1,
      }}
    >
      <Button onClick={signOut}>ログアウト</Button>
    </Box>
  )
}

export default Sidebar
