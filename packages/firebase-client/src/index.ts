import { initializeApp } from 'firebase/app'
import 'firebase/auth'
import 'firebase/storage'
import 'firebase/firestore'
import { getStorage } from 'firebase/storage'
import { getAuth } from 'firebase/auth'

const firebaseApp = initializeApp({
  apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY,
  authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN,
  projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID,
  storageBucket: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.NEXT_PUBLIC_FIREBASE_APP_ID
})

const auth = getAuth(firebaseApp)
auth.tenantId = process.env.NEXT_PUBLIC_FIREBASE_TENANT_ID || null
const storage = getStorage(firebaseApp)

class FirebaseAuthenticationError extends Error {
  constructor(error: any) {
    super(error.message)
    this.name = new.target.name
    Object.setPrototypeOf(this, new.target.prototype)

    // https://github.com/firebase/firebase-js-sdk/blob/master/packages/auth/src/error_auth.js
    switch (error.code) {
      case 'auth/email-already-in-use':
        this.message = '入力されたメールアドレスはすでに使用されています。'
        break
      case 'auth/invalid-email':
        this.message = '不正なメールアドレスです'
        break
      case 'auth/user-not-found':
        this.message = 'ユーザーが見つかりませんでした'
        break
      case 'auth/wrong-password':
        this.message = 'パスワードが一致しません'
        break
      default:
        this.message = 'エラーが発生しました'
        break
    }
  }
}

export { firebaseApp, auth, storage, FirebaseAuthenticationError }
