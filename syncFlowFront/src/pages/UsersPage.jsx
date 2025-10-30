import { useEffect, useState } from "react";
import { getUsers, createUser, deleteUser } from "../api";
export default function UsersPage() {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [form, setForm] = useState({ name: "", email: "", profilePhotoURL: "" });

  useEffect(() => {
    async function load() {
      try {
        setLoading(true);
        const data = await getUsers();
        setUsers(data || []);
      } catch (err) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, []);

  async function handleDelete(userId) {
    if (!window.confirm("Eliminar usuario?")) return;
    await deleteUser(userId);
    setUsers((prev) => prev.filter((u) => u.id !== userId));
  }

  async function handleCreate(e) {
    e.preventDefault();
    await createUser(form);
    setForm({ name: "", email: "", profilePhotoURL: "" });
    const data = await getUsers();
    setUsers(data || []);
  }

  return (
    <div className="max-w-2xl mx-auto py-6">
      <h2 className="text-2xl font-bold mb-4">Usuarios</h2>
      <form onSubmit={handleCreate} className="mb-4 flex gap-2">
        <input
          type="text"
          placeholder="Nombre"
          value={form.name}
          onChange={e => setForm((f) => ({ ...f, name: e.target.value }))}
          className="border px-2 py-1 rounded"
          required
        />
        <input
          type="email"
          placeholder="Email"
          value={form.email}
          onChange={e => setForm((f) => ({ ...f, email: e.target.value }))}
          className="border px-2 py-1 rounded"
          required
        />
        <button className="bg-green-500 text-white px-3 py-1 rounded">Crear usuario</button>
      </form>
      {loading ? (
        <div>Cargando usuarios...</div>
      ) : error ? (
        <div className="text-red-600">{error}</div>
      ) : (
        <ul>
          {users.map((u) => (
            <li key={u.id} className="flex items-center justify-between py-1">
              <span>{u.name} ({u.email})</span>
              <button
                className="text-xs px-2 py-1 bg-red-100 rounded"
                onClick={() => handleDelete(u.id)}
              >
                Eliminar
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
