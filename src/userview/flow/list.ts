export function menulist(){
    const list = [
        // {
        //     path: '/home',
        //     name: 'home',
        //     label: '| 返回首页',
        //     icon: '',
        //     url: '/home',
        // },
        {
            name: 'dataloading',
            label: '数据加载',
            children: [
                {
                    name: 'csvdir',
                    label: 'csv文件',
                    content:[
                    {
                        name:'csv文件1',
                    },
                    {
                        name:'csv文件2',

                    },
                    {
                        name:'csv文件3',
                    },
                    ]
                },
                {
                    name: 'imgdataset',
                    label: '数值数据集',
                    content:[
                    {
                        name:'数值数据集1',
                    },
                    {
                        name:'数值数据集2',
                    },
                    {
                        name:'数值数据集3',
                    },
                    ]
                },
                {
                    name: 'monitordata',
                    label: '仿真监听数据',
                },
                {
                    name: 'modellog',
                    label: '模型日志数据',
                },
            ]

        },
        {
            name: 'datapre',
            label: '数据预处理',
            children: [
                {
                    name: 'dataaugmentation',
                    label: '数据扩充',
                },
                {
                    name: 'dataaugmentation',
                    label: '数据清洗',
                },
            ]
        },
        {
            name: 'moudles',
            label: '模型模板',
            children: [
                {
                name: 'machinelearning',
                label: '机器学习',
                children: [
                    {
                        name: 'regression',
                        label: '回归',
                        content:[
                            {
                                name:'回归1',
                            },
                            {
                                name:'回归2',
                            },
                            {
                                name:'回归3',
                            },
                            ]
                    },
                    {
                        name: 'classification',
                        label: '分类',
                        content:[
                            {
                                name:'分类1',
                            },
                            {
                                name:'分类2',
                            },
                            ]
                    },
                ]
                },
                {
                name: 'deeplearning',
                label: '深度学习',
                children: [
                    {
                        name: 'CNN',
                        label: '卷积神经网络',
                        content:[
                            {
                                name:'卷积神经网络1',
                            },
                            {
                                name:'卷积神经网络2',
                            },
                            ]
                    },
                    {
                        name: 'RNN',
                        label: '循环神经网络',
                    },
                ]
                },
                {
                name: 'reinforcementlearning',
                label: '强化学习',
                children:[                  
                    {
                        name: 'SAC',
                        label: '值迭代',
                    },
                    {
                        name: 'DQN',
                        label: '策略学习',
                    },
                ]
                },
            ]
        },
        {
            name: 'moudlepredictive',
            label: '模型预测',
            children: [
                {
                    name: 'batchdeduction',
                    label: '批量推理',
                },
            ]
        },
        {
            name: 'resultvisualization',
            label: '模型结果可视化',
            children: [
                {
                    name: 'actdistribution',
                    label: '动作分布',
                },
                {
                    name: 'rewarddistribution',
                    label: '奖励分布',
                },
                {
                    name: 'learningrate',
                    label: '学习率',
                },
            ]
        },
        {
            name: 'scorefunction',
            label: '评价函数',
            children: [
                {
                    name: 'composite',
                    label: '综合指标',
                },
                {
                    name: 'precision',
                    label: '模型精度',
                },
                {
                    name: 'generalization',
                    label: '泛化性',
                },
                {
                    name: 'robustness',
                    label: '鲁棒性',
                },
                {
                    name: 'custom',
                    label: '自定义指标',
                },
            ]
        },
    ];

    return list;
}
