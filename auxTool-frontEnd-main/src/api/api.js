import myAxios from './myAxios'

export const getMenuInfo = (pageKind) => {
  return myAxios({
    method: 'GET',
    url: '/getPageMenu',
    params: {
      pageKind: pageKind
    }
  })
}

export const addMenuItem = (data) => {
  return myAxios({
    method: 'POST',
    url: '/addPageMenuItem',
    data: data
  })
}

export const isLogin = (data) => {
  return myAxios({
    method: 'POST',
    url: '/login',
    data: data
  })
}

export const isRegister = (data) => {
  return myAxios({
    method: 'POST',
    url: '/register',
    data: data
  })
}

export const getCard = (data) => {
  return myAxios({
    method: 'GET',
    url: '/getCard',
    params:data
  })
}

export const sendMenu = (data) => {
  //这里是添加菜单的接口
  return myAxios({
    method: 'POST',
    url: '/sendmenu',
    data: data
  })
}

export const getCsvData = (data) => {
  return myAxios({
    method: 'GET',
    url: '/getCsv',
    params:data
  })
}