export async function fetchToken(
  provider:
    | 'SELF'
    | 'GITHUB'
    | 'GOOGLE'
    | 'X'
    | 'FACEBOOK'
    | 'MICROSOFT'
    | 'LINKEDIN'
    | 'APPLE'
    | 'AMAZON',
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
  if (provider === 'SELF') {
    payload = JSON.stringify({ email, password });
    url = `${process.env.NEXT_PUBLIC_AUTH_HTTP_ADDR}/api/v1/auth/sign-in/password`;
  } else {
    payload = JSON.stringify({ provider, code });
    url = `${process.env.NEXT_PUBLIC_AUTH_HTTP_ADDR}/api/v1/auth/sign-in/oauth`;
  }
  const response = await fetch(url, {
    headers,
    method: 'POST',
    body: payload,
  });
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

export async function signUp(email: string, password: string): Promise<string> {
  const headers = new Headers({
    'Content-Type': 'application/json',
    Authorization: `Bearer ${localStorage.getItem('token')}`,
  });
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_AUTH_HTTP_ADDR}/api/v1/auth/sign-up`,
    {
      headers,
      method: 'POST',
      body: JSON.stringify({ email, password }),
    },
  );
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
