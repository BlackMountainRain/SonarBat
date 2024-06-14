import { OAuthProvider } from '@/app/types/auth';

export async function fetchToken(
  provider: 'self' | 'github' | 'google',
  code: string = '',
): Promise<string> {
  console.log('fetch token by:', provider, code);
  return new Promise((resolve) => {
    resolve('fake-token');
  });
}

export async function fetchOAuthProviders(): Promise<OAuthProvider[]> {
  return new Promise((resolve) => {
    resolve([
      {
        provider: 'github',
        authorize_url: 'https://github.com/login/oauth/authorize',
        client_id: process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID || '',
        redirect_uri: `${process.env.NEXT_PUBLIC_BASE_URL || ''}/auth/callback`,
        scope: 'read:user',
        response_type: 'code',
      },
      {
        provider: 'google',
        authorize_url: 'https://accounts.google.com/o/oauth2/auth',
        client_id: process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || '',
        redirect_uri: `${process.env.NEXT_PUBLIC_BASE_URL || ''}/auth/callback`,
        scope: 'openid profile email',
        response_type: 'code',
      },
    ]);
  });
}
