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
    globalGridAttr.Network_num = cell.getData()?.Network_num;
    globalGridAttr.Optimizer = cell.getData()?.Optimizer;
    globalGridAttr.batch = cell.getData()?.batch;
    globalGridAttr.learning_rate = cell.getData()?.learning_rate;
    globalGridAttr.Epoch_num = cell.getData()?.Epoch_num;
    globalGridAttr.Act_function = cell.getData()?.Act_function;
    globalGridAttr.Decay_factor = cell.getData()?.Decay_factor;
    globalGridAttr.Explore_rate = cell.getData()?.Explore_rate;
    globalGridAttr.Radom_seed = cell.getData()?.Radom_seed;


    globalGridAttr.nodemodeltype = cell.getData()?.modeltype;
    globalGridAttr.nodemodelbatch = cell.getData()?.modelbatch;
    if(cell.getData()?.name == undefined){cell.getData().name = cell.getData()?.selflabel}
    globalGridAttr.nodename = cell.getData()?.name;
    // console.log(cell.getData().name == undefined)
    globalGridAttr.selflabel = cell.getData()?.selflabel
  }
  return curCel;
}
