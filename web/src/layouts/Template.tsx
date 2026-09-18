import { Box, Stack, Typography } from '@mui/material'
import { useLocation } from 'react-router'
import Sidebar from './Sidebar'

type TemplateProps = { children: React.ReactNode }

const Template = ({ children }: TemplateProps) => {
  const location = useLocation()
  if (location.pathname.replaceAll('/', '') === 'login') {
    return children
  }
  return (
    <Box>
      <Box sx={{ bgcolor: 'grey.300' }}>
        <Box sx={{ height: '5vh' }}>
          <Typography variant='h2' align='center'>
            Todo List
          </Typography>
        </Box>
      </Box>
      <Stack
        direction='row'
        sx={{
          display: 'flex',
          height: '95vh',
          width: '100vw',
        }}
      >
        <Sidebar />
        <Box sx={{ display: 'block' }}>{children}</Box>
      </Stack>
    </Box>
  )
}

export default Template
