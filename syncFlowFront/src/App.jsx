import React from 'react'
import { Route, Switch, Link } from 'wouter'
import GroupsCreatePage from './pages/GroupCreatePage'
import GroupPage from './pages/GroupPage'
import GroupsListPage from './pages/GroupListPage'
import UsersPage from './pages/UsersPage'

export default function App() {
  return (
    <div className="min-h-screen bg-slate-50 text-slate-800">
      <header className="bg-white border-b">
        <div className="max-w-6xl mx-auto px-4 py-4 flex items-center justify-between">
          <Link href="/groups" className="text-xl font-semibold">SyncFlowGroups</Link>
          <nav className="space-x-4">
            <Link href="/groups" className="text-slate-600 hover:text-slate-900">Grupos</Link>
            <Link href="/create/group" className="text-slate-600 hover:text-slate-900">Crear grupo</Link>
            <Link href="/users" className="text-slate-600 hover:text-slate-900">Usuarios</Link>
          </nav>
        </div>
      </header>

      <main className="max-w-6xl mx-auto px-4 py-8">
        <Switch>
          <Route path="/" component={GroupsCreatePage} />
          <Route path="/create/group" component={GroupsCreatePage} />
          <Route path="/groups" component={GroupsListPage} />
          <Route path="/groups/:id" component={GroupPage} />
          <Route path="/users" component={UsersPage} />
        </Switch>
      </main>
    </div>
  )
}
