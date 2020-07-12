import React from 'react'
import { ExNextPageContext } from 'next'
import Router from 'next/router'
import { authorize } from '@boilerplate/fixtures'
import { useAuthCookie, useSignIn, useSignOut } from '@boilerplate/fixtures'

type InitialProps = {}
type Props = {} & InitialProps

const Index = (_: Props) => {
  const { idToken } = useAuthCookie()
  const { handleSignInWithGoogle } = useSignIn()
  const { handleSignOut } = useSignOut()

  return (
    <div>
      <button onClick={() => Router.push('/about')}>please click here!</button>
      {!idToken && <button onClick={() => handleSignInWithGoogle()}>sign in</button>}
      {idToken && <button onClick={() => handleSignOut()}>sign out</button>}
    </div>
  )
}

export const getServerSideProps = async (ctx: ExNextPageContext) => {
  await authorize(ctx)
  return { props: {} }
}

export default Index
