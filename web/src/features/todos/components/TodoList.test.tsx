import { render, screen } from '@testing-library/react'
import { afterEach, beforeAll, describe, expect, it, vi } from 'vitest'
import TodoList from './TodoList'
import useTodoList from '../hooks/useTodoList'

vi.mock('../hooks/useTodoList', () => ({ default: vi.fn() }))

describe('TodoList', () => {
  beforeAll(() => {
    vi.stubGlobal('fetch', vi.fn())
  })
  afterEach(() => {
    vi.unstubAllGlobals()
  })

  it('一覧を表示する', async () => {
    vi.mocked(useTodoList).mockReturnValue({
      data: {
        data: [
          {
            id: '01a08fe6-796a-776a-92be-97e022c979b9',
            title: 'todo_1',
            status: 'not_started',
          },
          {
            id: '01a08fe6-e426-75cf-86dc-7e7a5fa01162',
            title: 'todo_2',
            status: 'in_progress',
          },
        ],
      },
      isLoading: false,
      isError: false,
      error: null,
    } as unknown as ReturnType<typeof useTodoList>)
    render(<TodoList />)
    expect(await screen.findByText('todo_1')).toBeTruthy()
    expect(await screen.findByText('not_started')).toBeTruthy()
    expect(await screen.findByText('todo_2')).toBeTruthy()
    expect(await screen.findByText('in_progress')).toBeTruthy()
  })
  it('ローディング中ならその旨を表示する', async () => {
    vi.mocked(useTodoList).mockReturnValue({
      data: {
        data: [],
      },
      isLoading: true,
      isError: false,
      error: null,
    } as unknown as ReturnType<typeof useTodoList>)
    render(<TodoList />)
    expect(await screen.findByText('loading...')).toBeTruthy()
  })
  it('0件なら空状態を表示する', async () => {
    vi.mocked(useTodoList).mockReturnValue({
      data: {
        data: [],
      },
      isLoading: false,
      isError: false,
      error: null,
    } as unknown as ReturnType<typeof useTodoList>)
    render(<TodoList />)
    expect(await screen.findByText('Todo がありません')).toBeTruthy()
  })
  it('失敗ならエラーを表示する', async () => {
    vi.mocked(useTodoList).mockReturnValue({
      data: {
        data: [],
      },
      isLoading: false,
      isError: true,
      error: { message: 'HTTP error! status: 500' },
    } as unknown as ReturnType<typeof useTodoList>)
    render(<TodoList />)
    expect(await screen.findByText('HTTP error! status: 500')).toBeTruthy()
  })
})
