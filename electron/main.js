const {app, BrowserWindow,  dialog, Menu} = require('electron')
const path = require('path')

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

let mainWindow
//mainwindow
function createWindow() {
  mainWindow = new BrowserWindow({
        width:1600,
        height:900,
        center: true, // 是否出现在屏幕居中的位置
        minWidth: 1366,     // 最小宽度
        minHeight: 768,    // 最小高度
        maxWidth: 1920,    // 最大宽度
        maxHeight: 1080,    // 最大高度
        //fullscreen: true   //全屏
        // titleBarStyle: "hidden",
        // titleBarOverlay: {
        //   color: "#fff",
        //   symbolColor: "black"
        //  },
        frame: true,   	//设置为 false 时可以创建一个无边框窗口
        // icon: path.join(__dirname, './public/icons/girl.ico'),//应用运行时的标题栏图标
        webPreferences: {
            nodeIntegration: true,
            contextIsolation:false,
            webSecurity: false,
        }
    })

    process.env['ELECTRON_DISABLE_SECURITY_WARNINGS'] = 'true'
    if (process.env.VITE_DEV_SERVER_URL) {
      mainWindow.loadURL(process.env.VITE_DEV_SERVER_URL)
    } else {
      mainWindow.loadFile(path.resolve(__dirname, '../dist/index.html'));
    }

    mainWindow.on('close', e => {
      e.preventDefault(); //先阻止一下默认行为，不然直接关了，提示框只会闪一下
      dialog.showMessageBox({
          type: 'info',
          title: '提示',
          message:'确认退出？',
          buttons: ['确认', '取消'],   //选择按钮，点击确认则下面的idx为0，取消为1
          cancelId: 1, //这个的值是如果直接把提示框×掉返回的值，这里设置成和“取消”按钮一样的值，下面的idx也会是1
      }).then(idx => {
      //注意上面↑是用的then，网上好多是直接把方法做为showMessageBox的第二个参数，我的测试下不成功
          console.log(idx)
          if (idx.response == 1) {
              console.log('index==1,cancel')
              e.preventDefault();
              // mainWindow.minimize();
          } else {
              console.log('index==0,close')
              mainWindow = null
              app.exit();
          }
          })
      });

};

app.whenReady().then(()=>{
    createWindow()
})


