const graphData = {
  cells: [
    {
      position: {
        x: 420,
        y: 160,
      },
      size: {
        width: 80,
        height: 42,
      },
      attrs: {
        text: {
          text: '数据集',
        },
        body: {
          rx: 24,
          ry: 24,
        },
      },
      data:{
        canshu1:111,
        canshu2:222,
        canshu3:333,
      },
      shape: 'flow-chart-rect',
      ports: {
        groups: {
          top: {
            position: 'top',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
          right: {
            position: 'right',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
          bottom: {
            position: 'bottom',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
          left: {
            position: 'left',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
        },
        items: [
          {
            group: 'top',
            id: '45726225-0a03-409e-8475-07da4b8533c5',
          },
          {
            group: 'right',
            id: '06111939-bf01-48d9-9f54-6465d9d831c6',
          },
          {
            group: 'bottom',
            id: '6541f8dc-e48b-4b8c-a105-2ab3a47f1f21',
          },
          {
            group: 'left',
            id: '54781206-573f-4982-a21e-5fac1e0e8a60',
          },
        ],
      },
      id: '8650a303-3568-4ff2-9fac-2fd3ae7e6f2a',
      zIndex: 1,
    },
    {
      position: {
        x: 420,
        y: 250,
      },
      size: {
        width: 80,
        height: 42,
      },
      attrs: {
        text: {
          text: '自定义函数',
        },
      },
      data:{
        canshu1:11,
        canshu2:22,
        canshu3:33,
      },
      shape: 'flow-chart-rect',
      ports: {
        groups: {
          top: {
            position: 'top',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
          right: {
            position: 'right',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
          bottom: {
            position: 'bottom',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
          left: {
            position: 'left',
            attrs: {
              circle: {
                r: 3,
                magnet: true,
                stroke: '#5F95FF',
                strokeWidth: 1,
                fill: '#fff',
                style: {
                  visibility: 'hidden',
                },
              },
            },
          },
        },
        items: [
          {
            group: 'top',
            id: 'd1346f43-969a-4201-af5d-d09b7ef79980',
          },
          {
            group: 'right',
            id: 'd561926a-3a24-449a-abb1-0c20bc89947e',
          },
          {
            group: 'bottom',
            id: '0cbde5df-ef35-410e-b6c3-a6b1f5561e3f',
          },
          {
            group: 'left',
            id: '2fceb955-f7af-41ac-ac02-5a2ea514544e',
          },
        ],
      },
      id: '7b6fd715-83e6-4053-8c2b-346e6a857bf3',
      zIndex: 2,
    },
    {
      shape: 'edge',
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
      id: '00f3c401-8bad-46b9-b692-232aa011d4c5',
      router: {
        name: 'manhattan',
      },
      zIndex: 3,
      source: {
        cell: '8650a303-3568-4ff2-9fac-2fd3ae7e6f2a',
        port: '6541f8dc-e48b-4b8c-a105-2ab3a47f1f21',
      },
      target: {
        cell: '7b6fd715-83e6-4053-8c2b-346e6a857bf3',
        port: 'd1346f43-969a-4201-af5d-d09b7ef79980',
      },
    },

    

    
  ],
};

export default graphData;
