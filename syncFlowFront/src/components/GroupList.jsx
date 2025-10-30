import { useEffect, useState } from 'react'
import { Avatar, List, ListItem } from "flowbite-react";
import { Link } from 'wouter'
import { getGroups, deleteGroup } from '../api'

export function GroupList() {
  const [groups, setGroups] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

    const load = async () => {
      setLoading(true)
      try {
        const data = await getGroups()
        setGroups(data || [])
      } catch (err) {
        setError(err?.response?.data || err.message)
      } finally {
        setLoading(false)
      }
    }

  useEffect(() => { load() }, [])

  async function handleDelete(id) {
    if (!confirm('Desea eliminar el grupo?')) return
    try {
      await deleteGroup(id)
      setGroups((s) => s.filter((g) => g.id !== id))
    } catch (err) {
      alert('Fallo al eliminar: ' + err.message)
    }
  }


  if (loading) return <div>Cargando grupos...</div>
  if (error) return <div className="text-red-600">{error}</div>

  return (
    <>
      {groups.length === 0 && <div className="text-sm text-slate-500">Aun no hay grupos.</div>}
      <ul className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {groups.map((g) => (
          <li key={g.id} className="bg-white p-4 rounded-lg shadow-sm border flex flex-col">
            <div className="flex items-start justify-between">
              <div>
                <h3 className="text-lg font-medium text-slate-900">{g.name}</h3>
                <p className="text-sm text-slate-500 mt-1">{g.description}</p>
              </div>
              <div className="flex items-center gap-2">
                <Link href={`/groups/${g.id}`} className="text-indigo-600 hover:underline">Ver</Link>
                <button onClick={() => handleDelete(g.id)} className="text-sm text-red-600">Eliminar</button>
              </div>
            </div>
            <div className="mt-3 text-xs text-slate-400">Privacidad: {g.privacy} • Estado: {g.state}</div>
          </li>
        ))}
      </ul>
    </>
  )
}

  export default GroupList
