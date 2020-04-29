const isLoggedIn = (user) => {
  return user.roles && user.roles.length > 0
}

const authorize = (state, permission) => {
  let { roles, permissions } = state.rbac
  if (roles && permissions && state.roles && state.roles.length) {
    for (let i = 0; i < state.roles.length; i++) {
      let role = state.roles[i]
      if ((role === 'admin') || (roles[role] & permissions[permission])) {
        return true
      }
    }
  }
  return false
}

const hasRole = (user, roleName?) => {
  return user.roles && user.roles.indexOf(roleName) !== -1
}

const isAdmin = (user) => {
  return hasRole(user, 'admin')
}

export default {
  access: () => {
    return {}
  },
  uid: (state) => {
    return state._id
  },
  authorize: (state) => (accessLevel) => {
    return authorize(state, accessLevel)
  },
  isLoggedIn: (state) => {
    return isLoggedIn(state)
  },
  isAdmin: (state) => {
    return isAdmin(state)
  },
  isCheck: (state) => {
    return !!state.isCheck
  }
}
