import { QueryClientProvider } from '@tanstack/react-query'
import { queryClient } from './index'

const QueryProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
  )
}
export default QueryProvider
