export function EndUrl(){
    // 检查是否在浏览器环境中
    const isBrowser = typeof window !== 'undefined'
    
    const url_dict = {
        // 向后端发送请求的地址+端口 - 浏览器和Node.js环境都使用完整URL
        "backEndUrl": 'http://127.0.0.1:8080',
        
        // 局域网内暴露本机地址 - 浏览器和Node.js环境都使用完整URL
        "fileUrl": 'http://127.0.0.1:3002'
    }
    
    // 解析前端URL和端口 - 统一使用配置的URL
    var tmpurl = url_dict.fileUrl.split('//')[1].split(':')
    url_dict["frontEndUrl"] = tmpurl[0]
    url_dict["frontpost"] = tmpurl[1]
    
    return url_dict
}