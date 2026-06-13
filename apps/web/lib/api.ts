import ky from 'ky';

export const api = ky.create({
  prefixUrl: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
  hooks: {
    beforeRequest: [
      (request) => {
        // Here we could add access tokens from an in-memory store if needed.
        // However, we are using httpOnly cookies as per PRD for secure auth.
        // So we just ensure credentials are included.
        request.credentials = 'include';
      }
    ],
    afterResponse: [
      async (request, options, response) => {
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
