import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import QueryProvider from './lib/react-query/QueryProvider.tsx'
import { BrowserRouter } from 'react-router'
import Template from './layouts/Template.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryProvider>
      <BrowserRouter>
        <Template>
          <App />
        </Template>
      </BrowserRouter>
    </QueryProvider>
  </StrictMode>,
)
