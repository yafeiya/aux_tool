import axios from 'axios'
//import { backEndUrl } from '../../url'
const baseUrl = 'http://192.168.0.103:8080';
//const baseUrl = backEndUrl;
const myAxios = axios.create({
    baseURL: baseUrl,
    timeout: 200000 // 请求超时时间
})

export default myAxios