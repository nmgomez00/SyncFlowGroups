import React from 'react'
import GroupList from '../components/GroupList'
import GroupForm from '../components/GroupForm'
export default function GroupListPage() {
  return (
    <div className="max-w-4xl mx-auto py-6">
      <h2 className="text-2xl font-semibold mb-4">Lista de grupos</h2>
      <GroupList />
    </div>
  )
}
