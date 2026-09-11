import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { UpdateTodo } from '../../../lib/openapi/gen/schema'

type MutationParam = {
  id: string
} & UpdateTodo

const useUpdateTodo = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async ({ id, ...param }: MutationParam) => {
      const res = await fetch(`/todos/${id}`, {
        method: 'PATCH',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(param),
      })
      if (!res.ok) throw new Error(`HTTP error! status: ${res.status}`)
    },
    onSuccess: () => {
      // todoList の query のキャッシュを更新
      queryClient.invalidateQueries({ queryKey: ['todoList'] })
    },
  })
}
export default useUpdateTodo
