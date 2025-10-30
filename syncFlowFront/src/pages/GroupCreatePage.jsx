import React from 'react'
import GroupList from '../components/GroupList'
import GroupForm from '../components/GroupForm'
export default function GroupCreatePage() {
  return (
    <div className="max-w-4xl mx-auto py-6">
      <h2 className="text-2xl font-semibold mb-4">Crear grupo</h2>
      <GroupForm/>
    </div>
  )
}
