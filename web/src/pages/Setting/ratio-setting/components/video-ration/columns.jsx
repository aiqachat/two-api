import { Button, Modal, Space } from '@douyinfe/semi-ui';
import { IconDelete, IconEdit } from '@douyinfe/semi-icons';
import React from 'react';
import service from './service';
import { editModal } from './components/edit';

export const columns = (refresh) => {
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
      render: (_, { id }) => (
        <Space>
          <Button
            type='primary'
            icon={<IconEdit />}
            onClick={() => {
              editModal.open({ id }, () => {
                refresh();
              });
            }}
          ></Button>
          <Button
            icon={<IconDelete />}
            type='danger'
            onClick={() => {
              console.log('#');
              const modal = Modal.confirm({
                open: true,
                title: '删除模型',
                content: '确定删除该模型吗？',
                okText: '确定',
                cancelText: '取消',
                onOk: async () => {
                  try {
                    modal.update({ okButtonProps: { loading: true } });
                    await service.delWsVideoRation(id);
                    refresh();
                  } catch (e) {
                  } finally {
                    modal.update({ okButtonProps: { loading: false } });
                  }
                },
              });
            }}
          />
        </Space>
      ),
    },
  ];
};
