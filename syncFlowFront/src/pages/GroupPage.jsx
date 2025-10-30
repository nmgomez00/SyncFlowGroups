import React, { useEffect, useState, useCallback } from 'react'
import { useRoute } from 'wouter'
import { getCategories, getChannelsByCategory, createCategory, createChannel, deleteCategory, deleteChannel, joinGroup, ensureLocalUser, getAllUsersByGroup, getUsers, changeUserRole, leaveGroup } from '../api'


export default function GroupPage() {
  const [match, params] = useRoute('/groups/:id')
  const groupID = params?.id
  const [categories, setCategories] = useState([])
  const [channelsByCategory, setChannelsByCategory] = useState({})
  const [members, setMembers] = useState([])
  const [allUsers, setAllUsers] = useState([])
  const [loading, setLoading] = useState(true)

  const loadAll = useCallback(async () => {
    if (!groupID) return
    setLoading(true)
    try {
      const cats = await getCategories(groupID)
      setCategories(cats || [])
      if (cats === null) {
        setChannelsByCategory({});
      }else{
        const chObj = {}
        for (const cat of cats) {
          chObj[cat.id] = await getChannelsByCategory(groupID, cat.id)
        }
        setChannelsByCategory(chObj)
      }
      setMembers(await getAllUsersByGroup(groupID))
      setAllUsers(await getUsers())
    } catch (err) {
      console.error(err)
    } finally {
      setLoading(false)
    }
  }, [groupID])

  useEffect(() => { loadAll() }, [groupID, loadAll])

  async function handleCreateCategory(e) {
    e.preventDefault()
    const form = e.target
    const name = form.name.value
    const description = form.description.value
    const userCreatedID = await ensureLocalUser()
    await createCategory(groupID, { name, description, userCreatedID })
    form.reset()
    await loadAll()
  }

  async function handleCreateChannel(categoryID, { name }) {
    const userID = await ensureLocalUser()
    await createChannel(groupID, categoryID, { name, description: '', userID })
    await loadAll()
  }

  async function handleDeleteCategory(id) {
    if (!confirm('Eliminar categoría?')) return
    await deleteCategory(groupID, id)
    await loadAll()
  }

  async function handleDeleteChannel(categoryID, channelID) {
    if (!confirm('Eliminar canal?')) return
    await deleteChannel(groupID, categoryID, channelID)
    await loadAll()
  }

  async function handleRemoveMember(userID) {
    await leaveGroup(groupID, userID)
    await loadAll()
  }

  async function handleChangeRole(userID, role) {
    await changeUserRole(groupID, userID, role)
    await loadAll()
  }

  async function handleAddMember(userID) {
    await joinGroup(groupID, { userID })
    await loadAll()
  }

  let nonMembers = [];
  if (!members || members.length === 0) {
    nonMembers = allUsers;
  } else {
    nonMembers = allUsers.filter(u =>
      !members.some(m => (m.userID || m.id) === (u.userID || u.id))
    );
  }


  if (!groupID) return <div>Selecciona un grupo</div>
  if (loading) return <div>Cargando...</div>

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-semibold">Grupo</h2>
      </div>

      <div className="mb-8">
        <h3 className="font-semibold mb-2">Miembros</h3>
        <ul className="mb-2">
          {members !== null &&members.map((m) => (
            <li key={m.id || m.userID} className="flex items-center justify-between py-1">
              <span>
                {m.name} ({m.role})
              </span>
              <div>
                <button
                  className="text-xs px-2 py-1 bg-red-100 rounded mr-2"
                  onClick={() => handleRemoveMember(m.id || m.userID)}
                >
                  Eliminar
                </button>
                <select
                  value={m.role}
                  onChange={e => handleChangeRole(m.id || m.userID, e.target.value)}
                  className="text-xs px-2 py-1 border rounded"
                >
                  <option value="USER">USER</option>
                  <option value="ADMIN">ADMIN</option>
                </select>
              </div>
            </li>
          ))}
        </ul>
        <h4 className="font-semibold mt-4 mb-2">Agregar Miembros</h4>
        <ul>
          {nonMembers !== null && nonMembers.map((u) => (
            <li key={u.id || u.userID} className="flex items-center justify-between py-1">
              <span>{u.name}</span>
              <button
                className="text-xs px-2 py-1 bg-blue-100 rounded"
                onClick={() => handleAddMember(u.id || u.userID)}
              >
                Agregar al grupo
              </button>
            </li>
          ))}
        </ul>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <h3 className="text-lg font-medium mb-3">Categorías & Canales</h3>
          <form onSubmit={handleCreateCategory} className="mb-4 flex gap-2">
            <input
              name="name"
              required
              placeholder="Nombre de categoría"
              className="border px-2 py-1 rounded"
            />
            <input
              name="description"
              placeholder="Descripcion"
              className="border px-2 py-1 rounded"
            />
            <button className="bg-green-500 text-white px-3 py-1 rounded">Agregar Categoría</button>
          </form>
          {categories.map((cat) => (
            <div key={cat.id} className="mb-6 border rounded p-3 bg-gray-50">
              <div className="flex justify-between items-center mb-2">
                <span className="font-semibold">{cat.name}</span>
                <button
                  className="text-xs px-2 py-1 bg-red-100 rounded"
                  onClick={() => handleDeleteCategory(cat.id)}
                >
                  Eliminar Categoría
                </button>
              </div>
              <div className="ml-4">
                <h4 className="font-semibold mb-1">Canales</h4>
                <ul>
                  {(channelsByCategory[cat.id] || []).map((ch) => (
                    <li key={ch.id} className="flex items-center justify-between py-1">
                      <span>{ch.name}</span>
                      <button
                        className="text-xs px-2 py-1 bg-red-100 rounded"
                        onClick={() => handleDeleteChannel(cat.id, ch.id)}
                      >
                        Eliminar
                      </button>
                    </li>
                  ))}
                </ul>
                <form
                  onSubmit={e => {
                    e.preventDefault();
                    const name = e.target.name.value;
                    handleCreateChannel(cat.id, { name });
                    e.target.reset();
                  }}
                  className="flex gap-2 mt-2"
                >
                  <input
                    name="name"
                    type="text"
                    placeholder="Nombre del canal"
                    className="border px-2 py-1 rounded"
                    required
                  />
                  <button className="bg-blue-500 text-white px-3 py-1 rounded">
                    Agregar Canal
                  </button>
                </form>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}
