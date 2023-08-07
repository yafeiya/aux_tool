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
    globalGridAttr.nodeiterations = cell.getData()?.iterations;
    // console.log(cell.getData()?.canshu1)
    globalGridAttr.nodeoptimizer = cell.getData()?.optimizer;
    globalGridAttr.nodedecayfactor = cell.getData()?.decayfactor;
    globalGridAttr.nodelearningrate = cell.getData()?.learningrate;
    globalGridAttr.nodeloss = cell.getData()?.loss;
    globalGridAttr.nodeevalution = cell.getData()?.evalution;
  }
  return curCel;
}
