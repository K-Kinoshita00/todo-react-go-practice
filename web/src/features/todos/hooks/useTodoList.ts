import { useQuery } from '@tanstack/react-query'
import type { ResponseTodoWithPagination } from '../../../lib/openapi/gen/schema'

const useTodoList = () => {
  return useQuery({
    queryKey: ['todoList'],
    queryFn: async () => {
      const res = await fetch('/todos', {
        method: 'GET',
      })
      if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`)
      const body: ResponseTodoWithPagination['content']['application/json'] =
        await res.json()
      return body
    },
  })
}
export default useTodoList
