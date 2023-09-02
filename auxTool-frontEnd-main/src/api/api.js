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
export const isLogin = (username,password) => {
  return myAxios({
    method: 'GET',
    url: '/login',
    params: {
      username: username,
      password: password
    }
  })
}
export const postRegist = (data) => {
  return myAxios({
    method: 'POST',
    url: '/postRegist',
    data: data
  })
}

