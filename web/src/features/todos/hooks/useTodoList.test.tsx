import { renderHook, waitFor } from '@testing-library/react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { describe, expect, it, vi } from 'vitest'
import useTodoList from './useTodoList'
import { axios } from '../../../lib/axios'

vi.mock('../../../lib/axios', () => ({
  axios: {
    get: vi.fn(),
  },
}))

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
    vi.mocked(axios.get).mockResolvedValue({
      data: {
        data: [
          {
            id: '01a08fed-a81a-725f-9c24-1c5a372d27d0',
            title: 'todo_1',
            status: 'not_started',
          },
          {
            id: '01a0a4ab-290d-75c8-8121-5769bc3905c8',
            title: 'todo_2',
            status: 'in_progress',
          },
        ],
      },
    })
    const { result } = renderHook(() => useTodoList(), {
      wrapper: createWrapper(),
    })
    await waitFor(() => expect(result.current.isSuccess).toBe(true))
    expect(result.current.data?.data[0]?.title).toBe('todo_1')
    expect(result.current.data?.data[0]?.status).toBe('not_started')
    expect(result.current.data?.data[1]?.title).toBe('todo_2')
    expect(result.current.data?.data[1]?.status).toBe('in_progress')
  })

  it('空配列を返す', async () => {
    vi.mocked(axios.get).mockResolvedValue({
      data: {
        data: [],
      },
    })
    const { result } = renderHook(() => useTodoList(), {
      wrapper: createWrapper(),
    })
    await waitFor(() => expect(result.current.isSuccess).toBe(true))
    expect(result.current.data?.data.length).toBe(0)
  })

  it('失敗なら isError', async () => {
    vi.mocked(axios.get).mockRejectedValue(new Error('HTTP 500'))
    const { result } = renderHook(() => useTodoList(), {
      wrapper: createWrapper(),
    })
    await waitFor(() => expect(result.current.isError).toBe(true))
    expect(result.current.status).toBe('error')
  })
})
