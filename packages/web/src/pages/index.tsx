import React from 'react'
import Router from 'next/router'
import { useSignIn, useSignOut, useIsLoggedIn } from 'src/modules/auth/hooks'

type InitialProps = {}
type Props = {} & InitialProps

const Index = (_: Props) => {
  const { handleSignInWithGoogle } = useSignIn()
  const { handleSignOut } = useSignOut()
  const { isLoggedIn } = useIsLoggedIn()

  return (
    <div>
      <button onClick={() => Router.push('/about')}>please click here!</button>
      {!isLoggedIn && <button onClick={() => handleSignInWithGoogle()}>sign in</button>}
      {isLoggedIn && <button onClick={() => handleSignOut()}>sign out</button>}
    </div>
  )
}

export default Index
