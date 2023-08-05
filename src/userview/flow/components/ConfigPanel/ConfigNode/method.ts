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
    globalGridAttr.nodecanshu1 = cell.getData()?.canshu1;
    // console.log(cell.getData()?.canshu1)
    globalGridAttr.nodecanshu2 = cell.getData()?.canshu2;
    globalGridAttr.nodecanshu3 = cell.getData()?.canshu3;
    globalGridAttr.nodecanshu4 = cell.getData()?.canshu4;
    globalGridAttr.nodecanshu5 = cell.getData()?.canshu5;
  }
  return curCel;
}
