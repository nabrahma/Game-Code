import ky from 'ky';

export const api = ky.extend({
  prefix: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
  credentials: 'include',
  hooks: {
    afterResponse: [
      // @ts-ignore
      async ({ request, options, response }: any) => {
        if (response.status === 401 && !request.url.includes('/auth/refresh')) {
          // Attempt to refresh the token transparently
          try {
            await ky.post(`${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'}/auth/refresh`, {
              credentials: 'include',
            });
            // Retry the original request
            return ky(request);
          } catch (error) {
            // Refresh failed, redirect to login
            if (typeof window !== 'undefined') {
              window.location.href = '/login';
            }
          }
        }
        return response;
      }
    ]
  }
});
