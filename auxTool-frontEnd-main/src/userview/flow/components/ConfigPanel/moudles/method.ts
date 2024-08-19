import FlowGraph from '../../../graph/index';


export function nodeOpt(id: any, globalGridAttr: any) {
  let curCel: any = null;
  if (id.value) {
    const { graph } = FlowGraph;
    const cell = graph.getCellById(id.value);
    if (!cell || !cell.isNode()) {
      return;
    }
    curCel = cell;
    globalGridAttr.nodeStroke = curCel.attr('body/stroke');
    globalGridAttr.nodeStrokeWidth = cell.attr('body/strokeWidth');
    globalGridAttr.nodeFill = cell.attr('body/fill');
    globalGridAttr.nodeFontSize = cell.attr('text/fontSize');
    globalGridAttr.nodeColor = cell.attr('text/fill');
    globalGridAttr.nodenetworkdepth = cell.getData()?.networkdepth;
    globalGridAttr.nodeclassnum = cell.getData()?.classnum;
    globalGridAttr.nodefuturerewarddiscount = cell.getData()?.futurerewarddiscount;
    globalGridAttr.nodemodelurl = cell.getData()?.modelurl;
    globalGridAttr.nodemodeltype = cell.getData()?.modeltype;
    globalGridAttr.nodemodelbatch = cell.getData()?.modelbatch;
    if(cell.getData()?.name == undefined){cell.getData().name = cell.getData()?.selflabel}
    globalGridAttr.nodename = cell.getData()?.name;
    // console.log(cell.getData().name == undefined)
    globalGridAttr.selflabel = cell.getData()?.selflabel
  }
  return curCel;
}
