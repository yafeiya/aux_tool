import FlowGraph from '../../../graph/index';

export function gridOpt(globalGridAttr: any) {
  let options;
  if (globalGridAttr.type === 'doubleMesh') {
    options = {
      type: globalGridAttr.type,
      args: [
        {
          color: globalGridAttr.color,
          thickness: globalGridAttr.thickness,
        },
        {
          color: globalGridAttr.colorSecond,
          thickness: globalGridAttr.thicknessSecond,
          factor: globalGridAttr.factor,
        },
      ],
    };
  } else {
    options = {
      type: globalGridAttr.type,
      args: [
        {
          color: globalGridAttr.color,
          thickness: globalGridAttr.thickness,
        },
      ],
    };
  }
  const { graph } = FlowGraph;
  graph.drawGrid(options);
}

export function gridSizeOpt(globalGridAttr: any) {
  const { graph } = FlowGraph;
  graph.setGridSize(globalGridAttr.size);
}
