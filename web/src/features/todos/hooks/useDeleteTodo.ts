import { useMutation, useQueryClient } from '@tanstack/react-query'

const useDeleteTodo = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (id: string) => {
      const res = await fetch(`/todos/${id}`, {
        method: 'DELETE',
      })
      if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`)
    },
    onSuccess: () => {
      // todoList の query のキャッシュを更新
      queryClient.invalidateQueries({ queryKey: ['todoList'] })
    },
  })
}
export default useDeleteTodo
