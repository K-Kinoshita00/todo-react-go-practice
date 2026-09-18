import {
  CognitoUserPool,
  CognitoUser,
  AuthenticationDetails,
} from 'amazon-cognito-identity-js'

export const signIn = (username: string, password: string): Promise<string> => {
  const userPoolId: string = import.meta.env.VITE_COGNITO_USER_POOL_ID ?? ''
  const clientId: string = import.meta.env.VITE_COGNITO_CLIENT_ID ?? ''
  if (userPoolId === '')
    throw new Error('signIn: VITE_COGNITO_USER_POOL_ID is missing')
  if (clientId === '')
    throw new Error('signIn: VITE_COGNITO_CLIENT_ID is missing')

  // 接続先の情報をまとめる、このアプリ用のユーザー台帳を指す
  const userPool = new CognitoUserPool({
    UserPoolId: userPoolId,
    ClientId: clientId,
  })
  // 台帳内のどのユーザーかを指す
  const cognitoUser = new CognitoUser({
    Username: username,
    Pool: userPool,
  })
  // そのユーザーである証明情報
  const details = new AuthenticationDetails({
    Username: username,
    Password: password,
  })
  // CognitoUser が details を Cognito に渡して検証
  return new Promise((resolve, reject) => {
    cognitoUser.authenticateUser(details, {
      onSuccess: (session) => resolve(session.getIdToken().getJwtToken()),
      onFailure: (err) => reject(err),
    })
  })
}
