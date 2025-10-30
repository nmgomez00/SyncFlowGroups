import React from "react";

export default function UserList({ users, onDelete, onChangeRole, onLeaveGroup, roleEdit, setRoleEdit, groupId }) {

  return (
    <table className="min-w-full bg-white border rounded">
      <thead>
        <tr>
          <th className="px-4 py-2">Nombre</th>
          <th className="px-4 py-2">Email</th>
          <th className="px-4 py-2">Acciones</th>
        </tr>
      </thead>
      <tbody>
        {users.map((u) => (
          <tr key={u.id}>
            <td className="px-4 py-2">{u.name}</td>
            <td className="px-4 py-2">{u.email}</td>
            <td className="px-4 py-2 space-x-2">
              <button className="text-red-600" onClick={() => onDelete(u.id)}>Eliminar</button>
              <input
                type="text"
                placeholder="Role"
                value={roleEdit[u.id] || ""}
                onChange={e => setRoleEdit({ ...roleEdit, [u.id]: e.target.value })}
                className="border px-1 rounded w-20"
              />
              <button className="text-indigo-600" onClick={() => onChangeRole(u.id)}>Cambiar Rol</button>
              <button className="text-orange-600" onClick={() => onLeaveGroup(u.id)}>Eliminar del Grupo</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
