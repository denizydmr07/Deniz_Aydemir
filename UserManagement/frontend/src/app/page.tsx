"use client";

import UserList from './components/UserList'
import styles from './page.module.css'


/**
 * The home page
 * @returns The home page
*/
export default function Home() {

  // Return the JSX element
  return (
    <main className={styles.main} >
      <h1 >
        UserManagement
      </h1>
      <UserList/>
    </main>
  )
}
