import * as jose from 'jose';

const isTokenExpired = (token: string): boolean => {
  if (!token) {
    return true;
  }

  try {
    const claims = jose.decodeJwt(token);
    const exp = claims.exp ?? 0;
    if (exp === 0 || exp <= Date.now() / 1000) {
      return true;
    }
  } catch (error) {
    return true;
  }
  return false;
};

export { isTokenExpired };
