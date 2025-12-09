import { Button, Space } from '@douyinfe/semi-ui';
import { IconDelete, IconEdit } from '@douyinfe/semi-icons';
import React from 'react';

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
      title: '模型名称',
      dataIndex: 'model_name',
      key: 'model_name',
    },
    {
      title: '每秒价格',
      dataIndex: 'config',
      key: 'config',
      width: 300,
      render: (config) => {
        return (
          <>
            {Object.entries(config).map(([key, value]) => {
              return (
                <div key={key}>
                  分辨率({key}): {value}元/秒
                </div>
              );
            })}
          </>
        );
      },
    },
    {
      title: '操作',
      key: 'action',
      width: 130,
      render: (_, record) => (
        <Space>
          <Button
            type='primary'
            icon={<IconEdit />}
            onClick={() => {
              console.log('#');
            }}
          ></Button>
          <Button
            icon={<IconDelete />}
            type='danger'
            onClick={() => {
              console.log('#');
            }}
          />
        </Space>
      ),
    },
  ];
};
