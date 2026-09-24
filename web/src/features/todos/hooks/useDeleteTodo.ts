import { useMutation, useQueryClient } from '@tanstack/react-query'
import { axios } from '../../../lib/axios'

const useDeleteTodo = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (id: string) => {
      await axios.delete(`/todos/${id}`)
    },
    onSuccess: () => {
      // todoList の query のキャッシュを更新
      queryClient.invalidateQueries({
        queryKey: ['todoList'],
      })
    },
  })
}
export default useDeleteTodo
