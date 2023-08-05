import { Graph, FunctionExt, Shape, Addon } from '@antv/x6';
import './shape';
import graphData from './data';
import axios from 'axios';
export default class FlowGraph {
  static graph: Graph;
  static dnd: Addon.Dnd;

  static init() {
    this.graph = new Graph({
      container: document.getElementById('container')!,
      width: 1600,
      height: 1050,
      grid: {
        size: 10,
        visible: true,
        type: 'doubleMesh',
        args: [
          {
            color: '#cccccc',
            thickness: 1,
          },
          {
            color: '#5F95FF',
            thickness: 1,
            factor: 4,
          },
        ],
      },
      panning: true,
      snapline: true,
      history: true,
      // scroller: {
      //   enabled: true,
      //   pageVisible: true,
      //   pageBreak: true,
      //   pannable: false,
      // },
      mousewheel: {
        enabled: true,
        modifiers: ['ctrl', 'meta'],
        minScale: 0.5,
        maxScale: 2,
      },

      selecting: {
        enabled: true,
        multiple: true,
        rubberband: true,
        movable: true,
        showNodeSelectionBox: true,
      },
      connecting: {
        anchor: 'center',
        connectionPoint: 'anchor',
        allowBlank: false,
        highlight: true,
        snap: true,
        createEdge() {
          return new Shape.Edge({
            attrs: {
              line: {
                stroke: '#5F95FF',
                strokeWidth: 1,
                targetMarker: {
                  name: 'classic',
                  size: 8,
                },
              },
            },
            router: {
              name: 'manhattan',
            },
          });
        },
        validateConnection({ sourceView, targetView, sourceMagnet, targetMagnet }) {
          if (sourceView === targetView) {
            return false;
          }
          if (!sourceMagnet) {
            return false;
          }
          if (!targetMagnet) {
            return false;
          }
          return true;
        },
      },
      highlighting: {
        magnetAvailable: {
          name: 'stroke',
          args: {
            padding: 4,
            attrs: {
              strokeWidth: 4,
              stroke: 'rgba(223,234,255)',
            },
          },
        },
      },
      clipboard: {
        enabled: true,
      },
      keyboard: {
        enabled: true,
      },
      embedding: {
        enabled: true,
        findParent({ node }) {
          const bbox = node.getBBox();
          return this.getNodes().filter((node) => {
            // 只有 data.parent 为 true 的节点才是父节点
            const data = node.getData<any>();
            if (data && data.parent) {
              const targetBBox = node.getBBox();
              return bbox.isIntersectWithRect(targetBBox);
            }
            return false;
          });
        },
      },
    });

    this.dnd = new Addon.Dnd({
      target: this.graph,
      getDropNode(node) {
        const { width, height } = node.size()
        // 返回一个新的节点作为实际放置到画布上的节点
        return node.clone().size(width , height )
      },
    });


    this.initGraphShape();
    this.initEvent();
    return this.graph,this.dnd;
  }

  // 初始化设置好的画布上的元素
  static initGraphShape() {
    var url = decodeURI(window.location.href);
    var cs_arr = url.split('?')[1];//?后面的
    var iid = cs_arr.split('=')[1];
    axios.get("http://localhost:3000/design?id="+iid).then(response => {

      console.log(response);

      var graphData1 = response.data[0]
      this.graph.fromJSON(graphData1 as any);
    })
        .catch(error => {
          console.log(error);
        })
  }

  static showPorts(ports: NodeListOf<SVGAElement>, show: boolean) {
    for (let i = 0, len = ports.length; i < len; i = i + 1) {
      ports[i].style.visibility = show ? 'visible' : 'hidden';
    }
  }

  static initEvent() {
    const { graph } = this;
    const container = document.getElementById('container')!;

    graph.on('node:contextmenu', ({ cell, view }) => {
      console.log(view.container);
      const oldText = cell.attr('text/text') as string;
      cell.attr('text/style/display', 'none');
      const elem = view.container.querySelector('.x6-edit-text') as HTMLElement;
      if (elem) {
        elem.innerText = oldText;
        elem.focus();
      }
      const onBlur = () => {
        cell.attr('text/text', elem.innerText);
      };
      if (elem) {
        elem.addEventListener('blur', () => {
          onBlur();
          elem.removeEventListener('blur', onBlur);
        });
      }
    });
    graph.on(
      'node:mouseenter',
      FunctionExt.debounce(() => {
        const ports = container.querySelectorAll('.x6-port-body') as NodeListOf<SVGAElement>;
        this.showPorts(ports, true);
      }),
      500,
    );
    graph.on('node:mouseleave', () => {
      const ports = container.querySelectorAll('.x6-port-body') as NodeListOf<SVGAElement>;
      this.showPorts(ports, false);
    });

    graph.on('node:collapse', ({ node, e }: any) => {
      e.stopPropagation();
      node.toggleCollapse();
      const collapsed = node.isCollapsed();
      const cells = node.getDescendants();
      cells.forEach((n: any) => {
        if (collapsed) {
          n.hide();
        } else {
          n.show();
        }
      });
    });
    // backspace
    graph.bindKey('delete', () => {
      const cells = graph.getSelectedCells();
      if (cells.length) {
        graph.removeCells(cells);
      }
    });
  }
}
