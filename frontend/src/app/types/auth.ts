export interface OAuthProvider {
  provider: 'self' | 'github' | 'google';
  authorize_url: string;
  client_id: string;
  redirect_uri: string;
  scope: string;
  response_type: 'code' | 'token';
}
