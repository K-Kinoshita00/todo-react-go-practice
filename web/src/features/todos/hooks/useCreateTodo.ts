import { useMutation, useQueryClient } from '@tanstack/react-query'
import type { CreateTodo } from '../../../lib/openapi/gen/schema'

const useCreateTodo = () => {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (param: CreateTodo) => {
      const res = await fetch('/todos', {
        method: 'POST',
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
export default useCreateTodo
