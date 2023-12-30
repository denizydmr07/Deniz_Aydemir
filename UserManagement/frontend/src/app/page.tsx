"use client";

import Image from 'next/image'
import UserList from './components/UserList'
import styles from './page.module.css'

export default function Home() {
  return (
    <main className={styles.main} >
      <h1 >
        UserManagement
      </h1>
      <UserList/>
    </main>
  )
}
