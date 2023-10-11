export function EndUrl(){
  const url_dict =
  {
    // 局域网内暴露本机地址
    "frontEndUrl":'127.0.0.1',
    "frontpost":"5173",
    // 向后端发送请求的地址+端口
    "backEndUrl": 'http://127.0.0.1:8080',
    "fileUrl": 'http://127.0.0.1:5173'
}
    return url_dict

}
