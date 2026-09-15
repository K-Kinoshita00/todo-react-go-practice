const STORAGE_KEY = "dev.jwt"

export const getToken = (): string | null => {
  const token = localStorage.getItem(STORAGE_KEY)
  return token
}

export const setToken = (token: string): void => {
  localStorage.setItem(STORAGE_KEY, token)
}