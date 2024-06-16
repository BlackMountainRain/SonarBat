import { OAuthProvider } from '@/app/types/auth';

export async function fetchToken(
  provider: 'Self' | 'GitHub' | 'Google' | 'X' | 'Facebook' | 'Microsoft' | 'LinkedIn' | 'Apple' | 'Amazon',
  code: string = '',
  email: string = '',
  password: string = '',
): Promise<string> {
  const headers = new Headers({
    'Content-Type': 'application/json',
    Authorization: `Bearer ${localStorage.getItem('token')}`,
  });

  let url = '';
  let payload = '';
  if (provider === 'Self') {
    payload = JSON.stringify({ email, password });
    url = `${process.env.NEXT_PUBLIC_AUTH_HTTP_ADDR}/api/v1/auth/sign-in/password`;
  } else {
    payload = JSON.stringify({ provider, code });
    url = `${process.env.NEXT_PUBLIC_AUTH_HTTP_ADDR}/api/v1/auth/sign-in/oauth`;
  }
  const response = await fetch(
    url,
    {
      headers,
      method: 'POST',
      body: payload,
    },
  );
  if (response.status !== 200) {
    let msg = 'Get token failed.';
    try {
      const data = await response.json();
      if (data.message) {
        msg = data.message;
      }
    } catch (error) {
      throw new Error(msg);
    }
    throw new Error(msg);
  }
  const data = await response.json();
  return data.token;
}

export async function signUp(
  email: string,
  password: string,
): Promise<string> {
  const headers = new Headers({
    'Content-Type': 'application/json',
    Authorization: `Bearer ${localStorage.getItem('token')}`,
  });
  const response = await fetch(`${process.env.NEXT_PUBLIC_AUTH_HTTP_ADDR}/api/v1/auth/sign-up`, {
    headers,
    method: 'POST',
    body: JSON.stringify({ email, password }),
  });
  if (response.status !== 200) {
    let msg = 'Sign up failed.';
    try {
      const data = await response.json();
      if (data.message) {
        msg = data.message;
      }
    } catch (error) {
      throw new Error(msg);
    }
    throw new Error(msg);
  }
  const data = await response.json();
  return data.token;
}

export async function fetchOAuthProviders(): Promise<OAuthProvider[]> {
  return new Promise((resolve) => {
    resolve([
      {
        provider: 'GitHub',
        authorize_url: 'https://github.com/login/oauth/authorize',
        client_id: process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID || '',
        redirect_uri: `${process.env.NEXT_PUBLIC_BASE_URL || ''}/auth/callback`,
        scope: 'read:user',
        response_type: 'code',
      },
      {
        provider: 'Google',
        authorize_url: 'https://accounts.google.com/o/oauth2/auth',
        client_id: process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || '',
        redirect_uri: `${process.env.NEXT_PUBLIC_BASE_URL || ''}/auth/callback`,
        scope: 'openid profile email',
        response_type: 'code',
      },
    ]);
  });
}
