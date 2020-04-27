import axios from './http'

export const rbacInfo = () => { return axios.get('/rbac') }
