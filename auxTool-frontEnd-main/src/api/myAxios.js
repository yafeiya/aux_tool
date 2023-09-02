import axios from 'axios'
import { EndUrl } from '../../url_config'
const baseUrl  = EndUrl().backEndUrl

const myAxios = axios.create({
    baseURL: baseUrl,
    timeout: 200000 // 请求超时时间
})

export default myAxios