import { getMenuInfo } from '../../api/api.js'


export function readlist(menu){
    
    let list = menu
    console.log("readlist",list)

    // getMenuInfo("designmenu").then(res => {
    //     list =res.data
    //     // console.log("outgetmenu",list)
    //   })
    
    return list;
    
    
}
