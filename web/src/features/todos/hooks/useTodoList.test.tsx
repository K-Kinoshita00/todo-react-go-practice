import { renderHook, waitFor } from '@testing-library/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { describe, expect, it, vi } from 'vitest'
import useTodoList from './useTodoList'

const createWrapper = () => {
  const client = new QueryClient({
    defaultOptions: { queries: { retry: false } },
  })
  return ({ children }: { children: React.ReactNode }) => (
    <QueryClientProvider client={client}>{children}</QueryClientProvider>
  )
}
describe('useTodoList', () => {
  it('一覧データを返す', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({
          data: [
            {
              id: '01a08fed-a81a-725f-9c24-1c5a372d27d0',
              title: 'todo_1',
              status: 'not_started',
            },
          ],
        }),
      }),
    )
    const { result } = renderHook(() => useTodoList(), {
      wrapper: createWrapper(),
    })
    await waitFor(() => expect(result.current.isSuccess).toBe(true))
    expect(result.current.data?.data[0]?.title).toBe('todo_1')
    expect(result.current.data?.data[0]?.status).toBe('not_started')
  })

  it('空配列を返す', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({
          data: [],
        }),
      }),
    )
    const { result } = renderHook(() => useTodoList(), {
      wrapper: createWrapper(),
    })
    await waitFor(() => expect(result.current.isSuccess).toBe(true))
    expect(result.current.data?.data.length).toBe(0)
  })

  it('失敗なら isError', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({ ok: false, status: 500 }),
    )
    const { result } = renderHook(() => useTodoList(), {
      wrapper: createWrapper(),
    })
    await waitFor(() => expect(result.current.isSuccess).toBe(false))
    expect(result.current.status).toBe(500)
  })
})
