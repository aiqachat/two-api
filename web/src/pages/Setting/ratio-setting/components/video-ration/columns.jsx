export const columns = () => {
  return [
    {
      title: 'ID',
      dataIndex: 'id',
      key: 'id',
      width: 100,
      fixed: 'left',
    },
    {
      title: '积分包名称',
      dataIndex: 'name',
      key: 'name',
      width: 120,
    },
    {
      title: '积分包余量',
      dataIndex: 'info',
      key: 'info',
      render: (_) => {
        return <div style={{ display: 'flex', gap: 20 }}>{_}</div>;
      },
    },
  ];
};
