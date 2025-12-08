import { Form, Modal } from '@douyinfe/semi-ui';
import { wsCreateModalHandle } from '@components';
import { useEffect, useRef } from 'react';
import { WsError } from '@helpers';
import service from './service';

export const EditModal = ({ modalProps, onComplete, edit = true, id }) => {
  const formRef = useRef(null);
  useEffect(() => {
    formRef.current?.formApi.setValues({
      name: '视频比率',
    });
  }, [edit]);
  return (
    <Modal
      {...modalProps}
      title='编辑视频比率'
      onOk={async () => {
        try {
          const values = await formRef.current?.formApi.validate();
          await service.createWsVideoRation(values);
          onComplete(true);
          modalProps.onCancel();
        } catch (e) {
          WsError.handleError(e);
        }
      }}
    >
      <Form ref={formRef}>
        <Form.Input
          label='模型名称'
          field='modeName'
          rules={[{ required: true }]}
        />
        <Form.Input
          label='分辨率'
          field='resolution'
          rules={[{ required: true }]}
        />
        <Form.InputNumber
          label='每秒价格'
          field='price'
          rules={[{ required: true }]}
          precision={2}
          step={1}
          min={0}
          max={99999999999}
        />
      </Form>
    </Modal>
  );
};

export const editModal = wsCreateModalHandle(EditModal);
