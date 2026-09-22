import { test, expect } from '@playwright/test'

test.describe('Todo CRUD', () => {
  test('login then crud', async ({ page }) => {
    const email = process.env.E2E_EMAIL
    const password = process.env.E2E_PASSWORD
    if (email == null || password == null) {
      throw new Error('E2E_EMAIL / E2E_PASSWORD is not set')
    }

    await test.step('login', async () => {
      await page.goto('/login')
      await page.getByPlaceholder('email').fill(email)
      await page.getByPlaceholder('password').fill(password)
      await page.getByRole('button', { name: 'ログイン' }).click()
      await expect(page).toHaveURL('/')
    })

    const inputTitle = `e2e-${Date.now()}`
    await test.step('create', async () => {
      await page.getByRole('button', { name: 'TODO 作成' }).click()
      await page.getByLabel('title').fill(inputTitle)
      await page.getByRole('combobox').click()
      await page.getByRole('option', { name: 'in_progress' }).click()
      await page.getByRole('button', { name: '作成' }).click()
      const card = page.locator('.MuiCard-root').filter({ hasText: inputTitle })
      await expect(card.getByText(inputTitle)).toBeVisible()
      await expect(card.getByText('in_progress')).toBeVisible()
    })

    const card = page.locator('.MuiCard-root').filter({ hasText: inputTitle })
    const updateTitle = `e2e-update-${Date.now()}`
    await test.step('update', async () => {
      await card.getByRole('button', { name: '編集' }).click()
      await page.getByLabel('title').fill(updateTitle)
      await page.getByRole('combobox').click()
      await page.getByRole('option', { name: 'completed' }).click()
      await page.getByRole('button', { name: '更新' }).click()
      const updatedCard = page
        .locator('.MuiCard-root')
        .filter({ hasText: updateTitle })
      await expect(updatedCard.getByText(updateTitle)).toBeVisible()
      await expect(updatedCard.getByText('completed')).toBeVisible()
    })

    const updatedCard = page
      .locator('.MuiCard-root')
      .filter({ hasText: updateTitle })
    await test.step('delete', async () => {
      await updatedCard.getByRole('button', { name: '削除' }).click()
      await page.getByRole('button', { name: '削除' }).click()
      await expect(page.getByText(updateTitle)).toHaveCount(0)
    })
  })
})
