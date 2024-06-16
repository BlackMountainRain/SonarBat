export interface OAuthProvider {
  provider: 'Self' | 'GitHub' | 'Google' | 'X' | 'Facebook' | 'Microsoft' | 'LinkedIn' | 'Apple' | 'Amazon';
  authorize_url: string;
  client_id: string;
  redirect_uri: string;
  scope: string;
  response_type: 'code';
}
