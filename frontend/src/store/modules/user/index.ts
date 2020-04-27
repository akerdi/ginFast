import _ from 'lodash'
import Vue from 'vue'
import { userProfile } from '@/service/user'
import { rbacInfo } from '@/service/rbac'
import { USER_SIGNIN, USER_SIGNOUT, USER_CHECK, USER_INFO, RBAC_INFO } from '../../mutation-types'
import getters from './getters'

const state = {
  rbac: {},
  username: '',
  email: '',
  roles: [],
  avatar: '',
  isCheck: false
}

const INFO_ARR = ['username', 'email', 'roles', 'avatar']

export default {
  namespaced: true,
  state,
  getters,
  mutations: {
    [USER_CHECK] (state, data) {
      state.isCheck = data
    },
    [USER_SIGNIN] (state, data) {
      let user = _.pick(data, INFO_ARR)
      Object.keys(user).forEach(k => { Vue.set(state, k, user[k]) })
    },
    [USER_SIGNOUT] (state) {
      INFO_ARR.forEach(k => { Vue.delete(state, k) })
    },
    [RBAC_INFO] (state, data) {
      Vue.set(state, 'rbac', data)
    }
  },
  actions: {
    [USER_SIGNIN] ({ commit }, user) {
      commit(USER_SIGNIN, user)
    },
    [USER_SIGNOUT] ({ commit }) {
      commit(USER_SIGNOUT)
    },
    [USER_INFO] ({ commit }) {
      return new Promise((resolve, reject) => {
        userProfile().then((response) => {
          if (!response.data) {
            return reject(new Error('未获取用户信息'))
          }
          let data = JSON.parse(JSON.stringify(response.data))
          data.roles = _.map(data.roles, 'name')
          commit(USER_SIGNIN, data)
          commit(USER_CHECK, true)
          resolve()
        }).catch((error: Error) => {
          commit(USER_CHECK, true)
          reject(error)
        })
      })
    },
    [RBAC_INFO] ({ commit }) {
      return new Promise((resolve, reject) => {
        rbacInfo().then(response => {
          let data = JSON.parse(JSON.stringify(response.data))
          commit(RBAC_INFO, data)
          resolve()
        }).catch((error: Error) => {
          reject(error)
        })
      })
    }
  }
}
