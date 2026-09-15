import { useQuery } from '@tanstack/react-query'
import type { ResponseTodoWithPagination } from '../../../lib/openapi/gen/schema'
import { axios } from '../../../lib/axios'

const useTodoList = () => {
  return useQuery({
    queryKey: ['todoList'],
    queryFn: async () => {
      const res = await axios.get('/todos')
      const body: ResponseTodoWithPagination['content']['application/json'] =
        res.data
      return body
    },
  })
}
export default useTodoList
