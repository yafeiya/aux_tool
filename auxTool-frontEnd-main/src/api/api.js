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
export const deleteTable = (data) => {
  return myAxios({
    method: 'GET',
    url: '/deleteTable',
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
export const downloadModelFile = (data) => {
  //这里是下载的接口
  return myAxios({
    method: "POST",
    url: "/downloadModelFile",
    data: data
  })
}
export const getModelImage = (data) => {
  //这里是下载的接口
  return myAxios({
    method: "POST",
    url: "/getModelImage",
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

export const addDefinefunctionCard = (data) => {
  return myAxios({
    method: "GET",
    url: "/addDefinefunctionCard",
    params: data
  })
}
export const getDefinefunctionCards = (data) => {
  return myAxios({
    method: "GET",
    url: "/getDefinefunctionCards",
    params: data
  })
}
export const updateDefinefunctionCard = (data) => {
  return myAxios({
    method: "GET",
    url: "/updateDefinefunctionCard",
    params: data
  })
}
export const deleteDefinefunctionCard = (data) => {
  return myAxios({
    method: "GET",
    url: "/deleteDefinefunctionCard",
    params: data
  })
}
export const uploadModelFile = (data) => {
  return myAxios({
    method: 'POST',
    url: '/uploadModelFile',
    data: data
  })
}
export const uploadModelPNGFile = (data) => {
  return myAxios({
    method: 'POST',
    url: '/uploadModelPNGFile',
    data: data
  })
}
export const sendXmlInfo = (data) => {
  return myAxios({
    method: "POST",
    url: "/xmlinfo",
    responseType: "blob",
    params: data
  })
}
export const uploadReward = (data) => {
  return myAxios({
    method: 'POST',
    url: '/uploadReward',
    data: data
  })
}
export const uploadActions = (data) => {
  return myAxios({
    method: 'POST',
    url: '/uploadActions',
    data: data
  })
}
export const uploadLoss = (data) => {
  return myAxios({
    method: 'POST',
    url: '/uploadLoss',
    data: data
  })
}
export const getprocessFile = (data) => {
  return myAxios({
    method: 'GET',
    url: '/getprocessFile',
    params:data
  })
}


