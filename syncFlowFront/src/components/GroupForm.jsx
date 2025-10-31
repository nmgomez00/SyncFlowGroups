import React, { useState } from 'react'
import { createGroup, ensureLocalUser } from '../api'

export default function GroupForm({ onCreated }) {
  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(e) {
    e.preventDefault()
    setLoading(true)
    try {
      const userID = await ensureLocalUser()
      const res = await createGroup({ name, description, userID })
      setName('')
      setDescription('')
      if (onCreated) onCreated(res)
    } catch (err) {
      alert('Fallo al crear el grupo: ' + (err?.response?.data || err.message))
    } finally {
      setLoading(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-white p-4 rounded-lg shadow-sm border">
      <h4 className="text-sm font-medium mb-3">Nuevo Grupo</h4>
      <div className="space-y-2">
        <input value={name} onChange={(e) => setName(e.target.value)} required placeholder="Nombre" className="w-full rounded-md border px-3 py-2" />
        <textarea value={description} onChange={(e) => setDescription(e.target.value)} placeholder="Descripcion" className="w-full rounded-md border px-3 py-2" />
        <div className="flex items-center justify-end">
          <button disabled={loading} className="bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700">{loading ? 'Creando' : 'Crear'}</button>
        </div>
      </div>
    </form>
  )
}
