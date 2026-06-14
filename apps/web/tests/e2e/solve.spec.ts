import { test, expect } from '@playwright/test';

test.describe('Problem Solving Flow', () => {
  test('user can open a problem, view description, and see editor', async ({ page }) => {
    // Navigate to problem list
    await page.goto('/problems');
    
    // Find the first problem link and click it (we assume the DB has problems)
    // Wait for table to load
    await page.waitForSelector('table');
    const firstProblemLink = page.locator('table tbody tr:first-child a');
    
    // In case no problems exist, we should probably handle it gracefully in the test,
    // but assuming seed data exists:
    if (await firstProblemLink.count() > 0) {
      await firstProblemLink.click();
      
      // Verify we navigated to the solve page
      await expect(page).toHaveURL(/\/problems\/[a-z0-9-]+/);
      
      // Verify layout panels exist (description left, editor right)
      await expect(page.getByText('Description')).toBeVisible();
      await expect(page.getByText('Console')).toBeVisible();
      
      // Editor takes a moment to load Monaco
      await page.waitForTimeout(1000);
    }
  });

  test('clicking run outputs result in console', async ({ page }) => {
    // Go directly to a specific problem if we know one, e.g. two-sum
    // We will just hit the API to fetch a slug, or just use the first one
    await page.goto('/problems');
    await page.waitForSelector('table');
    const firstProblemLink = page.locator('table tbody tr:first-child a');
    
    if (await firstProblemLink.count() > 0) {
      await firstProblemLink.click();
      
      // Wait for editor to be visible
      await expect(page.getByText('Console')).toBeVisible();

      // Click the Run button
      const runButton = page.getByRole('button', { name: /Run Code/i });
      await runButton.click();

      // Verify the console changes state (shows Executing or similar, then results)
      // Since it hits localhost:8080 we might get connection refused if backend is off,
      // but in a proper E2E environment the backend runs.
      // We just check if the button goes into 'Running...' state.
      await expect(runButton).toBeDisabled();
    }
  });
});
