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