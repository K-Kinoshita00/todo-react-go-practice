import { Box, Typography } from '@mui/material'
import { useLocation } from 'react-router'

type TemplateProps = { children: React.ReactNode }

const Template = ({ children }: TemplateProps) => {
  const location = useLocation()
  if (location.pathname === '/login/') {
    return children
  }
  return (
    <Box
      sx={{
        marginX: 20,
        marginY: 3,
      }}
    >
      <Typography variant='h1' align='center'>
        Todo List
      </Typography>
      {children}
    </Box>
  )
}

export default Template
