import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { UpdateTodo } from '../../../lib/openapi/gen/schema'
import { axios } from '../../../lib/axios'

type MutationParam = {
  id: string
} & UpdateTodo

const useUpdateTodo = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async ({ id, ...param }: MutationParam) => {
      await axios.patch(`/todos/${id}`, param)
    },
    onSuccess: () => {
      // todoList の query のキャッシュを更新
      queryClient.invalidateQueries({ queryKey: ['todoList'] })
    },
  })
}
export default useUpdateTodo
