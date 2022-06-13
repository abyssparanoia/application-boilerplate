import React, { useEffect } from 'react'
import { useRouter } from 'next/router'
import { useCallback, useState } from 'react'
import { auth } from '@boilerplate/firebase-client'
import { GoogleAuthProvider, signInWithEmailAndPassword, signInWithPopup } from 'firebase/auth'

interface ISignInForm {
  email: string
  password: string
}

const handleError = (error: unknown, setError?: React.Dispatch<React.SetStateAction<Error | undefined>>) => {
  let errorMessage = 'Auth Error occurred'
  if (error instanceof Error) {
    setError?.(error)
    errorMessage += ` :${error.message}`
    errorMessage
  }
  //   Toaster.error(errorMessage)
}

export const useSignIn = () => {
  const router = useRouter()

  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<Error>()

  const handleSignin = useCallback(
    async ({ email, password }: ISignInForm) => {
      try {
        console.log(email, password)
        setIsLoading(true)
        setError(undefined)
        const { user } = await signInWithEmailAndPassword(auth, email, password)

        // const { claims } = (await auth.currentUser?.getIdTokenResult()) || {}
        // if (claims?.role !== 'ADMIN') {
        //   throw new Error('ログインする権限がありません')
        // }

        console.log(user)

        if (user) {
          router.push('/')
        } else {
          throw new Error('ユーザーが存在しません')
        }
      } catch (error: any) {
        console.log(error)
        handleError(error, setError)
      } finally {
        setIsLoading(false)
      }
    },
    [router]
  )

  const handleSignInWithGoogle = useCallback(async () => {
    try {
      setIsLoading(true)
      setError(undefined)
      const { user } = await signInWithPopup(auth, new GoogleAuthProvider())

      // const { claims } = (await auth.currentUser?.getIdTokenResult()) || {}
      // if (claims?.role !== 'ADMIN') {
      //   throw new Error('ログインする権限がありません')
      // }

      console.log(user)

      if (user) {
        router.push('/')
      } else {
        throw new Error('ユーザーが存在しません')
      }
    } catch (error: any) {
      console.log(error)
      handleError(error, setError)
    } finally {
      setIsLoading(false)
    }
  }, [router])

  const clearError = useCallback(() => {
    setError(undefined)
  }, [])

  return { isLoading, error, handleSignin, clearError, handleSignInWithGoogle }
}

export const useSignOut = () => {
  const router = useRouter()

  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<Error>()

  const handleSignOut = useCallback(async () => {
    try {
      setIsLoading(true)
      await auth.signOut()
      setIsLoading(false)
      router.push('/sign_in')
    } catch (error: any) {
      setError(error)
    } finally {
      setIsLoading(false)
    }
  }, [router])
  return { isLoading, error, handleSignOut }
}

export const useAuthRequired = () => {
  const router = useRouter()
  const [isInit, setIsInit] = useState(false)

  useEffect(() => {
    // theme の override 時に使用する
    const jssStyles = document.querySelector('#jss-server-side')
    if (jssStyles) {
      jssStyles.parentElement?.removeChild(jssStyles)
    }
    if (!isInit) {
      const unsubscriber = auth.onIdTokenChanged(user => {
        if (!user) {
          router.push('/sign_in')
        }
        setIsInit(true)
        unsubscriber()
      })
    }
  }, [isInit, router])

  return { isInit }
}

export const useIsLoggedIn = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const unsubscriber = auth.onIdTokenChanged(user => {
      setIsLoggedIn(!!user)
    })
    return () => {
      unsubscriber()
    }
  }, [])

  return { isLoggedIn }
}
