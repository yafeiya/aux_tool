import { app,Menu,BrowserWindow } from 'electron';
import path from 'path'

//auto open json-server
const jsonServer = require('json-server')
const server = jsonServer.create()
const router = jsonServer.router('./data/data.json')
const middlewares = jsonServer.defaults()

server.use(middlewares)

server.get('/echo', (req, res) => {
  res.jsonp(req.query)
})

server.use(jsonServer.bodyParser)
server.use((req, res, next) => {
  if (req.method === 'POST') {
    req.body.createdAt = Date.now()
  }
  next()
})

server.use(router)
server.listen(3000, () => {
  console.log('JSON Server is running')
})

//mainwindow
const createdWindow = () => {
    const win = new BrowserWindow({
        width:900,
        height:600,
        center: true, // 是否出现在屏幕居中的位置
        //  minWidth: 300,     // 最小宽度
        //  minHeight: 500,    // 最小高度
         maxWidth: 1920,    // 最大宽度
         maxHeight: 1080,    // 最大高度
        //fullscreen: true   //全屏
        frame: true,   	//设置为 false 时可以创建一个无边框窗口
        
        webPreferences: {
            webSecurity: false,
            //...
          },
      
        // icon: path.join(__dirname, './public/icons/girl.ico'),//应用运行时的标题栏图标
        webPreferences: {
            nodeIntegration: true,
            contextIsolation:false,
        }
    })

    if (process.env.VITE_DEV_SERVER_URL) {
        win.loadURL(process.env.VITE_DEV_SERVER_URL)
    } else {
        win.loadFile(path.resolve(__dirname, '../dist/index.html'));
    }

    // MenuTool
    // const menuTemplate = [
    //   {
    //     label:"文件", // 菜单名
    //     submenu:[ // 二级菜单
    //       //  {
    //       //   label:'新建',
    //       //   accelerator:'ctrl+n', // 快捷键
    //       //   click:()=>{ // 点击事件
    //       //     console.log("create file")
    //       //   }
    //       // },
    //       {type: 'separator'}, // 分割线
    //       {
    //         label:'退出',
    //           click:()=>{ // 点击事件
    //           console.log("exit mainwindow")
                  
    //         }
    //       },
    //       {type: 'separator'}, // 分割线
    //       // {label:'全屏切换',role:'togglefullscreen'},
    //       {label: '撤销', role: 'undo'},     
  
          

    //     ],
    //   },
        
    //   // {
    //   //   label:"编辑",
    //   //   submenu:[ 
    //   //      {
    //   //       label:'复制', 
    //   //       role:'copy', // 快捷键与系统冲突时可以使用role属性指定
    //   //       click:()=>{ 
    //   //         console.log("copy")
    //   //       }
    //   //     },
    //   //     {
    //   //       label:'粘贴',
    //   //       role:'paste',
    //   //       click:()=>{
    //   //         console.log("paste")
    //   //       }
    //   //     },
    //   //     {
    //   //       label:'保存',
    //   //     }
    //   //   ],
    //   // },
    // ]
    // const menuBuilder = Menu.buildFromTemplate(menuTemplate)
    // Menu.setApplicationMenu(menuBuilder)

}

app.whenReady().then(()=>{
    createdWindow()
})
