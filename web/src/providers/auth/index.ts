const STORAGE_KEY = 'idToken'

export const getToken = (): string | null => {
  const token = sessionStorage.getItem(STORAGE_KEY)
  return token
}

export const setToken = (token: string): void => {
  sessionStorage.setItem(STORAGE_KEY, token)
}

export const clearToken = (): void => {
  sessionStorage.removeItem(STORAGE_KEY)
}
