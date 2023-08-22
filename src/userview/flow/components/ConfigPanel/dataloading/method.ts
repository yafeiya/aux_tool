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
    globalGridAttr.nodedataurl = cell.getData()?.dataurl;
    globalGridAttr.nodeheadernum = cell.getData()?.headernum;
    globalGridAttr.nodedatasetsize = cell.getData()?.datasetsize;
    globalGridAttr.nodeimgsize = cell.getData()?.imgsize;
    globalGridAttr.nodename = cell.getData()?.name;
    globalGridAttr.selflabel = cell.getData()?.selflabel
    // console.log(globalGridAttr.selflabel)
    // console.log(globalGridAttr.nodename)
  }
  return curCel;
}
