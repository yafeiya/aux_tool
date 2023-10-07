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

export const getDesignCards = (data) => {
  return myAxios({
    method: 'GET',
    url: '/getDesignCards',
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

export const updataCard = (data) => {
  return myAxios({
    method: 'GET',
    url: '/updataCard',
    params:data
  })
}

export const updateDesignCard = (data) => {
  return myAxios({
    method: 'GET',
    url: '/updateDesignCard',
    params:data
  })
}

export const createCard = (data) => {
  return myAxios({
    method: 'GET',
    url: '/createCard',
    params:data
  })
}

export const deleteCard = (data) => {
  return myAxios({
    method: 'GET',
    url: '/deleteCard',
    params:data
  })
}
export const deleteDesignCard = (data) => {
  return myAxios({
    method: 'GET',
    url: '/deleteDesignCard',
    params:data
  })
}

export const addPageMenuItem = (data) => {
  return myAxios({
    method: 'GET',
    url: '/addPageMenuItem',
    params:data
  })
}
export const deletePageMenuItem = (data) => {
  return myAxios({
    method: 'GET',
    url: '/deletePageMenuItem',
    params:data
  })
}
export const downloadCsvFile = (data) => {
  //这里是下载的接口
  return myAxios({
    method: "POST",
    url: "/downloadCsvFile",
    data: data
  })
}

export const addDesignCard = (data) => {
  return myAxios({
    method: "GET",
    url: "/addDesignCard",
    params: data
  })
}

