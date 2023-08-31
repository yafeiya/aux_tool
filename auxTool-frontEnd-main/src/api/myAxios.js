import axios from 'axios'
const baseUrl = 'http://192.168.0.103:2345';
const myAxios = axios.create({
    baseURL: baseUrl,
    timeout: 200000 // 请求超时时间
})

export default myAxios