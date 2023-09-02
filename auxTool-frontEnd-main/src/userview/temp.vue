<template>
  <Table :columns="columns" :data="data" border height="500"></Table>
</template>
<script>
export default {
  data () {
    return {
      columns: [
        {
          title: 'Name',
          key: 'name',
          align: 'center',
          width: 200,
          fixed: 'left',
          filters: [
            {
              label: 'Joe',
              value: 1
            },
            {
              label: 'John',
              value: 2
            }
          ],
          filterMultiple: false,
          filterMethod (value, row) {
            if (value === 1) {
              return row.name === 'Joe';
            } else if (value === 2) {
              return row.name === 'John Brown';
            }
          }
        },
        {
          title: 'Other',
          align: 'center',
          children: [
            {
              title: 'Age',
              key: 'age',
              align: 'center',
              width: 200,
              sortable: true
            },
            {
              title: 'Address',
              align: 'center',
              children: [
                {
                  title: 'Street',
                  key: 'street',
                  align: 'center',
                  width: 200
                },
                {
                  title: 'Block',
                  align: 'center',
                  children: [
                    {
                      title: 'Building',
                      key: 'building',
                      align: 'center',
                      width: 200,
                      sortable: true
                    },
                    {
                      title: 'Door No.',
                      key: 'door',
                      align: 'center',
                      width: 200
                    }
                  ]
                }
              ]
            }
          ]
        },
        {
          title: 'Company',
          align: 'center',
          children: [
            {
              title: 'Company Address',
              key: 'caddress',
              align: 'center',
              width: 200
            },
            {
              title: 'Company Name',
              key: 'cname',
              align: 'center',
              width: 200
            }
          ]
        },
        {
          title: 'Gender',
          key: 'gender',
          align: 'center',
          width: 200,
          fixed: 'right'
        }
      ],
      data: []
    }
  },
  mounted () {
    const data = [];
    for (let i = 0; i < 20; i++) {
      data.push({
        key: i,
        name: 'John Brown',
        age: i + 1,
        street: 'Lake Park',
        building: 'C',
        door: 2035,
        caddress: 'Lake Street 42',
        cname: 'SoftLake Co',
        gender: 'M',
      });
    }
    this.data = data;
  }
}
</script>
