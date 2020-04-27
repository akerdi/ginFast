import axios from './http'

export const userProfile = () => { return axios.get(`/profile`) }
export const logout = () => { return axios.post('/logout') }
