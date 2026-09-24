import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { CreateTodo } from '../../../lib/openapi/gen/schema'
import { axios } from '../../../lib/axios'

const useCreateTodo = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (param: CreateTodo) => {
      await axios.post('/todos', param)
    },
    onSuccess: () => {
      // todoList の query のキャッシュを更新
      queryClient.invalidateQueries({
        queryKey: ['todoList'],
      })
    },
  })
}
export default useCreateTodo
