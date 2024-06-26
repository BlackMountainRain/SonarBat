interface Claims {
  iss?: string; // Issuer
  sub?: string; // Subject
  aud?: string; // Audience
  exp?: number; // Expiration Time
  nbf?: number; // Not Before
  iat?: number; // Issued At
  jti?: string; // JWT ID

  uid: string;
  username: string;
  email: string;
  status: boolean;
  avatar_url: string;
}

interface User {
  uid: string;
  username: string;
  email: string;
  status: boolean;
  avatar_url: string;
}

export { type Claims, type User };
